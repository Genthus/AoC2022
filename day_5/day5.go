package day5

import (
	"strconv"
	"strings"
)

func parseStacks(input string) map[int][]string {
	inputRows := strings.Split(input, "\n")
	numberOfColumns := (len(inputRows[0]) + 1) / 4
	columnMap := make(map[int][]string)
	for i := 0; i < numberOfColumns; i++ {
		columnStack := []string{}
		for j, v := range inputRows {
			if j == len(inputRows)-1 {
				break
			}
			if v[i*4+1] != ' ' {
				columnStack = append(columnStack, string(v[i*4+1]))
			}
		}
		columnMap[i+1] = columnStack
	}
	return columnMap
}

func parseInstructions(input string) [][3]int {
	inputRows := strings.Split(input, "\n")
	instructions := make([][3]int, 0)
	for _, v := range inputRows {
		vSplit := strings.Split(v, " ")
		if len(vSplit) == 6 {
			v1, _ := strconv.Atoi(vSplit[1])
			v2, _ := strconv.Atoi(vSplit[3])
			v3, _ := strconv.Atoi(vSplit[5])
			instructions = append(instructions, [3]int{v1, v2, v3})
		}
	}
	return instructions
}

func parseInput(input string) (map[int][]string, [][3]int) {
	splitInput := strings.Split(input, "\n\n")
	inputCrates := splitInput[0]
	inputInstructions := splitInput[1]
	stacks := parseStacks(inputCrates)
	instructions := parseInstructions(inputInstructions)
	return stacks, instructions
}

func reorderCrates(stacks map[int][]string, instructions [][3]int) map[int][]string {
	for _, v := range instructions {
		for i := 0; i < v[0]; i++ {
			stack := stacks[v[1]]
			crate := stack[0]
			stacks[v[1]] = stacks[v[1]][1:]
			stacks[v[2]] = append([]string{crate}, stacks[v[2]]...)
		}
	}
	return stacks
}
func reorderCrates9001(stacks map[int][]string, instructions [][3]int) map[int][]string {
	for _, v := range instructions {
		stack := stacks[v[1]]
		crate := make([]string, v[0]) // memory management is funny sometimes
		copy(crate, stack[:v[0]])
		stacks[v[1]] = stacks[v[1]][v[0]:]
		stacks[v[2]] = append(crate, stacks[v[2]]...)
	}
	return stacks
}

func readTopCrates(stacks map[int][]string) string {
	res := ""
	for i := 0; i < len(stacks); i++ {
		res = res + stacks[i+1][0]
	}
	return res
}

func Solve(input string) string {
	stacks, instructions := parseInput(input)
	resultStacks := reorderCrates(stacks, instructions)
	return readTopCrates(resultStacks)
}
func SolvePart2(input string) string {
	stacks, instructions := parseInput(input)
	resultStacks := reorderCrates9001(stacks, instructions)
	return readTopCrates(resultStacks)
}
