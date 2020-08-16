/*
	Node Depths

  The distance between a node in a Binary Tree and the tree's root is called the
	node's depth. Write a function that takes in a Binary Tree and returns the sum of its nodes'
  depths.
*/

package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type Level struct {
	Node  *BinaryTree
	Depth int
}

func NodeDepthsIterative(root *BinaryTree) int {
	totalDepth := 0
	stack := []Level{{Node: root, Depth: 0}}

	var top Level
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if top.Node == nil {
			continue
		}

		totalDepth += top.Depth

		stack = append(stack, Level{Node: top.Node.Left, Depth: top.Depth + 1})
		stack = append(stack, Level{Node: top.Node.Right, Depth: top.Depth + 1})
	}

	return totalDepth
}

func NodeDepthsRecursive(root *BinaryTree) int {
	return computeNodeDepths(root, 0)
}

func computeNodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}

	return depth + computeNodeDepths(node.Right, depth+1) + computeNodeDepths(node.Left, depth+1)
}
