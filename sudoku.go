package main

import "math/bits"

const GridSize = 9
const MiniGridSize = 3
const CellStateResolved = 0
const CellStateUnresolved = 0x1FF

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

func (grid *Grid) solve() bool {
	if len(grid.resolveQueue) == 0 {
		grid.processGroups()
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
