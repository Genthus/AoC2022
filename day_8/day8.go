package day8

import (
	"log"
	"strconv"
	"strings"
)

func formGrid(input string) [][]int {
	splitInput := strings.Split(input, "\n")
	multiSlice := make([][]int, len(splitInput)-1)
	for i, v := range splitInput {
		if v != "" {
			multiSlice[i] = make([]int, len(v))
			for j, v2 := range v {
				intVal, err := strconv.Atoi(string(v2))
				if err != nil {
					log.Fatal(err)
				}
				multiSlice[i][j] = intVal
			}
		}
	}
	return multiSlice
}

func checkView(x, y int, input [][]int) bool {
	treeVal := input[x][y]
	checkDirection := func(dirX, dirY int) bool {
		cursorX, cursorY := x+dirX, y+dirY
		for cursorX >= 0 &&
			cursorY >= 0 &&
			cursorX < len(input) &&
			cursorY < len(input[0]) {
			if input[cursorX][cursorY] >= treeVal {
				return false
			}
			cursorX += dirX
			cursorY += dirY
		}
		return true
	}

	if checkDirection(1, 0) ||
		checkDirection(-1, 0) ||
		checkDirection(0, 1) ||
		checkDirection(0, -1) {
		return true
	}

	return false
}

func checkForest(input [][]int) int {
	inputLenght := len(input) * len(input[0])
	jobs := make(chan [2]int, inputLenght)
	done := make(chan int)

	go func() {
		res := 0
		for {
			j, more := <-jobs
			if more {
				if checkView(j[0], j[1], input) {
					res++
				}
			} else {
				done <- res
			}
		}
	}()

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			jobs <- [2]int{i, j}
		}

	}
	close(jobs)

	res := <-done
	return res
}

func Solve(input string) int {
	grid := formGrid(input)
	res := checkForest(grid)
	return res
}
