package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitPair(input string) ([]int, []int) {
	pairs := strings.Split(input, ",")
	pair1 := strings.Split(pairs[0], "-")
	pair2 := strings.Split(pairs[1], "-")
	intPair1 := make([]int, 2)
	intPair1[0], _ = strconv.Atoi(pair1[0])
	intPair1[1], _ = strconv.Atoi(pair1[1])
	intPair2 := make([]int, 2)
	intPair2[0], _ = strconv.Atoi(pair2[0])
	intPair2[1], _ = strconv.Atoi(pair2[1])

	var slice1 []int
	for i := intPair1[0]; i <= intPair1[1]; i++ {
		slice1 = append(slice1, i)
	}
	var slice2 []int
	for i := intPair2[0]; i <= intPair2[1]; i++ {
		slice2 = append(slice2, i)
	}
	if len(slice1) > len(slice2) {
		return slice1, slice2
	} else {
		return slice2, slice1
	}
}

func PairCheck(input string) bool {
	p1, p2 := splitPair(input)

	count := 0
	for _, i := range p1 {
		for _, j := range p2 {
			if i == j {
				count++
				break
			}
		}
	}
	if count == len(p2) {
		return true
	} else {
		return false
	}

}

func OverlapCheck(input string) bool {
	p1, p2 := splitPair(input)

	count := 0
	for _, i := range p1 {
		for _, j := range p2 {
			if i == j {
				count++
				break
			}
		}
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

func QueueChecks(check func(input string) bool, input []string) int {
	res := 0
	inputChan := make(chan bool, len(input))
	done := make(chan bool)
	go func() {
		for {
			j, more := <-inputChan
			if more {
				if j {
					res++
				}
			} else {
				done <- true
				return
			}
		}
	}()
	for _, v := range input {
		if v != "" {
			inputChan <- check(v)
		}
	}
	fmt.Println("Queued", len(input), "checks")
	close(inputChan)
	<-done
	return res
}
