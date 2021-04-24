package GoAdvance

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	_ "unsafe" // for linkname
)

//const (
//	isDeleted = 1
//	unDeleted = 0
//)
//
//type IntList struct {
//	head   *intNode
//	length int64
//}
//
//type intNode struct {
//	value     int
//	next      *intNode
//	mu        sync.Mutex
//	isDeleted uint32
//}
//
//func newIntNode(value int) *intNode {
//	return &intNode{value: value}
//}
//
//func (n *intNode) getValue() int {
//	return int(atomic.LoadInt64((*int64)(unsafe.Pointer(&n.value))))
//}
//
//func (n *intNode) setValue(value int) {
//	atomic.StoreInt64((*int64)(unsafe.Pointer(&n.value)), int64(value))
//}
//
//func (n *intNode) getNextNode() *intNode {
//	return (*intNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&n.next))))
//}
//
//func (n *intNode) setNextNode(next *intNode) {
//	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&n.next)), unsafe.Pointer(next))
//}
//
//func NewInt() *IntList {
//	return &IntList{head: newIntNode(0)}
//}
//
//func (l *IntList) Insert(value int) bool {
//	var a, b *intNode
//	for {
//		a = l.head
//		b = a.getNextNode()
//		for b != nil && b.getValue() < value {
//			a = b
//			b = b.getNextNode()
//		}
//		if b != nil && b.value == value {
//			return false
//		}
//
//		a.mu.Lock()
//		// double check
//		if a.getNextNode() == b {
//			break
//		} else {
//			a.mu.Unlock()
//			continue
//		}
//	}
//
//	x := newIntNode(value)
//	x.next = b
//	a.setNextNode(x)
//
//	atomic.AddInt64(&l.length, 1)
//	a.mu.Unlock()
//	return true
//}
//
//func (l *IntList) Delete(value int) bool {
//	var a, b *intNode
//	for {
//		a = l.head
//		b = a.getNextNode()
//		for b != nil && b.getValue() < value {
//			a = b
//			b = b.getNextNode()
//		}
//
//		if b == nil || b.value != value {
//			return false
//		}
//
//		b.mu.Lock()
//		if atomic.LoadUint32(&b.isDeleted) == isDeleted {
//			b.mu.Unlock()
//			// 已经被删除
//			return false
//		}
//
//		// 检查a有没有变化
//		a.mu.Lock()
//		if a.getNextNode() != b || atomic.LoadUint32(&a.isDeleted) == isDeleted {
//			a.mu.Unlock()
//			b.mu.Unlock()
//			continue
//		} else {
//			// b为所寻找的值
//			break
//		}
//	}
//	atomic.StoreUint32(&b.isDeleted, isDeleted)
//	a.setNextNode(b.getNextNode())
//	atomic.AddInt64(&l.length, -1)
//	a.mu.Unlock()
//	b.mu.Unlock()
//	return true
//}
//
//func (l *IntList) Contains(value int) bool {
//	if l.Len() == 0 {
//		return false
//	}
//	x := l.head.getNextNode()
//	for x != nil && x.getValue() < value {
//		x = x.getNextNode()
//	}
//	if x == nil || atomic.LoadUint32(&x.isDeleted) == isDeleted || x.value != value {
//		return false
//	}
//	return true
//}
//
//func (l *IntList) Range(f func(value int) bool) {
//	x := l.head.getNextNode()
//	for x != nil {
//		if !f(x.getValue()) {
//			break
//		}
//		x = x.getNextNode()
//	}
//}
//
//func (l *IntList) Len() int {
//	return int(atomic.LoadInt64(&l.length))
//}

//go:linkname fastrand runtime.fastrand
func fastrand() uint32

//go:nosplit
func fastrandn(n uint32) uint32 {
	return uint32(uint64(fastrand()) * uint64(n) >> 32)
}

