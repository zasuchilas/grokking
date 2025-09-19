package selection

// Sort sorts the items using a selection sort - O(n * n).
func Sort(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := range arr {
		idx := findSmallest(arr)
		newArr[i] = arr[idx]

		// removing the added element from the source array
		arr = append(arr[:idx], arr[idx+1:]...) // arr = slices.Delete(arr, idx, idx+1)
		// TODO: it turns out that we are spoiling the incoming array
	}
	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
