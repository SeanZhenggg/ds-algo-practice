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
	nodes := make([]*TreeNode, 0)
	st := make([]*TreeNode, 0)

	node := root
	var lastNodeVisited *TreeNode = nil
	for len(st) > 0 || node != nil {
		if node != nil {
			st = append(st, node)
			node = node.Left
		} else {
			top := st[len(st)-1]

			if top.Right != nil && lastNodeVisited != top.Right {
				node = top.Right
			} else {
				nodes = append(nodes, top)
				st = st[:len(st)-1]
				lastNodeVisited = top
			}
		}
	}

	return nodes
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

	preTraversals := PreorderTraversal(root)
	printTree(preTraversals)
	preTraversals = IterativePreorderTraversal(root)
	printTree(preTraversals)

	inTraversals := InorderTraversal(root)
	printTree(inTraversals)
	inTraversals = IterativeInorderTraversal(root)
	printTree(inTraversals)

	postTraversals := PostorderTraversal(root)
	printTree(postTraversals)
	postTraversals = IterativePostorderTraversal(root)
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
