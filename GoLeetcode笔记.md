[TOC]

## 数组

#### 283 Move Zeros

- 第一种解法：遍历一边数组，将不为0的append进另一个数组。再遍历另一个数组，把另一个数组的值依次放入原始数组，剩余长度补0

  O(N) O(N)

```Go
func moveZeroes(nums []int) {
   retArr := make([]int, len(nums))
   index := 0
   for i := 0; i < len(nums); i++ {
      if nums[i] != 0 {
         retArr[index] = nums[i]
         index++
      }
   }
   for j := 0; j < len(nums); j++ {
      nums[j] = retArr[j]
   }
}
```

- 第二种解法：将不为0往前放,放到k的位置，最后从k往后补0

  O(N) O(1)

```go
func moveZeroes(nums []int) {
   retArr := make([]int, 0, len(nums))
   for i := 0; i < len(nums); i++ {
      if nums[i] != 0 {
         retArr = append(retArr, nums[i])
      }
   }
   for j := 0; j < len(retArr); j++ {
      nums[j] = retArr[j]
   }
   for k := len(retArr); k < len(nums); k++ {
      nums[k] = 0
   }
}
```

- 第三种解法：k和i 交换位置 当nums[i] != 0的时候，交换位置。k++

  O(N) O(1)

```go
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
```

#### 27 移除元素

```go
func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return -1
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	for j := index; j < len(nums); j++ {
		nums[j] = 0
	}
	return index
}

```



#### 26 删除排序数组中的重复项

#### 80 删除排序数组中的重复项



