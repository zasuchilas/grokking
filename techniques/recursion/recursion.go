package recursion

import "fmt"

func Countdown(i int) {
	fmt.Println(i)
	if i <= 0 {
		// base case
		return
	} else {
		// recursive case
		Countdown(i - 1)
	}
}

func Fact(x int) int {
	if x == 1 {
		return 1
	}
	return x * Fact(x-1)
}
