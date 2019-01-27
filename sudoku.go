package main

import (
	"fmt"
	"math/bits"
)

const GridSize = 9
const MiniGridSize = 3
const CellStateResolved uint = 0
const CellStateUnresolved uint = 0x1FF

type Cell struct {
	state           uint
	value           int
	associatedCells [20]int
}

type Grid struct {
	cells        [GridSize * GridSize]Cell
	groups       [GridSize * 3][]int
	resolveQueue []int
}

func (grid *Grid) updateState(i int, value int) {
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
	for _, group := range grid.groups {
		oneCell := CellStateResolved
		multiCell := CellStateResolved
		for _, cellId := range group {
			multiCell |= oneCell & grid.cells[cellId].state
			oneCell |= grid.cells[cellId].state
		}
		oneCell &^= multiCell
		for _, cellId := range group {
			if oneCell & grid.cells[cellId].state != 0 && bits.OnesCount(grid.cells[cellId].state) > 1 {
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

	// no more cells to process, check if grid has been solved
	if len(grid.resolveQueue) == 0 {
		for _, cell := range grid.cells {
			if cell.state != CellStateResolved {
				return false
			}
		}
		return true
	}

	cell := &grid.cells[grid.resolveQueue[0]]
	grid.resolveQueue = grid.resolveQueue[1:]
	if cell.state == CellStateResolved {
		// illegal state
		println("ILLEGAL STATE")
		return false
	}
	value := bits.TrailingZeros(cell.state)
	cell.value = value
	cell.state = CellStateResolved

	// update states of associated cells
	for _, associatedCell := range cell.associatedCells {
		grid.updateState(associatedCell, value)
	}
	return grid.solve()
}
