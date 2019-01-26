package main

import (
	"fmt"
)

func main() {
	grid := Grid{}
	grid.init()
	grid.addToResolveQueue(22, 6)
	grid.addToResolveQueue(66, 6)

	fmt.Println(grid.solve())

	grid.print()
}
