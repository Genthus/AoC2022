package main

import (
	"fmt"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	test := func(i string, e int) {
		res := Solve(i)
		if res != e {
			t.Error("Expected", e, "got", res)
		}
	}
	test("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7)
	test("bvwbjplbgvbhsrlpgdmjqwftvncz", 5)
	test("nppdvjthqldpwncqszvftbrmjlhg", 6)
	test("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10)
	test("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11)
}

func TestExamplePart2(t *testing.T) {
	test := func(i string, e int) {
		res := SolvePart2(i)
		if res != e {
			t.Error("Expected", e, "got", res)
		}
	}
	test("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
	test("bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
	test("nppdvjthqldpwncqszvftbrmjlhg", 23)
	test("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
	test("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := Solve(string(rawInput))
	fmt.Println("Part1 result", res)
}

func TestInputPart2(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	res := SolvePart2(string(rawInput))
	fmt.Println("Part2 result", res)
}
