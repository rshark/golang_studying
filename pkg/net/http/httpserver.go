package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var count int64
var lock sync.Mutex

func main() {
	defer fmt.Println("main end")

	handler := func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query()
		requestID := v.Get("requestID")
		fmt.Println("requesting ID: ", requestID, " and Time: ", time.Now())

		lock.Lock()
		count++
		localCount := count
		lock.Unlock()

		time.Sleep(time.Duration(localCount) * time.Second)

		resp := []byte(fmt.Sprintf("requestID %s resp count %d", requestID, localCount))

		w.Write(resp)
	}

	http.HandleFunc("/httpgoroutine", handler)

	http.ListenAndServe("", nil)
	fmt.Println("listening localhost:80")
}
