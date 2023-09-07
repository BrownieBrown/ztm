package challenges

import (
	"container/heap"
	"math"
)

type Point struct {
	x, y     int
	distance float64
}

type PointsHeap []Point

func (h PointsHeap) Len() int           { return len(h) }
func (h PointsHeap) Less(i, j int) bool { return h[i].distance > h[j].distance } // max-heap
func (h PointsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointsHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *PointsHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func calculateDistance(x, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

func kClosest(points [][]int, k int) [][]int {
	h := &PointsHeap{}

	for _, point := range points {
		distance := calculateDistance(point[0], point[1])
		heap.Push(h, Point{point[0], point[1], distance})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		result[i] = []int{p.x, p.y}
	}

	return result
}
