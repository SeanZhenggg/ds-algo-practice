package main

import "fmt"

type UnionFind struct {
	set []int
}

// 1. New Disjoint Set
func NewUnionFind(length int) *UnionFind {
	set := make([]int, 0, length)
	for i := 0; i < length; i++ {
		set = append(set, -1)
	}
	return &UnionFind{
		set: set,
	}
}

// 2. FindRoot method
func (u *UnionFind) FindRoot(idx int) int {
	var root = idx
	for u.set[root] >= 0 {
		root = u.set[root]
	}

	return root
}

// 3. Union method
func (u *UnionFind) Union(idx1 int, idx2 int) {
	idx1 = u.FindRoot(idx1)
	idx2 = u.FindRoot(idx2)
	if u.set[idx1] <= u.set[idx2] {
		temp := u.set[idx2]
		u.set[idx2] = idx1
		u.set[idx1] += temp
	} else {
		temp := u.set[idx1]
		u.set[idx1] = idx2
		u.set[idx2] += temp
	}
}

func (u *UnionFind) PrintSet() {
	for i := range u.set {
		fmt.Printf("% 3d\t", i)
	}
	fmt.Print("\n")
	for _, j := range u.set {
		fmt.Printf("% 3d\t", j)
	}
	fmt.Print("\n\n")
}

func main() {
	uf := NewUnionFind(6)
	uf.Union(0, 1)

	uf.PrintSet()

	uf.Union(2, 3)

	uf.PrintSet()

	uf.Union(5, 4)

	uf.PrintSet()

	uf.Union(2, 4)

	uf.PrintSet()

	uf.Union(1, 4)

	uf.PrintSet()
}
