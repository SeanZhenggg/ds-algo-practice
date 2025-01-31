package main

import (
	"fmt"
)

type node struct {
	Key    string
	Weight int
}

type minPriorityQueue struct {
	heap    []*node
	indexes map[string]int // key -> index
}

func NewMinPriorityQueue(list []node) *minPriorityQueue {
	minPQ := &minPriorityQueue{
		heap:    make([]*node, len(list)+1),
		indexes: make(map[string]int, len(list)),
	}

	for i := 0; i < len(list); i++ {
		minPQ.heap[i+1] = &list[i]
		minPQ.setIndex(list[i].Key, i+1)
	}

	minPQ.build()

	return minPQ
}

func (pq *minPriorityQueue) getParent(idx int) *node {
	return pq.heap[idx/2]
}

func (pq *minPriorityQueue) getIndex(key string) int {
	return pq.indexes[key]
}

func (pq *minPriorityQueue) setIndex(key string, idx int) {
	pq.indexes[key] = idx
}

func (pq *minPriorityQueue) heapify(n *node) {
	length := pq.getLength()
	index := pq.getIndex(n.Key)

	minIndex := index
	if 2*index < length && pq.heap[minIndex].Weight > pq.heap[2*index].Weight {
		minIndex = 2 * index
	}

	if 2*index+1 < length && pq.heap[minIndex].Weight > pq.heap[2*index+1].Weight {
		minIndex = 2*index + 1
	}

	pq.swap(minIndex, index)

	if minIndex != index {
		pq.heapify(pq.heap[minIndex])
	}
}

func (pq *minPriorityQueue) swap(idx1 int, idx2 int) {
	pq.setIndex(pq.heap[idx1].Key, idx2)
	pq.setIndex(pq.heap[idx2].Key, idx1)
	pq.heap[idx1], pq.heap[idx2] = pq.heap[idx2], pq.heap[idx1]
}

func (pq *minPriorityQueue) getLength() int {
	return len(pq.heap)
}

func (pq *minPriorityQueue) build() {
	for i := len(pq.heap) / 2; i >= 1; i-- {
		pq.heapify(pq.heap[i])
	}
}

func (pq *minPriorityQueue) poll() *node {
	minNode := pq.heap[1]

	pq.setIndex(pq.heap[len(pq.heap)-1].Key, 1)
	pq.heap[1] = pq.heap[len(pq.heap)-1]

	pq.heap = pq.heap[:len(pq.heap)-1]

	pq.heapify(pq.heap[1])

	return minNode
}

func (pq *minPriorityQueue) push(key string, weight int) {
	newNode := &node{Key: key, Weight: weight}
	pq.heap = append(pq.heap, newNode)
	pq.decrease(key, weight)
}

func (pq *minPriorityQueue) decrease(key string, weight int) {
	index := pq.getIndex(key)
	if pq.heap[index].Weight < weight {
		return
	}

	pq.heap[index].Weight = weight

	for index > 1 && pq.getParent(index).Weight > pq.heap[index].Weight {
		pq.swap(index/2, index)
		index /= 2
	}
}

func (pq *minPriorityQueue) print() {
	for i := 1; i <= len(pq.heap)-1; i++ {
		fmt.Printf("index: %d, key = %s, weight = %d\n", i, pq.heap[i].Key, pq.heap[i].Weight)
	}
}

type PriorityQueue struct {
	heap        []interface{}
	compareFunc func(a, b interface{}) int
}

func NewPriorityQueue(compareFunc func(a, b interface{}) int) *PriorityQueue {
	pq := &PriorityQueue{
		heap:        make([]interface{}, 0),
		compareFunc: compareFunc,
	}

	return pq
}

func (pq *PriorityQueue) build() {
	for i := len(pq.heap) / 2; i >= 1; i-- {
		pq.heapify(i)
	}
}

func (pq *PriorityQueue) heapify(index int) {
	length := pq.getLength()

	minIndex := index
	l := 2*index + 1
	r := 2*index + 2
	if l < length && pq.compareFunc(pq.heap[minIndex], pq.heap[l]) > 0 {
		minIndex = l
	}

	if r < length && pq.compareFunc(pq.heap[minIndex], pq.heap[r]) > 0 {
		minIndex = r
	}

	if minIndex != index {
		pq.swap(minIndex, index)
		pq.heapify(minIndex)
	}
}

func (pq *PriorityQueue) getLength() int {
	return len(pq.heap)
}

func (pq *PriorityQueue) swap(idx1 int, idx2 int) {
	pq.heap[idx1], pq.heap[idx2] = pq.heap[idx2], pq.heap[idx1]
}

func (pq *PriorityQueue) poll() interface{} {
	minIndex := 0
	minNode := pq.heap[minIndex]

	pq.heap[minIndex] = pq.heap[len(pq.heap)-1]

	pq.heap = pq.heap[:len(pq.heap)-1]

	pq.heapify(minIndex)

	return minNode
}

func (pq *PriorityQueue) top() interface{} {
	return pq.heap[0]
}

func (pq *PriorityQueue) push(n interface{}) {
	pq.heap = append(pq.heap, n)
	index := len(pq.heap) - 1
	for index > 0 && pq.compareFunc(pq.heap[(index-1)/2], pq.heap[index]) > 0 {
		pq.swap((index-1)/2, index)
		index = (index - 1) / 2
	}
}

func (pq *PriorityQueue) print() {
	for i := 0; i < len(pq.heap); i++ {
		fmt.Printf("index = %d, val = %v\n", i, pq.heap[i])
	}
}

func main() {
	//pq := NewMinPriorityQueue([]node{
	//	{"B", 22},
	//	{"E", 20},
	//	{"H", 15},
	//	{"C", 12},
	//	{"I", 10},
	//	{"G", 8},
	//	{"A", 7},
	//	{"F", 4},
	//	{"D", 2},
	//})

	//pq.print()

	//fmt.Println("============================================")
	//
	//pq.decrease("C", 3)
	//
	//pq.print()
	//
	//fmt.Println("============================================")
	//
	//pq.poll()
	//
	//pq.print()

	pq := NewPriorityQueue(func(a, b interface{}) int {
		return a.([2]int)[0] - b.([2]int)[0]
	})

	pq.push([2]int{22, 1})
	pq.push([2]int{20, 2})
	pq.push([2]int{15, 3})
	pq.push([2]int{12, 4})
	pq.push([2]int{10, 5})
	pq.push([2]int{8, 6})
	pq.push([2]int{7, 7})
	pq.push([2]int{4, 8})
	pq.push([2]int{2, 9})

	pq.print()
	fmt.Println("============================================")

	pq.push([2]int{3, 10})

	pq.print()

	fmt.Println("============================================")

	pq.poll()

	pq.print()
}
