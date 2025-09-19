// Package dc means Divide & conquer technique (D&C)
package dc

import "errors"

// LoopSum returns the sum of the elements using a loop.
func LoopSum(arr []int) int {
	total := 0
	for _, item := range arr {
		total += item
	}
	return total
}

// Sum returns the sum of the items using the D&C technique.
func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + Sum(list[1:])
}

// Count returns the number of elements using recursion (D&C technique).
func Count(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + Count(list[1:])
}

// Max returns the maximum element using recursion (D&C technique).
func Max(list []int) (int, error) {
	if len(list) < 2 {
		return 0, errors.New("the number of elements is less than two")
	}
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0], nil
		}
		return list[1], nil
	}
	subMax, err := Max(list[1:])
	if err != nil {
		return 0, err
	}
	if list[0] > subMax {
		return list[0], nil
	}
	return subMax, err
}

// FarmSquarePlots evenly divides the farm into as large square plots as possible.
func FarmSquarePlots(w, h int) int {
	// correcting the sides if they are mixed up
	if w < h {
		w, h = h, w
	}

	// processing invalid data
	if w == 0 || h == 0 {
		return 0
	}

	// the base case
	if w%h == 0 {
		return h
	}

	w1 := h
	h1 := w - (w/h)*h
	return FarmSquarePlots(w1, h1)
}
