package main


import (
	"strings"
)


func extractUsernameAndHost(s string) (string, string) {
	a := strings.SplitN(s, "@", 2)

	switch len(a) {
	case 0, 1:
		return "", s
	default:
		return a[0], a[1]
	}
}
