package main

import "fmt"

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	if m == 0 {
//		copy(nums1, nums2)
//		fmt.Println(nums1)
//		return
//	}
//	temp := make([]int, m+n)
//	copy(temp, nums1)
//	index, i, j := 0, 0, 0
//
//	for i < m && j < n {
//		if nums1[i] < nums2[j] {
//			temp[index] = nums1[i]
//			i++
//			index++
//		} else {
//			temp[index] = nums2[j]
//			j++
//			index++
//		}
//	}
//	for j < n {
//		temp[index] = nums2[j]
//		j++
//		index++
//	}
//	for i < m {
//		temp[index] = nums1[i]
//		i++
//		index++
//	}
//	copy(nums1, temp)
//}

func mergeTail(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	tail := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[tail] = nums1[index1]
			tail--
			index1--
		} else {
			nums1[tail] = nums2[index2]
			tail--
			index2--
		}
	}
	for tail >= 0 && index2 >= 0 {
		nums1[tail] = nums2[index2]
		index2--
		tail--
	}
	fmt.Println(nums1)
}
func testMerge() {
	arr1 := []int{1}
	arr2 := []int{1}
	mergeTail(arr1, 0, arr2, 1)
}
