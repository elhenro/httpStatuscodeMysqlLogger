package main

import (
	"strings"
)

func Join(strs ...string) string {
	b := strings.Builder{}
	for _, str := range strs {
		b.WriteString(str)
	}
	r := b.String()
	return r
}
