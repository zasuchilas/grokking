package simple

import "errors"

func Search(list []int, item int) (int, error) {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}
