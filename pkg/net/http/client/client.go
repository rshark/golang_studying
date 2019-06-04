package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go request(i)
	}
	time.Sleep(30 * time.Second)
}

func request(n int) {
	rURL := fmt.Sprintf("http://localhost:3080/httpgoroutine?requestID=%d", n)
	r, err := http.NewRequest("GET", rURL, nil)
	if err != nil {
		fmt.Println("request: ", n, "new request error")
		return
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("request: ", n, "client do error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request: ", n, "ioutil read all error")
		return
	}

	fmt.Println("request: ", n, "response: ", string(body))

}
