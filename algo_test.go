package algo 

import (
	"testing"
	"fmt"
)

func TestMaxInt(t *testing.T) {
	a := 5
	b := 6
	c := Max(a, b)
	if c != b {
		t.Errorf("Max(5, 6) = %d, Expected: 6", c)
	}
}

func TestMaxStr(t *testing.T) {
	a := "abc"
	b := "cde"
	c := Max(a, b)
	if c != b {
		t.Errorf("Max('abc', 'cde') = %s, Expected: 'cde'", c)
	}
}

func TestHeapStr(t *testing.T) {
	heap := NewHeap[int](false)
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	fmt.Print(heap.String())
}
