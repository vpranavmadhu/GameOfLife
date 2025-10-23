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
