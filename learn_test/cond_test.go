package learn

import (
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/#cond

// 共享变量，也是控制并发的条件
var shared int32 = 0

func TestCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	// 创建多个 goroutine, 分别监听共享变量为 1 和 2 的情况
	for i := 0; i < 4; i++ {
		go listenFor1(cond)
		go listenFor2(cond)
	}
	time.Sleep(5 * time.Second)
	// 此 goroutine 随机决定先赋值 1 还是 2
	go broadcastRnd(cond)

	// 监听中断信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcastRnd(c *sync.Cond) {

	rand.Seed(time.Now().Unix())
	f := rand.Float32()

	n1, n2 := int32(1), int32(2)
	if f > 0.5 {
		n1, n2 = int32(2), int32(1)
	}
	// 通知第一个变量
	c.L.Lock()
	atomic.StoreInt32(&shared, n1)
	c.Broadcast()
	c.L.Unlock()

	time.Sleep(1 * time.Second)

	// 通知第二个变量
	c.L.Lock()
	atomic.StoreInt32(&shared, n2)
	c.Broadcast()
	c.L.Unlock()
}

func listenFor1(c *sync.Cond) {
	c.L.Lock()
	// 自旋监听共享变量的值，不是 1 就一直等
	for atomic.LoadInt32(&shared) != 1 {
		c.Wait()
	}
	println("listen1")
	c.L.Unlock()
}

func listenFor2(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt32(&shared) != 2 {
		c.Wait()
	}
	println("listen2")
	c.L.Unlock()
}
