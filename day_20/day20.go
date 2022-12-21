package main

func ProcessString(nums []int) []int {
	inputMap := make(map[int]int)
	for i, val := range nums {
		inputMap[i] = val
	}
}

func find(s []int, f int) int {
	for i, val := range s {
		if val == f {
			return i
		}
	}
	return -1
}

func remove(s []int, del int) ([]int, int) {
	for i, val := range s {
		if val == del {
			return append(s[:i], s[i+1:]...), i
		}
	}
	return s, -1
}

func insert(s []int, ins int, index int) []int {
	index = index % len(s)
	if index <= 0 {
		index = len(s) + index
	}
	return append(s[:index], append([]int{ins}, s[index:]...)...)
}

func ShiftVal(val int, nums *[]int) {
	newNums, index := remove(*nums, val)
	*nums = insert(newNums, val, index+val)
}

func GetCoordinates(nums []int) (int, int, int) {
	placeof0 := find(nums, 0)
	placeafter0 := func(placesAfer int) int {
		return nums[(placeof0+placesAfer)%len(nums)]
	}

	return placeafter0(1000), placeafter0(2000), placeafter0(3000)
}
