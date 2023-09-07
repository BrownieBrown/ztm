package heaps

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("cannot extract because length is zero")
		return -1
	}

	extracted := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	if len(h.array) != 0 { // Only heapify down if there are elements left in the heap.
		h.heapifyDown(0)
	}

	return extracted
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(parent int) int {
	return (2 * parent) + 1
}

func rightChild(parent int) int {
	return (2 * parent) + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
