package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestCoord(t *testing.T) {
	input := []int{1, 2, -3, 3, -2, 0, 4}
	_, zero := mix(input)
	//printCircle(zero)
	testVal := findValAfter(zero, 0, 7)
	if testVal != 0 {
		t.Error("Expected 0, got", testVal)
	}
	testVal = findValAfter(zero, len(input), 7)
	if testVal != 0 {
		t.Error("Expected 0, got", testVal)
	}
	testVal = findValAfter(zero, 1, 7)
	if testVal != 3 {
		t.Error("Expected 3, got", testVal)
	}
	testVal = findValAfter(zero, 2, 7)
	if testVal != -2 {
		t.Error("Expected -2, got", testVal)
	}
	testVal = findValAfter(zero, 3, 7)
	if testVal != 1 {
		t.Error("Expected 1, got", testVal)
	}
	testVal = findValAfter(zero, 1000, 7)
	if testVal != 4 {
		t.Error("Expected 4, got", testVal)
	}
	testVal = findValAfter(zero, 2000, 7)
	if testVal != -3 {
		t.Error("Expected 1, got", testVal)
	}
	testVal = findValAfter(zero, 3000, 7)
	if testVal != 2 {
		t.Error("Expected 1, got", testVal)
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
	intInput := make([]int, len(input))
	for i, v := range input {
		intInput[i], _ = strconv.Atoi(v)
	}
	//_, p := mix(intInput)
	//printCircle(p)
	a, b, c := Day22(intInput)
	if a != 4 || b != -3 || c != 2 {
		t.Error("Expected 4,3,2 got", a, b, c)
	}
}

func TestInput(t *testing.T) {
	rawInput, _ := os.ReadFile("input")
	input := strings.Split(string(rawInput), "\n")
	var intInput []int
	for _, v := range input {
		if v != "" {
			temp, err := strconv.Atoi(v)
			if err != nil {
				t.Fatal(err)
			}
			intInput = append(intInput, temp)
		}
	}
	a, b, c := Day22(intInput)
	fmt.Println("Part 1")
	fmt.Println(a, b, c)
	fmt.Println("Sum:", a+b+c)
}

func TestExample2(t *testing.T) {
	rawInput := `1
2
-3
3
-2
0
4
`
	input := strings.Split(rawInput, "\n")
	var intInput []int
	for _, v := range input {
		if v != "" {
			temp, err := strconv.Atoi(v)
			if err != nil {
				t.Fatal(err)
			}
			intInput = append(intInput, temp)
		}
	}
	a, b, c, _ := Day22Part2(intInput, 10, 811589153)
	//printCircle(p)
	if a != 811589153 || b != 2434767459 || c != -1623178306 {
		t.Error("Expected something else, got", a, b, c)
	}
}

func TestInputPart2(t *testing.T) {
	rawInput, _ := os.ReadFile("input")
	input := strings.Split(string(rawInput), "\n")
	var intInput []int
	for _, v := range input {
		if v != "" {
			temp, err := strconv.Atoi(v)
			if err != nil {
				t.Fatal(err)
			}
			intInput = append(intInput, temp)
		}
	}
	fmt.Println("Part 2")
	a, b, c, _ := Day22Part2(intInput, 10, 811589153)
	fmt.Println(a, b, c)
	fmt.Println("Sum:", a+b+c)

}
