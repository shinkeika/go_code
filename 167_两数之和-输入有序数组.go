package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{-1, -1}
	}

	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l = l + 1
		} else {
			r = r - 1
		}
	}
	return []int{-1, -1}
}

func testTwoSum() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
