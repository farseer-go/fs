package flog

import (
	"fmt"
	"strconv"
)

// brush is a color join function
type brush func(any) string

var Colors = []brush{
	newBrush("1;32"), // Trace              green
	newBrush("1;36"), // Debug              Background blue
	newBrush("1;34"), // Informational      blue
	newBrush("1;33"), // Warning            yellow
	newBrush("1;31"), // Error              red
	newBrush("1;35"), // Critical           magenta
	newBrush("1;37"), // NoneLevel          white
	newBrush("1;44"), // Alert              cyan
	newBrush("4"),    // datetime           Underline
}

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text any) string {
		return fmt.Sprintf("%s%sm%v%s", pre, color, text, reset)
	}
}

// RedInt64 转为红色字体
func RedInt64(v int64) string {
	return Colors[4](strconv.FormatInt(v, 10))
}

// RedInt64 转为红色字体
func RedInt(v int) string {
	return Colors[4](strconv.Itoa(v))
}
