package main

func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 如果 i 和 index 相等 那没必要交换，直接++
			if i != index {
				swap(nums, i, index)
				index++
			} else {
				index++
			}
		}
	}
}

func swap(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func testMoveZeroes() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
}
