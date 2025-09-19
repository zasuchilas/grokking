package listutil

func MakeOrderedIntList(length int) []int {
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = i
	}
	return res
}
