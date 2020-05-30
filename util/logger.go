package util

import "github.com/fatih/color"

var (
	info  = color.New(color.FgCyan).PrintlnFunc()
	err   = color.New(color.FgRed).PrintlnFunc()
	fatal = color.New(color.BgRed).PrintlnFunc()
)

func Info(contents string) {
	info("→", contents)
}

func Error(contents string) {
	err("→", contents)
}

func Fatal(contents string) {
	fatal("→", contents)
}
