package learn

// https://tonybai.com/2022/03/25/intro-generics/
import (
	"testing"
)

type Number interface {
	int | float32 | float64
}

func TestGeneric(t *testing.T) {
	// 泛型实例化，变成非泛型函数。只支持泛型方法声明时的类型
	f := To[float64]
	println(f(4))

	// 也可以直接使用，并且不提供类型参数，自己推断
	to := To(3.4)
	println(to)
}

// To 泛型支持
// 声明泛型类型形参，支持联合类型
func To[T int | float32 | float64](s T) T {
	return s
}
