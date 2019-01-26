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


}
