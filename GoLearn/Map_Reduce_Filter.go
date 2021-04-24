package GoLearn

import (
	"fmt"
)


func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}
var list = []string{"Hao", "Chen", "MegaEase"}

func TestReduce()  {
	x := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", x)
	// 15
}

