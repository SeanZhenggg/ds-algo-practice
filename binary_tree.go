package main

import (
	"log"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		nodes = append(nodes, root)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return nodes
}

// 4
// 4 out, 7 2 in -> 7 2
// 2 out, 3 1 in -> 7 3 1
// 1 out -> 7 3
// 3 out -> 7
// 7 out, 9 6 in -> 9 6
// 6 out -> 9
// 9 out

func IterativePreorderTraversal(root *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	st := make([]*TreeNode, 0)
	st = append(st, root)

	for len(st) > 0 {
		top := st[len(st)-1]
		nodes = append(nodes, top)
		st = st[:len(st)-1]

		if top.Right != nil {
			st = append(st, top.Right)
		}

		if top.Left != nil {
			st = append(st, top.Left)
		}
	}

	return nodes
}

func InorderTraversal(root *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nodes = append(nodes, root)
		dfs(root.Right)
	}
	dfs(root)
	return nodes
}

func IterativeInorderTraversal(root *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	st := make([]*TreeNode, 0)

	node := root
	for len(st) > 0 || node != nil {
		if node != nil {
			st = append(st, node)
			node = node.Left
		} else {
			node = st[len(st)-1]
			st = st[:len(st)-1]
			nodes = append(nodes, node)
			node = node.Right
		}
	}
	return nodes
}

func PostorderTraversal(root *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		nodes = append(nodes, root)
	}
	dfs(root)
	return nodes
}

func IterativePostorderTraversal(root *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)
	st := make([]*TreeNode, 0)

	current := root           // 記錄當前走到的節點
	var lastVisited *TreeNode // 記錄最近一個走過的節點
	for len(st) > 0 || current != nil {
		if current != nil { // 遍歷節點
			st = append(st, current) // 將節點放入 stack 中
			current = current.Left   // 往左遍歷
		} else {
			top := st[len(st)-1].Right            // stack 頂端的節點
			if top != nil && top != lastVisited { // 判斷其有沒有右邊的子節點，而且這個右邊的子節點不是最近走過的
				current = top // 往右遍歷
			} else {
				// 沒有的話則將頂端的節點放至結果陣列，並且pop出頂端的節點
				top = st[len(st)-1]
				st = st[:len(st)-1]
				result = append(result, top)
				lastVisited = top // 記錄最近一個走過的節點
			}
		}
	}

	return result
}

func IterativePostorderTraversal2(root *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)
	st := make([]*TreeNode, 0)

	current := root
	for len(st) > 0 || current != nil {
		if current != nil {
			st = append(st, current)
			current = current.Left
		} else {
			topRight := st[len(st)-1].Right

			if topRight != nil {
				current = topRight
			} else {
				top := st[len(st)-1]
				st = st[:len(st)-1]
				result = append(result, top)

				// 如果遇到堆疊頂端的節點其右子節點是當前的 top，代表頂端節點的右邊子樹已經遍歷完了，可以連續彈出
				for len(st) > 0 && st[len(st)-1].Right == top {
					top = st[len(st)-1]
					st = st[:len(st)-1]
					result = append(result, top)
				}
			}
		}
	}
	return result
}

//	  4
//	2	7
// 1 3 6 9

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
	}

	//preTraversals := PreorderTraversal(root)
	//printTree(preTraversals)
	//preTraversals = IterativePreorderTraversal(root)
	//printTree(preTraversals)
	//
	//inTraversals := InorderTraversal(root)
	//printTree(inTraversals)
	//inTraversals = IterativeInorderTraversal(root)
	//printTree(inTraversals)
	//
	//postTraversals := PostorderTraversal(root)
	//printTree(postTraversals)
	//postTraversals = IterativePostorderTraversal(root)
	//printTree(postTraversals)
	postTraversals := IterativePostorderTraversal2(root)
	printTree(postTraversals)
}

func printTree(nodes []*TreeNode) {
	builder := strings.Builder{}
	for i, treeNode := range nodes {
		builder.WriteString(strconv.Itoa(treeNode.Val))
		if i != len(nodes)-1 {
			builder.WriteString(",")
		}
	}
	log.Println(builder.String())
}
