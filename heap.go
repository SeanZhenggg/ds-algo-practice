package main

import (
	"fmt"
)

type heapNode struct {
	Key   string
	Value int
}

type MinHeap struct {
	Nodes []heapNode
}

func (h *MinHeap) add(key string, value int) {
	h.Nodes = append(h.Nodes, heapNode{Key: key, Value: value})
	index := len(h.Nodes) - 1
	parent := (index - 1) / 2
	for index > 0 && h.Nodes[index].Value < h.Nodes[parent].Value {
		h.Nodes[index], h.Nodes[parent] = h.Nodes[parent], h.Nodes[index]
		index = parent
		parent /= 2
	}
}

func (h *MinHeap) delete() heapNode {
	index := 0
	top := h.Nodes[index]
	h.Nodes[index] = h.Nodes[len(h.Nodes)-1]
	h.Nodes = h.Nodes[:len(h.Nodes)-1]

	parent := index
	for index < len(h.Nodes) {
		left := 2*index + 1
		right := 2*index + 2

		if left < len(h.Nodes) && h.Nodes[index].Value > h.Nodes[left].Value {
			index = left
		}

		if right < len(h.Nodes) && h.Nodes[index].Value > h.Nodes[right].Value {
			index = right
		}

		if h.Nodes[parent].Value != h.Nodes[index].Value {
			h.Nodes[parent], h.Nodes[index] = h.Nodes[index], h.Nodes[parent]
			parent = index
		}
	}

	return top
}

func main() {
	TestAdd()
}

func TestAdd() {
	heap := MinHeap{}

	heap.add("4", 4)
	heap.add("3", 3)
	heap.add("5", 5)
	heap.add("2", 2)
	heap.add("1", 1)

	fmt.Println(heap.Nodes)
}
