package benchutil

import "github.com/zasuchilas/grokking/utils/listutil"

var (
	List  = listutil.MakeOrderedIntList(1_000_000)
	Items = []int{0, 250_000, 500_000, 750_000, 1_000_000}
)
