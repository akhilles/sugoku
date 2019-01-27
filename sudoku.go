package main

import (
	"fmt"
	"math/bits"
)

const GridSize = 9
const MiniGridSize = 3
const CellStateResolved uint = 0
const CellStateUnresolved uint = 0x1FF

var gridInfo = initGridInfo()

type Cell struct {
	state uint
	value int
}

type Grid struct {
	cells        [GridSize * GridSize]Cell
	// TODO: switch to fixed size stack (more performant)
	resolveQueue []int
	resolved     int
}

type GridInfo struct {
	linkedCells [GridSize*GridSize][20]int
	groups      [GridSize * 3][]int
}

func (grid *Grid) updateCellState(i int, value int) {
	cell := &grid.cells[i]
	var stateModifier uint = 1 << uint(value)
	if cell.state&stateModifier != 0 {
		cell.state &^= stateModifier
		if bits.OnesCount(cell.state) == 1 {
			grid.resolveQueue = append(grid.resolveQueue, i)
		}
	}
}

// TODO: maybe optimize with references instead of values
func (grid *Grid) processGroups() {
	for _, group := range gridInfo.groups {
		oneCell := CellStateResolved
		multiCell := CellStateResolved
		for _, cellId := range group {
			multiCell |= oneCell & grid.cells[cellId].state
			oneCell |= grid.cells[cellId].state
		}
		oneCell &^= multiCell
		for _, cellId := range group {
			if oneCell&grid.cells[cellId].state != 0 && bits.OnesCount(grid.cells[cellId].state) > 1 {
				grid.cells[cellId].state &= oneCell
				grid.resolveQueue = append(grid.resolveQueue, cellId)
			}
		}
	}
}

func (grid *Grid) solve() bool {
	// fmt.Println("Q LEN --> ", len(grid.resolveQueue))
	if len(grid.resolveQueue) == 0 {
		fmt.Print("GROUP PROCESSING!!! --> ")
		grid.processGroups()
		fmt.Println(len(grid.resolveQueue))
	}

	if len(grid.resolveQueue) == 0 {
		if grid.resolved == GridSize*GridSize {
			return true
		}
		numSolutions := 0
		// start guessing
		return numSolutions == 1
	}

	cellIndex := grid.resolveQueue[0]
	grid.resolveQueue = grid.resolveQueue[1:]
	cell := &grid.cells[cellIndex]
	if cell.state == CellStateResolved {
		// illegal state
		println("ILLEGAL STATE")
		return false
	}
	value := bits.TrailingZeros(cell.state)
	cell.value = value
	cell.state = CellStateResolved
	grid.resolved++

	// update states of associated linkedCells
	for _, associatedCell := range gridInfo.linkedCells[cellIndex] {
		grid.updateCellState(associatedCell, value)
	}
	return grid.solve()
}
