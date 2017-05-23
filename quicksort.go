package sort

import "math/rand"

// QuickSort rearranges the array to it's sequential format.
func QuickSort(array []int) {
	if len(array) > 2 {
		left, right := 0, len(array)-1

		// Pick a pivot
		pivotIndex := rand.Int() % len(array)

		// Move the pivot to the right
		array[pivotIndex], array[right] = array[right], array[pivotIndex]

		// Pile elements smaller than the pivot on the left
		for i := range array {
			if array[i] < array[right] {
				array[i], array[left] = array[left], array[i]
				left++
			}
		}

		// Place the pivot after the last smaller element
		array[left], array[right] = array[right], array[left]

		// Go down the rabbit hole
		QuickSort(array[:left])
		QuickSort(array[left+1:])
	}
}
