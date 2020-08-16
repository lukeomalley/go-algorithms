/*
	Longest Peak

		Write a function that takes in an array of integers and returns the length of
		the longest peak in the array.

		A peak is defined as adjacent integers in the array that are strictly
		increasing until they reach a tip (the highest value in the peak), at which
		point they become strictly  decreasing. At least three integers are required to
		form a peak.
*/

package main

func LongestPeak(array []int) int {
	longestPeak := 0
	i := 1
	for i < len(array)-1 {
		// PHASE 1: Find a peak
		isPeak := array[i-1] < array[i] && array[i+1] < array[i]

		if !isPeak {
			i++
			continue
		}

		// PHASE 2: Expand around the peak and find its length
		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx--
		}

		rightIdx := i + 2
		for rightIdx <= len(array)-1 && array[rightIdx] < array[rightIdx-1] {
			rightIdx++
		}

		peakLength := rightIdx - leftIdx - 1
		if peakLength > longestPeak {
			longestPeak = peakLength
		}

		i++
	}

	return longestPeak
}
