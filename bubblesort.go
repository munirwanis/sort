package sort

// BubbleSort rearranges the array to it's sequential format.
func BubbleSort(array []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i+1] < array[i] {
				tmp := array[i+1]
				array[i+1] = array[i]
				array[i] = tmp
				swapped = true
			}
		}
	}
}
