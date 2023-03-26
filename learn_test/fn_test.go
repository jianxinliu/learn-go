package learn_test

import (
	"testing"
)

func ifNotNilThen(o interface{}, fn func(o interface{})) {
	if o != nil {
		fn(o)
	} else {
		println("no.....")
	}
}

func TestFn(t *testing.T) {
	fn := func(o interface{}) {
		if v, ok := o.(string); ok {
			println(v)
		}
		if v, ok := o.(int); ok {
			println(v)
		}

	}
	ifNotNilThen(nil, fn)
	ifNotNilThen("aaa", fn)
	ifNotNilThen(23, fn)
}

// 自定义类型
type MyFn func(int) int

// Adder 创建一个初始累加器
func Adder(a int) MyFn {
	return func(b int) int {
		return a + b
	}
}

func TestFnFn(t *testing.T) {
	adder := Adder(1)
	println(adder(3))
	println(adder(4))
}
