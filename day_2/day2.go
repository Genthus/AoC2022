package main

func ScoreCalc(oponent string, strategy string) int {
	points := 0
	switch strategy {
	case "X":
		points += 1
		if oponent == "A" {
			points += 3
		} else if oponent == "C" {
			points += 6
		}
	case "Y":
		points += 2
		if oponent == "B" {
			points += 3
		} else if oponent == "A" {
			points += 6
		}
	case "Z":
		points += 3
		if oponent == "C" {
			points += 3
		} else if oponent == "B" {
			points += 6
		}
	}

	return points
}

func Optimal(oponent string, strategy string) int {
	points := 0
	handVals := make(map[string]int)
	handVals["A"] = 1
	handVals["B"] = 2
	handVals["C"] = 3
	switch strategy {
	case "X":
		points += 0
		if oponent == "A" {
			points += 3
		} else {
			points += handVals[oponent] - 1
		}
	case "Y":
		points += 3
		points += handVals[oponent]
	case "Z":
		points += 6
		if oponent == "C" {
			points += 1
		} else {
			points += handVals[oponent] + 1
		}
	}

	return points
}
