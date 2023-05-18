package learn

import (
	"context"
	"testing"
	"time"
)

/**
context.Context 一般用来同步多个 goroutine 中的状态。同一个 context 会一层层传递到所有  goroutine 中
当某个  goroutine 触发  context 结束时，所有的 goroutine 都会停止自己的工作

https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

*/

func TestContext(t *testing.T) {
	// 创建一个会过期的 context。时间到之后，context.Done 返回的 channel 会被写入值
	// 在监听这个  channel 的 goroutine 就会知道
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancelFunc()
	// 新开 goroutine, 花费一定时间处理
	go handler(ctx, 3000*time.Millisecond)

	select {
	case <-ctx.Done():
		println("main time out", ctx.Err())
	}

	time.Sleep(5 * time.Second)
}

func handler(ctx context.Context, timeout time.Duration) {
	select {
	// 新开一个 context,则与原先的 context 不关联，原先的 context  done 后，不影响当前的执行
	// 如果监听传过来的  context，则会受影响
	case <-context.Background().Done():
		println("ctx done", ctx.Err())
	case <-time.After(timeout):
		println("process success")
	}
}
