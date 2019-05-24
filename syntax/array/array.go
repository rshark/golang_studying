package main

import (
	"fmt"
)

func main() {
	adIDs := make([]int, 0, 1)
	adIDs = append(adIDs, 1)
	adIDs = append(adIDs, 2)
	fmt.Println(adIDs)
}
