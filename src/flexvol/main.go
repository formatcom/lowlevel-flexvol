package main

import (
	"flexvol/cmd"
)

var (
	// VERSION is set during build
	VERSION string
)

func main() {
	cmd.Execute(VERSION)
}
