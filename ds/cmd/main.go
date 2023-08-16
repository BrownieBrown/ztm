package main

import (
	"fmt"
	"ztm/ds/pkg/heaps"
)

func main() {
	m := &heaps.MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 11, 14, 20, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
