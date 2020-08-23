/*
	Max Path Sums

	Write a function that takes in a Binary Tree and returns its max path sum.

	A path is a collection of connected nodes in a tree where no node is connected
  to more than two other nodes; a path sum is the sum of the values of the nodes
  in a particular path.
*/

package main

import "math"

// BinaryTree represents a binary tree where each node has a max of two children nodes
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// MaxPathSum returns the path down the tree with the max value
func MaxPathSum(tree *BinaryTree) int {
	_, maxSum := findMaxSums(tree)
	return maxSum
}

func findMaxSums(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, math.MinInt32
	}

	leftMaxSumsAsBranch, leftMaxPathSum := findMaxSums(tree.Left)
	rightMaxSumsAsBranch, rightMaxPathSum := findMaxSums(tree.Right)
	maxChildSumAsBranch := max(leftMaxSumsAsBranch, rightMaxSumsAsBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumsAsBranch+value+rightMaxSumsAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)

	return maxSumAsBranch, maxPathSum
}

func max(first int, vals ...int) int {
	for _, val := range vals {
		if val > first {
			first = val
		}
	}

	return first
}
