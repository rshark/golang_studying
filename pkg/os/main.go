package main

import (
	"fmt"
	"os"
)

func main() {
	goPath := os.Getenv("GOPATH")
	fmt.Println("goPath: ", goPath)
	path := os.Getenv("PATH")
	fmt.Println("path: ", path)
}
