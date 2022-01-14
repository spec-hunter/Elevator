package handlers

func countSort(arr []int) []int {
	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	array := make([]int, max-min+1)
	for _, value := range arr {
		array[value-min]++
	}

	for i := 1; i < len(array); i++ {
		array[i] += array[i-1]
	}

	sortArray := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sortArray[array[arr[i]-min]-1] = arr[i]
		array[arr[i]-min]--
	}

	return sortArray
}
