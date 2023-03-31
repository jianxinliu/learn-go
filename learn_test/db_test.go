package learn

import (
	"learn-go/base"
	"testing"
)

func TestDbConn(t *testing.T) {
	for i := 0; i < 10; i++ {
		base.GetConn()
	}
}
