package main

import (
	"nthr/files"
)

func main() {
	fs := files.GetFileStructure("tempdir")
	fs.Display()

	//server.Start()
}
