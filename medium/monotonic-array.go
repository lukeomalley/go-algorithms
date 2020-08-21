/*
	Monotonic Array

	Write a function that takes in an array of integers and returns a boolean representing
	whether the array is monotonic

	An array is said to be monotonic if its elements from left to right are entirely non
	increasing or entirely non decreasing
*/

package main

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}

	direction := array[1] - array[0] // Negative if decreasing

	for i := 2; i < len(array); i++ {
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}

		if breaksDirection(direction, array[i-1], array[i]) {
			return false
		}
	}

	return true
}

func breaksDirection(direction, prevInt, currentInt int) bool {
	currentDir := currentInt - prevInt
	if direction > 0 {
		return currentDir < 0
	}

	return currentDir > 0
}
