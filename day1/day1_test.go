package main

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	// t.Fatal("not implemented")
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	sol := Solution(input)
	fmt.Println(sol)
	if sol != 24000 {
		t.Error("Solution incorrect, expected 24000, got {}", sol)
	}
}
