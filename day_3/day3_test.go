package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	input := strings.Split(rawInput, "\n")
	// t.Fatal("not implemented")
	sum := 0
	for _, val := range input {
		if len(val)%2 == 0 {
			common := FindCommon(val)
			sum += ValueRune(common)
		}
	}

	if sum != 157 {
		t.Error("Expected 157, got", sum)
	}
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	sum := 0
	for _, val := range input {
		if len(val)%2 == 0 {
			common := FindCommon(val)
			sum += ValueRune(common)
		}
	}
	fmt.Println("Input Result:", sum)
}

func TestExample2(t *testing.T) {
	rawInput := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	input := strings.Split(string(rawInput), "\n")
	sum := 0
	for i := 0; i < len(input)-2; i += 3 {
		common := FindCommons(input[i], input[i+1])
		common = FindCommons(common, input[i+2])
		sum += ValueRune(rune(common[0]))
	}
	if sum != 70 {
		t.Error("Expected 70, got", sum)
	}
}
func TestInput2(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	sum := 0
	for i := 0; i < len(input)-2; i += 3 {
		common := FindCommons(input[i], input[i+1])
		common = FindCommons(common, input[i+2])
		sum += ValueRune(rune(common[0]))
	}
	fmt.Println("Input2 Result:", sum)
}
