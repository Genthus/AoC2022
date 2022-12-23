package main

func findFirst4Unique(input string) int {
	count := 4
	for len(input) > 3 {
		commons := 0
		for i, v := range input[:4] {
			for j, v2 := range input[:4] {
				if i != j && v == v2 {
					commons++
				}
			}
		}
		if commons > 0 {
			input = input[1:]
			count++
		} else {
			return count
		}
	}
	return 0
}

func findFirstNUnique(input string, n int) int {
	count := n
	for len(input) > n {
		commons := 0
		for i, v := range input[:n] {
			for j, v2 := range input[:n] {
				if i != j && v == v2 {
					commons++
				}
			}
		}
		if commons > 0 {
			input = input[1:]
			count++
		} else {
			return count
		}
	}
	return 0
}

func Solve(input string) int {
	return findFirst4Unique(input)
}

func SolvePart2(input string) int {
	return findFirstNUnique(input, 14)
}
