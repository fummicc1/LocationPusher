package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	fmt.Print(r.URL.RawQuery)
}

func httpRun() {
	http.HandleFunc("/location", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go httpRun()
	go pushNotification()
	wg.Wait()
}
