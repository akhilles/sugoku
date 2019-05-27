package main

import (
	"fmt"
	"strings"
)

func initGridInfo() *GridInfo {
	gridLinks := &GridInfo{}

	var cellGroups [GridSize * GridSize][3]int
	for i := range gridLinks.linkedCells {
		rowId := i / GridSize
		colId := i % GridSize
		miniGridId := (rowId/MiniGridSize)*MiniGridSize + (colId / MiniGridSize)
		cellGroups[i][0] = rowId
		cellGroups[i][1] = GridSize+colId
		cellGroups[i][2] = GridSize*2+miniGridId
		for _, groupId := range cellGroups[i] {
			gridLinks.groups[groupId] = append(gridLinks.groups[groupId], i)
		}
	}
	for i := range gridLinks.linkedCells {
		links := &gridLinks.linkedCells[i]
		associatedCells := make(map[int]bool)
		for _, groupId := range cellGroups[i] {
			for _, cellId := range gridLinks.groups[groupId] {
				associatedCells[cellId] = true
			}
		}
		delete(associatedCells, i)
		i := 0
		for k := range associatedCells {
			links[i] = k
			i++
		}
	}
	return gridLinks
}

func initGrid() *Grid {
	grid := &Grid{}
	grid.resolveQueue = make([]int, 0, GridSize*GridSize)
	for i := range grid.cells {
		cell := &grid.cells[i]
		cell.state = CellStateUnresolved
		cell.value = -1
	}
	return grid
}

func (grid *Grid) addToResolveQueue(i int, value int) {
	grid.cells[i].state = 1 << uint(value)
	grid.resolveQueue = append(grid.resolveQueue, i)
}

func (grid *Grid) load(gridState string) {
	gridState = strings.Join(strings.Fields(gridState), "")
	for i, c := range gridState {
		val := c - '1'
		if val < 0 || val > 8 {
			continue
		}
		grid.addToResolveQueue(i, int(val))
	}
}

func (grid *Grid) print(debug bool) {
	for i, cell := range grid.cells {
		if debug {
			fmt.Printf("(%9b) ", cell.state)
		} else {
			fmt.Printf("%2v ", cell.value+1)
		}
		if (i+1)%GridSize == 0 {
			fmt.Println()
		}
	}
}
