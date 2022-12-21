package day5

import (
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	res := Solve(rawInput)
	if res != "CMZ" {
		t.Error("Expected CMZ, got", res)
	}
}
func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := Solve(string(rawInput))
	fmt.Println("Result part1:", res)
}

func TestExample2(t *testing.T) {
	rawInput := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	res := SolvePart2(rawInput)
	if res != "MCD" {
		t.Error("Expected MCD, got", res)
	}
}
func TestInputPart2(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := SolvePart2(string(rawInput))
	fmt.Println("Result part2:", res)
}
