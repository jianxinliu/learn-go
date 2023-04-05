package learn

import (
	"testing"
)

func TestDeadLock(t *testing.T) {
	doTest2(100_000, 2000)
}

func doTest2(max, step int) {
	chanList := make(chan int64, max/step+1)
	Mapper1(BigNumSum, step, max, chanList)
	ret := int64(0)

	// chan unclosed can not be ranged
	// cause deadlock
	for c := range chanList {
		ret += c
	}
	println(ret)
}

func Mapper1(fn func(start, max int, c chan int64), width int, max int, c chan int64) {
	j := 0
	for i := 0; i < max; i += width {
		//println(utils.Format("[{}, {}]", i, i+width))
		go fn(i, i+width, c)
		j += 1
	}
	// 将多出来的 chan push 一个 0
	curJ := j - 1
	for i := 0; i < cap(c)-curJ; i++ {
		c <- 0
	}
}
