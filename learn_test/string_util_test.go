package learn

import (
	"fmt"
	"learn-go/utils"
	"testing"
)

type Name struct {
	Age int
}

func (name Name) ToString() string {
	return fmt.Sprintf("name:[%d]", name.Age)
}

func TestFormat(t *testing.T) {
	format := utils.Format("is {}, {}, {}, {}, {}, {}",
		"sfd", 4, 123.3, Name{23}, []int{12, 3, 4, 5},
		map[string]string{"a": "as"})

	println(format)
}
