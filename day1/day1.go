package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maxInSet(nums []string) int {
	total := 0
	for _, num := range nums {
		numInt, _ := strconv.Atoi(num)
		total += numInt
	}
	return total
}

func Solution(input string) int {
	sets := strings.Split(input, "\n\n")
	fmt.Println("Sets:", len(sets))
	currentMax := 0
	for _, set := range sets {
		temp := maxInSet(strings.Split(set, "\n"))
		if temp > currentMax {
			currentMax = temp
		}
	}
	return currentMax
}

func SolutionOfThree(input string) int {
	sets := strings.Split(input, "\n\n")
	fmt.Println("Sets:", len(sets))
	var currentMax [3]int
	for _, set := range sets {
		temp := maxInSet(strings.Split(set, "\n"))
		if temp > currentMax[2] {
			if temp > currentMax[1] {
				if temp > currentMax[0] {
					currentMax[2] = currentMax[1]
					currentMax[1] = currentMax[0]
					currentMax[0] = temp
				} else {
					currentMax[2] = currentMax[1]
					currentMax[1] = temp
				}
			} else {
				currentMax[2] = temp
			}
		}
	}
	return currentMax[0] + currentMax[1] + currentMax[2]
}

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	answer := Solution(string(content))
	fmt.Println(answer)
	answer2 := SolutionOfThree(string(content))
	fmt.Println(answer2)
}
