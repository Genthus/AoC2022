package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func arrangeMap(input string) [][]string {
	rows := strings.Split(input, "\n")
	longest := 0
	for _, v := range rows {
		if len(v) > longest {
			longest = len(v)
		}
	}
	fmt.Println("longest", longest)
	grid := make([][]string, len(rows))
	for i, v := range rows {
		grid[i] = make([]string, longest)
		for j := 0; j < longest; j++ {
			if j < len(v) {
				grid[i][j] = string(v[j])
			} else {
				grid[i][j] = " "
			}
		}
		for j, v2 := range v {
			grid[i][j] = string(v2)
		}
	}
	return grid

}

func printGrid(input [][]string, x, y int) {
	for i, v := range input {
		for j, v2 := range v {
			if i == y && j == x {
				fmt.Print("@")
			} else {
				fmt.Print(v2)
			}
		}
		fmt.Println(i)
	}
}

func parseInstructions(input string) []string {
	res := []string{}
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "R", " R ", -1)
	input = strings.Replace(input, "L", " L ", -1)
	res = strings.Split(input, " ")
	return res
}

func move(input [][]string, x int, y int, direction int, ammount int) (int, int) {
	fmt.Println("starting in", x, y, direction, ammount)
	determineDir := func(d int) (int, int) {
		switch d {
		case 270:
			return 0, -1
		case 90:
			return 0, +1
		case 0:
			return 1, 0
		case 180:
			return -1, 0
		}
		log.Fatal("Unknown direction")
		return 0, 0
	}

	directionX, directionY := determineDir(direction)
	newPosY, newPosX := y+directionY, x+directionX
	for i := 0; i < ammount; i++ {
		if newPosY < 0 {
			newPosY = len(input) - 1
		} else if newPosY > len(input)-1 {
			newPosY = 0
		} else if newPosX < 0 {
			newPosX = len(input[newPosY]) - 1
		} else if newPosX > len(input[newPosY])-1 {
			newPosX = 0
		}
		switch input[newPosY][newPosX] {
		case ".":
			y = newPosY
			x = newPosX
		case "#":
			return x, y
		case " ":
			i--
		}
		newPosY += directionY
		newPosX += directionX
	}
	//printGrid(input, x, y)
	return x, y
}

func initialPos(grid [][]string) int {
	for i, v := range grid[0] {
		if string(v) == "." {
			return i
		}
	}
	log.Fatal("unable to find starting position")
	return 0
}

func TankControls(grid [][]string, instructions []string) int {
	direction := 0
	x := initialPos(grid)
	y := 0
	for len(instructions) > 0 {
		if instructions[0] == "L" {
			direction -= 90
			if direction < 0 {
				direction = 360 + direction
			}
		} else if instructions[0] == "R" {
			direction += 90
			direction %= 360
		} else {
			if instructions[0] != "" {
			}
			times, err := strconv.Atoi(instructions[0])
			if err != nil {
				log.Fatal("error reading times to move")
				return 0
			}
			x, y = move(grid, x, y, direction, times)
		}
		instructions = instructions[1:]
	}

	// scoring
	rowScore := 1000 * (y + 1)
	colScore := 4 * (x + 1)
	facingScore := direction / 90
	return rowScore + colScore + facingScore
}

func Solve(input string) int {
	splitInput := strings.Split(input, "\n\n")
	inputMap := arrangeMap(splitInput[0])
	instructions := parseInstructions(splitInput[1])
	printGrid(inputMap, 0, 0)
	fmt.Println(instructions)
	res := TankControls(inputMap, instructions)
	return res
}
