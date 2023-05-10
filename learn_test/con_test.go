package learn

import (
	"sync"
	"testing"
)

var lock sync.Mutex

func TestCon(t *testing.T) {

}

type Op struct {
	Num int
}

func (op *Op) Incr() {
	lock.Lock()
	defer lock.Unlock()
	op.Num++
}
