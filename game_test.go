package main

import (
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	size := 3
	actual := newGrid(uint(size))
	expected := Grid{
		size: 3,
		data: [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		},
	}

	if expected.size != actual.size {
		t.Errorf("Expected size: %d but got %d", expected.size, actual.size)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected : %v but got : %v", expected, actual)
	}

}

func TestDisplayGrid(t *testing.T) {

	grid := Grid{
		size: 3,
		data: [][]bool{
			{true, false, false},
			{false, true, false},
			{false, false, true},
		},
	}
	actual := displayGrid(grid)
	expected := "■ - - \n- ■ - \n- - ■ \n"

	if actual != expected {
		t.Errorf("Expected : \n%s ,but got : %s", expected, actual)
	}

}

func TestCountAliveNeighboursTopLeftCorner(t *testing.T) {
	grid := newGrid(5, 0, 1)
	actual := countAliveNeighbours(grid, 0, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 1, 1, 1)
	actual = countAliveNeighbours(grid, 0, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 0)
	actual = countAliveNeighbours(grid, 0, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursTopRightCorner(t *testing.T) {
	grid := newGrid(5, 0, 3)
	actual := countAliveNeighbours(grid, 0, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 3, 1, 3)
	actual = countAliveNeighbours(grid, 0, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 3, 1, 3, 1, 4)
	actual = countAliveNeighbours(grid, 0, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursBottomLeftCorner(t *testing.T) {
	grid := newGrid(5, 3, 0)
	actual := countAliveNeighbours(grid, 4, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 3, 0, 3, 1)
	actual = countAliveNeighbours(grid, 4, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 3, 0, 3, 1, 4, 1)
	actual = countAliveNeighbours(grid, 4, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursBottomRightCorner(t *testing.T) {
	grid := newGrid(5, 4, 3)
	actual := countAliveNeighbours(grid, 4, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 4, 3, 3, 3)
	actual = countAliveNeighbours(grid, 4, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 4, 3, 3, 3, 3, 4)
	actual = countAliveNeighbours(grid, 4, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}
}

func TestCountAliveNeighboursTopEdge(t *testing.T) {
	grid := newGrid(5, 0, 1)
	actual := countAliveNeighbours(grid, 0, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 0, 1, 1, 1)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2, 1, 3)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 0, 1, 1, 1, 1, 2, 1, 3, 0, 3)
	actual = countAliveNeighbours(grid, 0, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursBottomEdge(t *testing.T) {
	grid := newGrid(5, 4, 1)
	actual := countAliveNeighbours(grid, 4, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 4, 1, 3, 1)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2, 3, 3)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 4, 1, 3, 1, 3, 2, 3, 3, 4, 3)
	actual = countAliveNeighbours(grid, 4, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursLeftEdge(t *testing.T) {
	grid := newGrid(5, 1, 0)
	actual := countAliveNeighbours(grid, 2, 0)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 0, 1, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1, 3, 1)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 0, 1, 1, 2, 1, 3, 1, 3, 0)
	actual = countAliveNeighbours(grid, 2, 0)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursRightEdge(t *testing.T) {
	grid := newGrid(5, 1, 4)
	actual := countAliveNeighbours(grid, 2, 4)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 4, 1, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3, 3, 3)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 4, 1, 3, 2, 3, 3, 3, 3, 4)
	actual = countAliveNeighbours(grid, 2, 4)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}
}

func TestCountAliveNeighboursMiddle(t *testing.T) {
	grid := newGrid(5, 1, 2)
	actual := countAliveNeighbours(grid, 2, 2)

	if actual != 1 {
		t.Errorf("Expected neighbour count to be 1, but got %d", actual)
	}
	grid = newGrid(5, 1, 2, 1, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 2 {
		t.Errorf("Expected neighbour count to be 2, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 3 {
		t.Errorf("Expected neighbour count to be 3, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 4 {
		t.Errorf("Expected neighbour count to be 4, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 5 {
		t.Errorf("Expected neighbour count to be 5, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 6 {
		t.Errorf("Expected neighbour count to be 6, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 7 {
		t.Errorf("Expected neighbour count to be 7, but got %d", actual)
	}

	grid = newGrid(5, 1, 2, 1, 3, 2, 3, 3, 3, 3, 2, 3, 1, 2, 1, 1, 1)
	actual = countAliveNeighbours(grid, 2, 2)
	if actual != 8 {
		t.Errorf("Expected neighbour count to be 8, but got %d", actual)
	}
}

func TestRule1(t *testing.T) { //underpopulation
	grid := newGrid(4, 0, 0, 3, 3)
	actual := runGeneration(grid)
	expected := newGrid(4)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v but got: %v", expected, actual)
	}

}

func TestRule2(t *testing.T) { //survival
	grid := newGrid(4, 1, 1, 1, 2, 1, 3, 3, 3)
	actual := runGeneration(grid)
	expected := newGrid(4, 1, 2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v but got: %v", expected, actual)
	}
}

func TestRule3(t *testing.T) { //overpopulation
	grid := newGrid(4, 1, 1, 1, 2, 1, 3, 2, 1, 2, 2, 2, 3)
	actual := runGeneration(grid)
	expected := newGrid(4, 1, 1, 1, 3, 2, 1, 2, 3)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v but got: %v", expected, actual)
	}
}
