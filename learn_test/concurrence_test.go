package learn

import (
	"learn-go/utils"
	"math"
	"testing"
	"time"
)

func TestConcurrence(t *testing.T) {
	doTest(100_000, 20000)
	doTest(100, 3)
}

func doTest(max, step int) {
	chanList := Mapper(BigNumSum, step, max)
	ret := Reducer(chanList)
	println(utils.Format("[{}, {}] => {}", max, step, ret))
}

func Mapper(fn func(start, max int, c chan int64), step int, max int) []chan int64 {
	chanList := make([]chan int64, int(math.Ceil(float64(max)/float64(step))))
	j := 0
	for i := 0; ; i += step {
		chanList[j] = make(chan int64, 1)
		curMax := math.Min(float64(i+step), float64(max))

		go fn(i, int(curMax), chanList[j])
		if curMax == float64(max) {
			break
		}
		j += 1
	}
	return chanList
}

func Reducer(chanList []chan int64) int64 {
	ret := int64(0)

	// buffered channels need to be closed before iterating over them.
	// 所以改用 unbuffered channel
	for _, c := range chanList {
		select {
		case r := <-c:
			ret += r
		case <-time.After(time.Duration(5) * time.Second):
			ret += 0
		}
	}
	return ret
}

func BigNumSum(start int, max int, retChan chan int64) {
	ret := 0
	for i := start + 1; i <= max; i++ {
		ret += i
	}
	retChan <- int64(ret)
}
