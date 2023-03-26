package learn_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(*testing.T) {
	// 创建一个错误
	e := errors.New("好一个错误")
	// 使用 %w 将错误进行包装
	wrapper := fmt.Errorf("出错啦：%w", e)
	// 包装后的错误
	fmt.Println(wrapper)

	w2 := fmt.Errorf("xx: %w", wrapper)
	w3 := fmt.Errorf("xx: %w", w2)
	fmt.Println(w2)
	fmt.Println(w3)

	// 也可以将包装后的错误还原，但是只能还原一层
	fmt.Println(errors.Unwrap(wrapper))
	fmt.Println(errors.Unwrap(w3))
}
