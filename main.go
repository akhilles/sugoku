package main

import "fmt"

func main() {
	grid := initGrid()
	grid.load("700893014354000000100600070060004000000389005503000940000506800000007400000400031")
	grid.print(true)
	fmt.Println(grid.solve())
	fmt.Println()
	grid.print(false)
}
