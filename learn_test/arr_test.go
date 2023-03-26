package learn

import (
	"strings"
	"testing"
)

func TestArr1(t *testing.T) {
	arr := []string{"a", "b"}
	for i, v := range arr {
		println(i, v)
	}
}

func TestArr2(t *testing.T) {
	pointArr := []struct{ x, y int32 }{
		{x: 1, y: 4},
	}

	for _, p := range pointArr {
		println(p.x, p.y)
	}

	str := "123"
	i := str[0:strings.Index(str, "2")]
	println(i)
	var b = strings.Builder{}
	b.Grow(9)
	b.WriteString("qweer")
	println(b.String())
}
