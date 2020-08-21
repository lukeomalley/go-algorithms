/*
	Spiral Traverse

	Spiral order starts at the top left corner of the two-dimensional array, goes
  to the right, and proceeds in a spiral pattern all the way until every element
  has been visited.

  Write a function that takes in an n x m two-dimensional array (that can be
  square-shaped when n == m) and returns a one-dimensional array of all the
  array's elements in spiral order.
*/

package main

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	result := []int{}
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		// Top Edge
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}

		// Right Edge
		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}

		// Bottom Edge
		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}

			result = append(result, array[endRow][col])
		}

		// Left Eddge
		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}

			result = append(result, array[row][startCol])
		}

		startCol++
		startRow++
		endCol--
		endRow--
	}

	return result
}
