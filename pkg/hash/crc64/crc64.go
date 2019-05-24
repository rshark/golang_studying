package main

import (
	"fmt"
	// "math"
	"strings"
)

const (
	baseStr string = "9AbCd1EfGhIJkLm2NoPq3RsTuVw8XyZ5aBcDe6FgHijK4lMnOp7QrS0tUvWxYz"
	base    int    = 62
)

var count int64 = 1001

func main() {

	URL := "https://www.google.com"
	sURL := shortenURL(URL)
	sID := unshortenURL(sURL)
	fmt.Printf("URL %s sURL %s sID %d \n", URL, sURL, sID)

	for i := 0; i < 100000; i++ {
		URL = "https://www.google.com"
		sURL = shortenURL(URL)
		sID = unshortenURL(sURL)
		fmt.Printf("URL %s sURL %s sID %d \n", URL, sURL, sID)
	}

}

func checksum(str string) int {
	bytes := []byte(str)
	var sum int
	for _, b := range bytes {
		sum += int(rune(b) - '-')
	}
	return sum % base
}

func digestStr(str string) string {
	index := checksum(str)
	fmt.Printf("digestStr index %d\n", index)
	return string(baseStr[index])
}

func signatureStr(str string) string {
	digest := digestStr(str)
	return str + digest
}

func validateSig(str string) bool {
	length := len(str)
	if length > 1 {
		preStr := str[:length-1]
		digest := str[length-1:]
		preDigest := digestStr(preStr)
		fmt.Printf("preStr %s digest %s preDigest %s\n", preStr, digest, preDigest)
		if preDigest == digest {
			return true
		}
	}
	return false
}

func shortenURL(url string) string {
	count++
	shortID := count
	shortURL := ""
	for shortID/int64(base) != 0 {
		shortURL += string(baseStr[shortID%int64(base)])
		shortID = shortID / int64(base)
	}
	if shortID != 0 {
		shortURL += string(baseStr[shortID])
	}
	fmt.Printf("shortenURL sURL %s\n", shortURL)
	shortURL = signatureStr(shortURL)

	return shortURL
}

func unshortenURL(str string) int64 {
	var shortID int64
	if validateSig(str) {
		preStr := str[:len(str)-1]
		length := len(preStr)
		indexOfByte := 0
		for i := length - 1; i >= 0; i-- {
			indexOfByte = strings.IndexByte(baseStr, preStr[i])
			shortID = shortID*int64(base) + int64(indexOfByte)
		}
	} else {
		panic("无效短链接")
	}

	return shortID
}
