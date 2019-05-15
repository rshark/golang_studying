package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// RandomString ...
func RandomString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)

	for i, b := range bytes {
		fmt.Printf("index %d b %d \n", i, b)
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}
	fmt.Println(string(bytes))
	return string(bytes)
}

// 62 characters table.
var charTable = [...]rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
	'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6',
	'7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
	'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

// ShortenURL, func to short the input URL
func ShortenURL(url string) []string {
	shortURLList := make([]string, 0, 4)
	sumData := md5.Sum([]byte(url))
	// Split MD5 checksum into 4 pieces, 4byte for each piece
	for i := 0; i < 4; i++ {
		part := sumData[i*4 : i*4+4]
		// Convert 4 byte to int32
		partUint := binary.BigEndian.Uint32(part)
		// Only reserve last 30 bits.
		partUint &= 0x3fffffff
		shortURLBuffer := &bytes.Buffer{}
		// Split 30bit into 6 pieces, 5bit for each piece
		for j := 0; j < 6; j++ {
			index := partUint & 0x3d
			shortURLBuffer.WriteRune(charTable[index])
			partUint = partUint >> 5
		}
		shortURLList = append(shortURLList, shortURLBuffer.String())
	}
	return shortURLList
}

func main() {
	/*
		const length = 10000
		for i := 0; i < length; i++ {
			RandomString(7)
		}
	*/

	shortURLList := ShortenURL("https://dev-api.idaddy.cn")
	fmt.Println(shortURLList)
	shortURLList = ShortenURL("http://www.baidu.com")
	fmt.Println(shortURLList)
}
