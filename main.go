package main

import (
	"fmt"
)

const GridSize = 9
const MiniGridSize = 3
const CellStateResolved = 0
const CellStateUnresolved = 0x1FF

type Cell struct {
	state uint
	value int
}

type Grid struct {
	cells        [GridSize * GridSize]Cell
	resolveQueue chan int
}

func (grid *Grid) init() {
	grid.resolveQueue = make(chan int, GridSize * GridSize)
	for i := range grid.cells {
		grid.cells[i].state = CellStateUnresolved
	}
}

func (grid *Grid) print() {
	for i, cell := range grid.cells {
		if cell.state == CellStateResolved {
			fmt.Printf("%2v(%9b) ", cell.value, cell.state)
		} else {
			fmt.Printf(" -(%9b) ", cell.state)
		}
		if (i + 1) % GridSize == 0 {
			fmt.Println()
		}
	}
}
