package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestGetCoords(t *testing.T) {
	input := []int{1, 2, -3, 4, 0, 3, -2}
	a, b, c := GetCoordinates(input)
	if a != 4 && b != -3 && c != 3 {
		t.Error("Expected 4 -3 3, got:", a, b, c)
	}
}

func TestExample(t *testing.T) {
	rawInput := `1
2
-3
3
-2
0
4`
	input := strings.Split(rawInput, "\n")
	intInput := make([]int, 0)
	for _, val := range input {
		temp, err := strconv.Atoi(val)
		if err != nil {
			t.Fatal(err)
		}
		intInput = append(intInput, temp)
	}
	result := ProcessString(intInput)
	a, b, c := GetCoordinates(result)
	if a != 4 && b != -3 && c != 3 {
		t.Error("Expected 4 -3 3, got:", a, b, c)
	}
}
