package algo

import (
	"fmt"
	"math"
)

/*
	@brief: The heap is based on Completed Binary Tree, so it can be stored in an array.
			For a node, its index in the array is i.
			Its parent's index is (i - 1) / 2.
			Its left child index is 2 * i + 1, and right child index is 2 * i + 2.
*/
type Heap [T ComparableType] struct {
	data []T
	isMaxHeap bool
}

func NewHeap[T ComparableType] (isMaxHeap bool) *Heap[T] {
	return &Heap[T] {
		data: make([]T, 0),
		isMaxHeap: isMaxHeap,
	}
}

func (h *Heap[T]) up(index int) {
	if !h.chkIdx(index) {
		return 
	}
	parent := (index - 1) / 2
	if parent < 0 {
		return
	}
	if h.isMaxHeap && h.data[index] > h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		h.up(parent)
	} else if !h.isMaxHeap && h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		h.up(parent)
	}
}

func (h *Heap[T]) down(index int) {
	if !h.chkIdx(index) {
		return 
	}
	left := 2 * index + 1
	right := 2 * index + 2
	n := len(h.data)
	if (h.isMaxHeap) {
		maxIdx := index
		if left < n && h.data[maxIdx] < h.data[left] {
			maxIdx = left
		}
		if right < n && h.data[maxIdx] < h.data[right] {
			maxIdx = right
		}
		if maxIdx != index {
			h.data[maxIdx], h.data[index] = h.data[index], h.data[maxIdx]
			h.down(maxIdx)
		}
	} else {
		minIdx := index
		if left < n && h.data[minIdx] > h.data[left] {
			minIdx = left
		}
		if right < n && h.data[minIdx] > h.data[right] {
			minIdx = right
		}
		if minIdx != index {
			h.data[minIdx], h.data[index] = h.data[index], h.data[minIdx]
			h.down(minIdx)
		}
	}
}

func (h *Heap[T]) Push(elem T) {
	h.data = append(h.data, elem)
	h.up(len(h.data) - 1)
}

func (h *Heap[T]) Pop() (interface{}) {
	if (len(h.data) == 0) {
		return nil
	}
	top := h.data[0]
	n := len(h.data)
	h.data[0], h.data[n-1] = h.data[n-1], h.data[0]
	h.data = h.data[0: n-1]
	return top
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}

func (h *Heap[T]) chkIdx(index int ) bool {
	if index < 0 || index > len(h.data) {
		fmt.Println("Invalid index<{}>", index)
		return false
	}
	return true
}

func (h* Heap[T]) String() string {
	// 1, 2, 4, 8, 16 ...
	n := len(h.data)
	layer := 0
	start := 0
	str := ""
	for {
		layerNodeNum := int(math.Pow(2, (float64(layer))))
		i := start
		for ; i < n && i < start + layerNodeNum; i++ {
			fmt.Sprintf(str, " %d", h.data[i])
		}
		if i >= n {
			break
		}
		start += layerNodeNum
		layer++
		fmt.Sprintf(str, "\n")
	}
	return str
}
