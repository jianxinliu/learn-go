package learn

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t2 := time.Now()
	// layout: 2006 01 02 15 05 07
	// Only these values are recognized
	format := t2.Format("2006=01-02 15:04_05")
	println(format)
}
