package main

import (
	"fmt"
	"math/bits"
)

type SolveResult int

const (
	Solved SolveResult = iota
	NoSolution
	MultipleSolutions
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
	cells [GridSize * GridSize]Cell
	// TODO: switch to fixed size stack (more performant)
	resolveQueue []int
}

type GridInfo struct {
	linkedCells [GridSize * GridSize][20]int
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
		for _, i := range group {
			multiCell |= oneCell & grid.cells[i].state
			oneCell |= grid.cells[i].state
		}
		oneCell &^= multiCell
		for _, i := range group {
			if oneCell&grid.cells[i].state != 0 && bits.OnesCount(grid.cells[i].state) > 1 {
				grid.cells[i].state &= oneCell
				grid.resolveQueue = append(grid.resolveQueue, i)
			}
		}
	}
}

func (grid *Grid) guessSolve(i int) SolveResult {
	numSolutions := 0
	var solutionGrid *Grid
	for grid.cells[i].state != CellStateResolved {
		value := bits.TrailingZeros(grid.cells[i].state)
		fmt.Println(">>> cell", i, "- guessing", value+1)

		grid.cells[i].state &^= 1 << uint(value)
		newGrid := &Grid{}
		newGrid.cells = grid.cells
		newGrid.cells[i].state = 1 << uint(value)
		newGrid.resolveQueue = []int{i}
		solveResult := newGrid.solve()
		if solveResult == Solved {
			numSolutions++
			solutionGrid = newGrid
		}
		fmt.Println("<<< cell", i, "- guessing", value+1)
		if solveResult == MultipleSolutions || numSolutions > 1 {
			return MultipleSolutions
		}
	}

	if numSolutions == 1 {
		grid.cells = solutionGrid.cells
		return Solved
	}
	return NoSolution
}

func (grid *Grid) solve() SolveResult {
	// fmt.Println("Q LEN --> ", len(grid.resolveQueue))
	if len(grid.resolveQueue) == 0 {
		//fmt.Print("group processing -> ")
		grid.processGroups()
		//fmt.Println(len(grid.resolveQueue))
	}

	if len(grid.resolveQueue) == 0 {
		// find an unresolved cell
		for i := range grid.cells {
			if grid.cells[i].state != CellStateResolved {
				// make guesses
				return grid.guessSolve(i)
			}
		}
		return Solved
	}

	cellIndex := grid.resolveQueue[0]
	grid.resolveQueue = grid.resolveQueue[1:]
	cell := &grid.cells[cellIndex]
	if cell.state == CellStateResolved {
		// illegal state
		return NoSolution
	}

	// resolve cell
	value := bits.TrailingZeros(cell.state)
	cell.value = value
	cell.state = CellStateResolved
	for _, associatedCell := range gridInfo.linkedCells[cellIndex] {
		grid.updateCellState(associatedCell, value)
	}
	return grid.solve()
}
