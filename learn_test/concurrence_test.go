package learn

import (
	"learn-go/utils"
	"math"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	const num = 4000
	doLockSequential(num)
	doLockConcurrentAndWaitAll(num)
	doLockConcurrentWithoutWaitAll(num)
}

func doLockSequential(cnt int) {
	var sharedNum = 0
	var lock sync.Mutex

	var increase = func() {
		defer lock.Unlock()
		sharedNum++
	}

	for i := 0; i < cnt; i++ {
		// 相当于同步执行
		lock.Lock()
		go increase()
	}
	// 最后一个 goroutine 还没结束就打印了，所以会少一个
	println(sharedNum)
}

func doLockConcurrentAndWaitAll(cnt int) {
	var sharedNum = 0
	var wg sync.WaitGroup
	var lock sync.Mutex

	var increase = func() {
		lock.Lock()
		defer lock.Unlock()
		sharedNum++
		wg.Done()
	}

	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go increase()
	}
	wg.Wait()
	println(sharedNum)
}

func doLockConcurrentWithoutWaitAll(cnt int) {
	var sharedNum = 0
	var lock sync.Mutex

	var increase = func() {
		lock.Lock()
		defer lock.Unlock()
		sharedNum++
	}

	for i := 0; i <= cnt; i++ {
		go increase()
	}
	// 未等待所有 goroutine 执行完成就打印，所以结果会一直变
	// 当 goroutine 很少的时候，几乎会一直打印 0
	println(sharedNum)
}

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
