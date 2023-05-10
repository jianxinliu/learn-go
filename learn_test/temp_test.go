package learn

import "testing"

func TestDo(t *testing.T) {
	do(900000, 3, 1)
	do(900000, 3, 2)
	do(900000, 3, 3)
}

func do(amount, totalCount, curCount int) {
	println((amount / totalCount) * curCount)
}
