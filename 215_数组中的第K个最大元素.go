package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 构造大顶堆
//type intHeap []int
//
//func (h *intHeap) Len() int {
//	return len(*h)
//}
//
//func (h intHeap) Less(i, j int) bool {
//	return h[i] > h[j]
//}
//func (h intHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h *intHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//func (h *intHeap) Pop() interface{} {
//	x := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return x
//}

//func findKthLargest(nums []int, k int) int {
//	if len(nums) < 1 || len(nums) < k {
//		return -1
//	}
//	h := intHeap{}
//	heap.Init(&h)
//	for _, i := range nums {
//		heap.Push(&h, i)
//	}
//	for j := 0; j < k-1; j++ {
//		heap.Pop(&h)
//	}
//	x := heap.Pop(&h)
//	return x.(int)
//}

func testFindKthLargest() {
	arr1 := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(arr1, 2))
}

func findKthLargest(arr []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(arr, 0, len(arr)-1, len(arr)-k)
}

func quickSelect(arr []int, l, r, k int) int {
	index := randomSelect(arr, l, r)
	if index == k {
		return arr[index]
	} else if index < k {
		return quickSelect(arr, index+1, r, k)
	}
	return quickSelect(arr, l, index-1, k)
}

func randomSelect(arr []int, l, r int) int {
	tag := rand.Int()%(r-l+1) + l
	arr[tag], arr[r] = arr[r], arr[tag]
	return find(arr, l, r)
}

func find(arr []int, l, r int) int {
	x := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] < x {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
