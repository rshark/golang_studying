package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 测试ReadAll()
	//callReadAll()

	// 测试ReadDir()
	//callReadDir()

	// 测试ReadFile()
	//callReadFile()

	// 测试TempDir()
	callTempDir()
}

func callReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}

func callReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func callReadFile() {
	content, err := ioutil.ReadFile("./ioutil.go")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File Content: %s", content)
}

func callTempDir() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}
