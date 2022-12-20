package main

func ProcessString(nums []int) []int {
	return nums
}

func GetCoordinates(nums []int) (int, int, int) {
	a := nums[1000%len(nums)-3]
	b := nums[2000%len(nums)-3]
	c := nums[3000%len(nums)-3]
	return a, b, c
}
