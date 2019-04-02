package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")
	fmt.Print(&buf)

	logger = log.New(&buf, "INFO", log.Lshortfile)
	infof := func(info string) {
		logger.Output(2, info)
	}

	infof("Hello world!")
	fmt.Print(&buf)
}
