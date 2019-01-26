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
