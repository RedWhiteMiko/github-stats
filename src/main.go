package main

import (
	"os"
	"scanner/file"
	"scanner/shell"
)

func main() {
	if len(os.Args) > 1 {
		file.Parse(os.Args[1])
		return
	}
	shell.Parse()
	return
}
