package main

import "fmt"

type Grid struct {
	size uint
	grid [][]bool
}

func newGrid(size uint) Grid {

	return Grid{
		size: size,
		grid: make([][]bool, size),
	}
}

func displayGrid(grid Grid) string {
	return "hi"
}

func newGeneration(grid Grid) Grid {
	return grid
}

func countNeighbours(grid Grid, x uint, y uint) uint {
	return 0
}

func countAliveNeighbours(grid Grid, x uint, y uint) uint {
	return 0
}

func applyRules(grid Grid) Grid {
	return grid
}

func main() {
	size := 10
	grid := newGrid(uint(size))
	//fmt.Println(grid)

	for {

		fmt.Println(displayGrid(grid))
		grid = applyRules(grid)
		grid = newGeneration(grid)

	}
}
