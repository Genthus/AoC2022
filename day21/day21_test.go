package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	rawInput := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`
	input := strings.Split(rawInput, "\n")
	res := QueueChecks(input)
	if res != 152 {
		t.Error("Expected 152, got", res)
	}
}
func TestExample2(t *testing.T) {
	rawInput := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`
	input := strings.Split(rawInput, "\n")
	res := MonkeyBrain(input)
	if res != 301 {
		t.Error("Expected 301, got", res)
	}
}

func TestInput(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	res := QueueChecks(input)
	fmt.Println("Part 1 Result:", res)
}

func TestInputPart2(t *testing.T) {
	rawInput, err := os.ReadFile("input")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Split(string(rawInput), "\n")
	res := MonkeyBrain(input)
	fmt.Println("Part 2 Result:", res)
}
