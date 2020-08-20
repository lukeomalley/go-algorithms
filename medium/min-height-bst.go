/*
   Min Height BST

	  Write a function that takes in a non-empty sorted array of distinct integers,
		constructs a BST from the integers, and returns the root of the BST.

		The function sould minimize the height of the BST
*/

package main

import "fmt"

func MinHeightBST(array []int) *BST {
	return constructMinHeightBst(array, nil, 0, len(array)-1)
}

func constructMinHeightBst(array []int, bst *BST, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return bst
	}

	midIdx := +((startIdx + endIdx) / 2)
	valueToAdd := array[midIdx]
	if bst == nil {
		bst = &BST{Value: valueToAdd}
	} else {
		bst.Insert(valueToAdd)
	}

	constructMinHeightBst(array, bst, startIdx, midIdx-1)
	constructMinHeightBst(array, bst, midIdx+1, endIdx)
	return bst
}

// Improved version operates in O(n) time as we avoid using the insert method
func constructMinHeightBstImproved(array []int, bst *BST, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return bst
	}

	midIdx := +((startIdx + endIdx) / 2)
	newBstNode := &BST{Value: array[midIdx]}

	if bst == nil {
		bst = newBstNode
	} else {
		if newBstNode.Value < bst.Value {
			// insert to left
			bst.Left = newBstNode
			bst = bst.Left
		} else {
			// insert to right
			bst.Right = newBstNode
			bst = bst.Right
		}
	}

	constructMinHeightBstImproved(array, bst, startIdx, midIdx-1)
	constructMinHeightBstImproved(array, bst, midIdx+1, endIdx)
	return bst
}

// This is the same as the second method, only a much cleaner version
func constructMinHeightBstClean(array []int, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}

	midIdx := +((startIdx + endIdx) / 2)

	bst := &BST{Value: array[midIdx]}
	bst.Left = constructMinHeightBstClean(array, startIdx, midIdx-1)
	bst.Right = constructMinHeightBstClean(array, midIdx+1, endIdx)
	return bst
}

// ============================================================================
// BST IMPLEMENTATION
// ============================================================================

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func main() {
	fmt.Println(MinHeightBST([]int{1, 2, 5, 7, 10, 13, 14, 15, 22}).Value)
}
