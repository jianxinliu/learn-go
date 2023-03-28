package learn

// https://tonybai.com/2022/03/25/intro-generics/
import (
	"fmt"
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

func Uniq[T comparable](collection []T) []T {
	innerMap := make(map[T]string, len(collection))
	ret := make([]T, 0, len(collection))

	for _, n := range collection {
		if _, ok := innerMap[n]; ok {
			continue
		}
		innerMap[n] = ""
		ret = append(ret, n)
	}
	return ret
}

func printArr[T comparable](list []T) {
	for _, n := range list {
		print(n)
		print(" ")
	}
	println()
}

func TestUniq(t *testing.T) {
	ints := []int{3, 2, 3, 3, 2, 3, 4, 3, 2, 3, 2}
	printArr(Uniq(ints))

	floats := []float32{.2, .43, .4, .4, .2, .43}
	printArr(Uniq(floats))

	strs := []string{"s", "a", "s", "v"}
	printArr(Uniq(strs))
}

func GroupBy[K comparable, V any](list []V, keyFn func(i V) K) map[K][]V {
	ret := map[K][]V{}
	for _, n := range list {
		fn := keyFn(n)
		ret[fn] = append(ret[fn], n)
	}
	return ret
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) ToString() string {
	return fmt.Sprintf(" [%d, %d]", c.X, c.Y)
}

func TestGroupBy(t *testing.T) {
	points := []Coordinate{
		{1, 2},
		{1, 3},
		{2, 5},
		{3, 6},
		{2, 4},
	}
	grp := GroupBy(points, func(i Coordinate) int {
		return i.X
	})
	for k, v := range grp {
		print(k)
		for _, n := range v {
			print(n.ToString())
			print(" ")
		}
		println()
	}
}
