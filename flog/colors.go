package flog

import (
	"fmt"
	"github.com/farseer-go/fs/parse"
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
	//pre := "\033["
	//reset := "\033[0m"
	return func(text any) string {
		return fmt.Sprintf("\u001B[%sm%v\u001B[0m", color, text)
	}
}

// Red 转为红色字体
func Red(text any) string {
	return "\u001B[1;31m" + parse.Convert(text, "") + "\u001B[0m"
}

// Yellow 转为黄色字体
func Yellow(text any) string {
	return "\u001B[1;33m" + parse.Convert(text, "") + "\u001B[0m"
}

// Green 转为绿色字体
func Green(text any) string {
	return "\u001B[1;32m" + parse.Convert(text, "") + "\u001B[0m"
}
