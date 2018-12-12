package main

import (
	"flag"
	"go_playground/coursera/tree/tree"
)

func main() {
	path := flag.String("path", ".", "Path to directory")
	needFiles := flag.Bool("files", false, "Set to view files")
	flag.Parse()

	tree.Tree(*path, needFiles)
}
