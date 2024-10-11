package strings

import "strings"

func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsAbsBlank(s string) bool {
	return len(s) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}
