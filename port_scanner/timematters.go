package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	tt := time.Now()
	
	for i := 1; i <= 5000; i++ {
		go func(j int) {
			address := fmt.Sprintf("aubit.edu.in:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			defer conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	time.Sleep(20*time.Second)
	fmt.Println(time.Since(tt))
}
