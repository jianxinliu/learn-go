package base

import (
	"testing"
)

func TestDbConn(t *testing.T) {
	for i := 0; i < 10; i++ {
		GetConn()
	}
}
