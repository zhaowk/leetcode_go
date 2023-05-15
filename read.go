package _0xx

import (
	"io"
	"os"
)

func ReadString(name string) string {
	f, e := os.Open(name)
	if e != nil {
		return ""
	}
	b, e := io.ReadAll(f)
	return string(b)
}
