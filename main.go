package main

import "fmt"

func main() {
	grid := Grid{}
	grid.init()

	grid.load(`000500000035000100600009300020005490940206000057003600060000000000704856074001002`)

	grid.print(true)
	fmt.Println(grid.solve())
	fmt.Println()
	grid.print(false)
}
