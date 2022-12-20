package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestHelloWorld2(t *testing.T) {
	input := `A Y
B X
C Z`
	sum := 0
	for _, val := range strings.Split(input, "\n") {
		s := strings.Split(val, " ")
		sum += Optimal(s[0], s[1])
		fmt.Println("CurrentScore:", sum)
	}
	if sum != 12 {
		t.Error("Expected 12, got", sum)
	} else {
		fmt.Println("Score:", sum)
	}
}

func TestInput2(t *testing.T) {
	input, _ := os.ReadFile("input")
	sum := 0
	for _, val := range strings.Split(string(input), "\n") {
		s := strings.Split(val, " ")
		if len(s) == 2 {
			sum += Optimal(s[0], s[1])
		}
	}
	fmt.Println("Score:", sum)
}
func TestHelloWorld(t *testing.T) {
	input := `A Y
B X
C Z`
	sum := 0
	for _, val := range strings.Split(input, "\n") {
		s := strings.Split(val, " ")
		sum += ScoreCalc(s[0], s[1])
	}
	if sum != 15 {
		t.Error("Expected 15, got", sum)
	} else {
		fmt.Println("Score:", sum)
	}
}

func TestInput(t *testing.T) {
	input, _ := os.ReadFile("input")
	sum := 0
	for _, val := range strings.Split(string(input), "\n") {
		s := strings.Split(val, " ")
		if len(s) == 2 {
			sum += ScoreCalc(s[0], s[1])
		}
	}
	fmt.Println("Score:", sum)
}
