/*

	Same BSTs


		An array of integers is said to represent the Binary Serach Tree (BST)
		obtained by inserting each integer in the array, from left to right
		into the BST.

		Write a function that takes in two arrays of integers and determines
		whether these arrays represent the same BST. Note that you're not
		allowed to construct any BSTs in your code
*/

package main

// SameBSTs validates that two arrays can construct the same BST
func SameBSTs(arrayOne, arrayTwo []int) bool {
	// Check that the lengths of the arrays are the same
	if len(arrayOne) != len(arrayTwo) {
		return false
	}

	// Check that the two arrays have some length
	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}

	// Check that the root nodes are the same
	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	leftOne := getSmaller(arrayOne)
	rightOne := getGreaterOrEqual(arrayOne)

	leftTwo := getSmaller(arrayTwo)
	rightTwo := getGreaterOrEqual(arrayTwo)

	return SameBSTs(leftOne, leftTwo) && SameBSTs(rightOne, rightTwo)
}

func getSmaller(array []int) []int {
	smaller := []int{}
	for i := 1; i < len(array); i++ {
		if array[i] < array[0] {
			smaller = append(smaller, array[i])
		}
	}

	return smaller
}

func getGreaterOrEqual(array []int) []int {
	greater := []int{}
	for i := 1; i < len(array); i++ {
		if array[i] >= array[0] {
			greater = append(greater, array[i])
		}
	}

	return greater
}
