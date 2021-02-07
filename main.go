package main

type LabelTree struct {
	Title    string      `json:"title"`
	Value    string      `json:"value"`
	Label    string      `json:"label"`
	ParentId int         `json:"parent_id"`
	Id       int         `json:"id"`
	Children []LabelTree `json:"children"`
}

func main() {
	//testBinarySearch()g
	//testMoveZeroes()
	//testRemoveElement()
	//testRemoveDuplicates2()
	//bank.Func1()
	//testSortColors()
	//testMerge()
	//testFindKthLargest()
	testTwoSum()
	testLongestPalindrome()
}

//var _ Shape = (*Square)(nil)
//
//type Shape interface {
//	Sides() int
//	Area() int
//}
//
//type Square struct {
//	len int
//}
//func (s* Square) Sides() int {
//	return 4
//}
//func (s* Square) Area() int {
//	return 4
//}
//
