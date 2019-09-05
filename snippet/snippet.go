package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(jobID int) {
			fmt.Printf("This is job: %v\n", jobID)
		}(i)
	}
	for true {
	}
}
