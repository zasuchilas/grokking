package quick

// Sort sorts the items using a quick sort - O(n * log n).
func Sort(list []int) []int {
	// base case
	if len(list) < 2 {
		return list
	}

	// recursive case
	pivot := list[0]
	var less, greater []int
	for _, item := range list[1:] {
		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	res := append(Sort(less), pivot)
	res = append(res, Sort(greater)...)
	return res
}
