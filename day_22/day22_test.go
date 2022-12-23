package main

import (
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`
	res := Solve(rawInput)
	if res != 6032 {
		t.Error("Expected 6032, got", res)
	}
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := Solve(string(rawInput))
	fmt.Println("Part1 res:", res)
}
