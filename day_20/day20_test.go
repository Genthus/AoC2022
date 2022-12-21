package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	input := []int{1, 2, -3, 4, 0, 3, -2}
	res := find(input, -3)
	if res != 2 {
		t.Error("Expected 4, got", res)
	}
}
func TestRemove(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, -8, -1}
	res, _ := remove(input, 3)
	expected := []int{1, 2, 4, 5, 6, -8, -1}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected", expected[i], "got", res[i], "in result", res)
		}
	}
}

func TestInsert(t *testing.T) {
	input := []int{1, 2, -3, 4, 0, 3, -2}
	res := insert(input, 50, 4)
	expected := []int{1, 2, -3, 4, 50, 0, 3, -2}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected", expected[i], "got", res[i], "in result", res)
		}
	}
}

func TestGetCoords(t *testing.T) {
	input := []int{1, 2, -3, 4, 0, 3, -2}
	a, b, c := GetCoordinates(input)
	if a != 4 && b != -3 && c != 2 {
		t.Error("Expected 4 -3 2, got:", a, b, c)
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
	var intInput []int
	for _, val := range input {
		if val != "" {
			temp, err := strconv.Atoi(val)
			if err != nil {
				t.Fatal(err)
			}
			intInput = append(intInput, temp)
		}
	}
	result := ProcessString(intInput)
	a, b, c := GetCoordinates(result)
	if a != 4 && b != -3 && c != 2 {
		t.Error("Expected 4 -3 2, got:", a, b, c, "in array", result)
	}
}

func TestInput(t *testing.T) {
	rawInput, _ := os.ReadFile("input")
	input := strings.Split(string(rawInput), "\n")
	var intInput []int
	for _, val := range input {
		if val != "" {
			temp, err := strconv.Atoi(val)
			if err != nil {
				t.Fatal(err)
			}
			intInput = append(intInput, temp)
		}
	}
	result := ProcessString(intInput)
	a, b, c := GetCoordinates(result)
	fmt.Println("SUM", a+b+c)
}