func TestIntSet(t *testing.T) {
	// Correctness.
	l := NewInt()
	if l.Len() != 0 {
		t.Fatal("invalid length")
	}
	if l.Contains(0) {
		t.Fatal("invalid contains")
	}
	if l.Delete(0) {
		t.Fatal("invalid delete")
	}

	if !l.Insert(0) || l.Len() != 1 {
		t.Fatal("invalid insert")
	}
	if !l.Contains(0) {
		t.Fatal("invalid contains")
	}
	if !l.Delete(0) || l.Len() != 0 {
		t.Fatal("invalid delete")
	}

	if !l.Insert(20) || l.Len() != 1 {
		t.Fatal("invalid insert")
	}
	if !l.Insert(22) || l.Len() != 2 {
		t.Fatal("invalid insert")
	}
	if !l.Insert(21) || l.Len() != 3 {
		t.Fatal("invalid insert")
	}

	var i int
	l.Range(func(score int) bool {
		if i == 0 && score != 20 {
			t.Fatal("invalid range")
		}
		if i == 1 && score != 21 {
			t.Fatal("invalid range")
		}
		if i == 2 && score != 22 {
			t.Fatal("invalid range")
		}
		i++
		return true
	})

	i = 0
	l.Range(func(_ int) bool {
		i++
		return i != 2
	})
	if i != 2 {
		t.Fatal("invalid range")
	}

	if !l.Delete(21) || l.Len() != 2 {
		t.Fatal("invalid delete")
	}

	i = 0
	l.Range(func(score int) bool {
		if i == 0 && score != 20 {
			t.Fatal("invalid range")
		}
		if i == 1 && score != 22 {
			t.Fatal("invalid range")
		}
		i++
		return true
	})
	const num = 10000
	// Make rand shuffle array.
	// The testArray contains [1,num]
	testArray := make([]int, num)
	testArray[0] = num + 1
	for i := 1; i < num; i++ {
		// We left 0, because it is the default score for head and tail.
		// If we check the skiplist contains 0, there must be something wrong.
		testArray[i] = int(i)
	}
	for i := len(testArray) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := fastrandn(uint32(i + 1))
		testArray[i], testArray[j] = testArray[j], testArray[i]
	}

	// Concurrent insert.
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		i := i
		wg.Add(1)
		go func() {
			l.Insert(testArray[i])
			wg.Done()
		}()
	}
	wg.Wait()
	if l.Len() != num {
		t.Fatalf("invalid length expected %d, got %d", num, l.Len())
	}

	// Don't contains 0 after concurrent insertion.
	if l.Contains(0) {
		t.Fatal("contains 0 after concurrent insertion")
	}

	// Concurrent contains.
	for i := 0; i < num; i++ {
		i := i
		wg.Add(1)
		go func() {
			if !l.Contains(testArray[i]) {
				wg.Done()
				panic(fmt.Sprintf("insert doesn't contains %d", i))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Concurrent delete.
	for i := 0; i < num; i++ {
		i := i
		wg.Add(1)
		go func() {
			if !l.Delete(testArray[i]) {
				wg.Done()
				panic(fmt.Sprintf("can't delete %d", i))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if l.Len() != 0 {
		t.Fatalf("invalid length expected %d, got %d", 0, l.Len())
	}

	// Test all methods.
	const smallRndN = 1 << 8
	for i := 0; i < 1<<16; i++ {
		wg.Add(1)
		go func() {
			r := fastrandn(num)
			if r < 333 {
				l.Insert(int(fastrandn(smallRndN)) + 1)
			} else if r < 666 {
				l.Contains(int(fastrandn(smallRndN)) + 1)
			} else if r != 999 {
				l.Delete(int(fastrandn(smallRndN)) + 1)
			} else {
				var pre int
				l.Range(func(score int) bool {
					if score <= pre { // 0 is the default value for header and tail score
						panic("invalid content")
					}
					pre = score
					return true
				})
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// Correctness 2.
	var (
		x     = NewInt()
		y     = NewInt()
		count = 10000
	)

	for i := 0; i < count; i++ {
		x.Insert(i)
	}

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			x.Range(func(score int) bool {
				if x.Delete(score) {
					if !y.Insert(score) {
						panic("invalid insert")
					}
				}
				return true
			})
			wg.Done()
		}()
	}
	wg.Wait()
	if x.Len() != 0 || y.Len() != count {
		t.Fatal("invalid length")
	}

	// Concurrent Insert and Delete in small zone.
	x = NewInt()
	var (
		insertcount uint64 = 0
		deletecount uint64 = 0
	)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				if fastrandn(2) == 0 {
					if x.Delete(int(fastrandn(10))) {
						atomic.AddUint64(&deletecount, 1)
					}
				} else {
					if x.Insert(int(fastrandn(10))) {
						atomic.AddUint64(&insertcount, 1)
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if insertcount < deletecount {
		panic("invalid count")
	}
	if insertcount-deletecount != uint64(x.Len()) {
		panic("invalid count")
	}
}
