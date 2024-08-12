package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup
func main() {

	for i := 1 ; i <= 50000; i ++ {
		defer wg.Done()
		wg.Add(1)
	go func () {
		resp, err := http.Get("https://blog.openwaves.in/blog")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	}()
}
	wg.Wait()
}
