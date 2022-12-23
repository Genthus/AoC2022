package day8

import (
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `30373
25512
65332
33549
35390
`
	res := Solve(rawInput)
	if res != 21 {
		t.Error("Expected 21, got", res)
	}
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := Solve(string(rawInput))
	fmt.Println("Part1 result:", res)
}
