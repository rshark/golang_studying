package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// intent() example
	jsonIntentExample()

	// Unmarshal
	jsonUnmarshal()
}

func jsonIntentExample() {
	type Road struct {
		Name   string
		Number int
	}

	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)

	//	json.Compact()
}

type studentLowC struct {
	name string
	age  string
}

type studentUpperC struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type studentNoTag struct {
	Name string
	Age  string
}

func jsonUnmarshal() {
	data := []byte(`{"name":"张三","age":"9岁","class":"三年级"}`)

	var stuLowC studentLowC
	var stuUpperC studentUpperC
	var stuNoTag studentNoTag

	json.Unmarshal(data, &stuLowC)
	json.Unmarshal(data, &stuUpperC)
	json.Unmarshal(data, &stuNoTag)

	fmt.Println(stuLowC)
	fmt.Println(stuUpperC)
	fmt.Println(stuNoTag)
}
