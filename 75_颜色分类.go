package main

func sortColors(nums []int) {
	var count [3]int
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	index := 0
	for j := 0; j < len(count); j++ {
		for k := 0; k < count[j]; k++ {
			nums[index] = j
			index++
		}
	}
}

func testSortColors() {
	arr := []int{2, 1, 1, 2, 0, 0}
	sortColors(arr)
}
