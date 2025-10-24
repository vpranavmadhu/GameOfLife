package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	size uint
	data [][]bool
}

func newGrid(size uint, live ...int) Grid {

	grid := make([][]bool, size)

	for i := 0; i < int(size); i++ {
		grid[i] = make([]bool, size)
	}

	for i := 0; i < len(live); i = i + 2 {
		x := live[i]
		y := live[i+1]
		grid[x][y] = true
	}

	return Grid{
		size: size,
		data: grid,
	}
}

func displayGrid(grid Grid) string {
	var display strings.Builder

	for _, row := range grid.data {
		for _, val := range row {
			if val {
				display.WriteString("■")
			} else {
				display.WriteString("-")
			}
			display.WriteString(" ")
		}
		display.WriteString("\n")
	}
	return display.String()
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func countAliveNeighbours(grid Grid, x uint, y uint) int {

	if grid.size == 1 {
		return 0
	}

	if x == 0 {
		if y == 0 { //top left corner
			return b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x+1][y])
		} else if y == grid.size-1 { //top right corner
			return b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y])
		} else { //top edge
			return b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x][y+1])
		}
	}

	if x == grid.size-1 {
		if y == 0 { //bottom left corner
			return b2i(grid.data[x][y+1]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x-1][y])
		} else if y == grid.size-1 { //bottom right corner
			return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1])
		} else { //bottom edge
			return b2i(grid.data[x][y-1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1])
		}

	}

	if y == 0 { //left edge
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x+1][y])
	}
	if y == grid.size-1 { //right edge
		return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x+1][y-1])
	}

	return b2i(grid.data[x-1][y]) + b2i(grid.data[x+1][y]) + b2i(grid.data[x][y+1]) + b2i(grid.data[x][y-1]) + b2i(grid.data[x-1][y+1]) + b2i(grid.data[x+1][y+1]) + b2i(grid.data[x-1][y-1]) + b2i(grid.data[x+1][y-1])
}

func runGeneration(grid Grid) Grid {
	newGen := newGrid(grid.size)
	for i := 0; i < int(grid.size); i++ {
		for j := 0; j < int(grid.size); j++ {
			cell := grid.data[i][j]
			count := countAliveNeighbours(grid, uint(i), uint(j))

			if cell && count < 2 { //underpopulation
				newGen.data[i][j] = false
			}
			if cell && (count == 2 || count == 3) { //survival
				newGen.data[i][j] = true
			}
			if cell && count > 3 {
				newGen.data[i][j] = false
			}
			if !cell && count == 3 {
				newGen.data[i][j] = true
			}
		}
	}

	return newGen
}

func main() {
	size := 10
	grid := newGrid(uint(size), 0, 9, 4, 4, 9, 0, 0, 0, 9, 9)

	fmt.Println(displayGrid(grid))
	// for {

	// 	fmt.Println(displayGrid(grid))
	// 	grid = applyRules(grid)
	// 	grid = newGeneration(grid)

	// }

}
