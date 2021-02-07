package main

func RemoveDuplicates2(nums []int) int {
	i:=0
	for j:=0; j < len(nums);j++ {
		if i < 2 || nums[j] > nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func testRemoveDuplicates2() {
	arr := []int{1, 1, 1, 2, 2, 3}
	RemoveDuplicates2(arr)
}
