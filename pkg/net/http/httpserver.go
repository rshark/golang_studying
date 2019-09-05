package main

import (
	"net/http"
	"sync"
)

var count int64
var lock sync.Mutex

func main() {

	req := http.NewRequest("POST", "https://dev-api.idaddy.cn")

}
