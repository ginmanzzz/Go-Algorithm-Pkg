package algo

import (
	"fmt"
)

/*
	@brief:
*/
type Heap struct {
	data []interface{}
	isMaxHeap bool
}

func (h *Heap) up(index int) {
	if !h.chkIdx(index) {
		fmt.Println("Invalid index<{}>", index)
		return 
	}
}

func (h *Heap) down(index int) {
	if !h.chkIdx(index) {
		fmt.Println("Invalid index<{}>", index)
		return 
	}
}

func (h *Heap) Push(elem interface{}) {
}

func (h *Heap) Pop() (interface{}) {
	if (len(h.data) == 0) {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) chkIdx(index int ) bool {
	if index < 0 || index > len(h.data) {
		return false
	}
	return true
}

func less(a, b interface{}) bool {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a < b
		}
	case int64:
		if b, ok := b.(int64); ok {
			return a < b
		}
	case string:
		if b, ok := b.(string); ok {
			return a < b
		}
	default:
		fmt.Println("Unsupported type.")
		return false
	}
	return false
}
