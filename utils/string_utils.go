package utils

import (
	"fmt"
	"strings"
)

func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(strings.Trim(s, " ")) == 0 {
		return true
	}
	return false
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

type Printable interface {
	ToString() string
}

// Format supports placeholder {} for any types
func Format(formatter string, param ...any) string {
	for i := 0; i < len(param); i++ {
		item := fmt.Sprint(param[i])
		if printable, ok := param[i].(Printable); ok {
			item = printable.ToString()
		}
		formatter = strings.Replace(formatter, "{}", item, 1)
	}
	return formatter
}
