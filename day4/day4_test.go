package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`
	input := strings.Split(rawInput, "\n")
	res := QueueChecks(PairCheck, input)

	if res != 2 {
		t.Error("Expected 2, got", res)
	}
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	res := QueueChecks(PairCheck, input)
	fmt.Println("Result:", res)
}
func TestOverlaps(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	res := QueueChecks(OverlapCheck, input)
	fmt.Println("Result:", res)
}
