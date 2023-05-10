package handler

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	doLock1(40)
	doLock2(40)
}

func doLock1(cnt int) {
	var sharedNum = 0

	var increase = func(lock *sync.Mutex) {
		defer lock.Unlock()
		sharedNum++
	}

	var lock sync.Mutex
	for i := 0; i <= cnt; i++ {
		lock.Lock()
		go increase(&lock)
	}
	println(sharedNum)
}

func doLock2(cnt int) {
	var sharedNum = 0

	var increase = func(lock *sync.Mutex, wg *sync.WaitGroup) {
		lock.Lock()
		defer lock.Unlock()
		sharedNum++
		wg.Done()
	}

	// why?
	var wg = sync.WaitGroup{}
	var lock sync.Mutex
	for i := 0; i <= cnt; i++ {
		wg.Add(1)
		go increase(&lock, &wg)
	}
	wg.Wait()
	println(sharedNum)
}
