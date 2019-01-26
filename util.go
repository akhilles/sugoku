package main

import "fmt"

func (grid *Grid) init() {
	grid.resolveQueue = make([]int, 0, GridSize*GridSize)
	var cellGroups [GridSize * GridSize][3]int
	for i := range grid.cells {
		rowId := i / GridSize
		colId := i % GridSize
		miniGridId := (rowId/MiniGridSize)*MiniGridSize + (colId / MiniGridSize)
		cellGroups[i][0] = rowId
		cellGroups[i][1] = GridSize+colId
		cellGroups[i][2] = GridSize*2+miniGridId
		for _, groupId := range cellGroups[i] {
			grid.groups[groupId] = append(grid.groups[groupId], i)
		}
	}
	for i := range grid.cells {
		cell := &grid.cells[i]
		cell.state = CellStateUnresolved

		associatedCells := make(map[int]bool)
		for _, groupId := range cellGroups[i] {
			for _, cellId := range grid.groups[groupId] {
				associatedCells[cellId] = true
			}
		}
		delete(associatedCells, i)
		i := 0
		for k := range associatedCells {
			cell.associatedCells[i] = k
			i++
		}
	}
}

func (grid *Grid) addToResolveQueue(i int, value int) {
	grid.cells[i].state = 1 << uint(value)
	grid.resolveQueue = append(grid.resolveQueue, i)
}

func (grid *Grid) print() {
	for i, cell := range grid.cells {
		if cell.state == CellStateResolved {
			fmt.Printf("%2v(%9b) ", cell.value, cell.state)
		} else {
			fmt.Printf(" -(%9b) ", cell.state)
		}
		if (i+1)%GridSize == 0 {
			fmt.Println()
		}
	}
}
