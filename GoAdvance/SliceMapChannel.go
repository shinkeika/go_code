package GoAdvance

import "fmt"

//func Slice() {
//	a := make([]int, 100)
//	b := [100]int{}
//	fmt.Println(a, b)
//}
//
//func PrintSlice(s []int) {
//	fmt.Println(len(s))
//	fmt.Println(cap(s))
//}
//
//func PrintArray(s [100]int) {
//	fmt.Println(len(s))
//	fmt.Println(cap(s))
//}
//
////go:noinline
//func callSlice(s []int) {
//}
//
//// go:noinline
//func callArray(s [10000]int) {
//}

//func BenchmarkCallSlice(b *testing.B) {
//	s := make([]int, 10000)
//	for i := 0; i < b.N; i++ {
//		callSlice(s)
//	}
//}
//
//func BenchmarkCallArray(b *testing.B) {
//	var a [10000]int
//	for i := 0; i < b.N; i++ {
//		callArray(a)
//	}
//}

// go:noinline
//func LenAndCap() {
	var s []int
	//fmt.Println(len(s), cap(s))
	//s = append(s, 0)
	//// go:noinline
	//s = append(s, 1)
	//fmt.Println(len(s), cap(s))
	//
	//s = append(s, 2)
	//fmt.Println(len(s), cap(s))
	//
	//
	//for i := 3; i < 1025; i++ {
	//	s = append(s, i)
	//}
	//fmt.Println(len(s), cap(s))

//	s = append(s, 0, 1, 2)
//	fmt.Println(len(s), cap(s))
//}

func Chan() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		default:
			break
		}
	}
}

//func Test() {
//N := 100
//from := make([]int, 100, 100)
//to := make([]int, 100, 100)
//for i := 0; i < 100; i++ {
//	from = append(from, i)
//	to = append(to, i)
//}
////fmt.Println(from, to)
//for i := 0; i < N; {
//	to[i] = from[i]
//	i++
//	to[i] = from[i]
//	i++
//}
//n := (N + 7) / 8
//i := 0
//switch N % 8 {
//case 0:
//	for {
//		to[i] = from[i]
//		i++
//case 1:
//	to[i] = from[i]
//	i++
//case 2:
//	to[i] = from[i]
//	i++
//case 3:
//	to[i] = from[i]
//	i++
//case 4:
//	to[i] = from[i]
//	i++
//case 5:
//	to[i] = from[i]
//	i++
//case 6:
//	to[i] = from[i]
//	i++
//case 7:
//	to[i] = from[i]
//	i++
//	if i < N {
//		break
//	}
//}
//}
