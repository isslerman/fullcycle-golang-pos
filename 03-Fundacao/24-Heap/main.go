package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) PrintValues() {
	fmt.Printf("%v\n", h)
	// for _, v := range *h {
	// 	fmt.Printf("v: %v", v)
	// }
}

// Add adds an element to the heap.
func (h *IntHeap) Add(x int) {
	heap.Push(h, x)
}

// Remove removes and returns the smallest element from the heap.
func (h *IntHeap) Remove() int {
	return heap.Pop(h).(int)
}

// Min returns the minimum value of the heap.
func (h IntHeap) Min() int {
	if h.Len() == 0 {
		return 0 // Or any other default value you prefer
	}
	return h[0]
}

// RemoveValue removes a specific value from the heap.
func (h *IntHeap) RemoveValue(val int) {
	for i, v := range *h {
		if v == val {
			// Remove the element at index i.
			heap.Remove(h, i)
			break // Assuming you only want to remove the first occurrence.
		}
	}
}

func main() {
	var h IntHeap
	h.Add(10)
	h.Add(2)
	h.Add(15)
	h.PrintValues() // printa ordenado
	h.Pop()         // remove o ultimo (maior)
	h.PrintValues() // printa ordenado
	h.Add(15)
	h.Remove() // remove o menor (primeiro)
	h.PrintValues()
	h.Add(2)
	h.PrintValues()
	h.RemoveValue(15)
	h.PrintValues()
	fmt.Println(h.Min())
}
