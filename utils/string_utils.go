package utils

import "strings"

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
