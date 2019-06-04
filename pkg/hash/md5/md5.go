package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("https://www.baidu.com")
	sum := md5.Sum(data)

	fmt.Printf("%x\n", sum)
	digest := fmt.Sprintf("%x", sum)
	fmt.Println("digest: ", digest)
}
