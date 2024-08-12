

package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	tt := time.Now()
	var address string = "aubit.edu.in"
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go PortScanner(address, i)
	}
	wg.Wait()
	fmt.Println(time.Since(tt))
}

// portscanner
func PortScanner(address string, port int) {
	//fmt.Println(Ping(address))
	defer wg.Done()
	fulladdress := fmt.Sprintf("%s:%d", address, port)

	conn, err := net.DialTimeout("tcp", fulladdress, 30*time.Second)
	if err != nil {
		//fmt.Println("closed : ", port, err)
		//wg.Done()
		return
	}
	conn.Close()
	fmt.Println("open : ", port)

}

func Ping(address string) time.Duration{
	begin := time.Now()
	conn , err := net.Dial("tcp","mugund10.openwaves.in:443")
	if err != nil {
		return	time.Duration(0)
	}
	defer conn.Close()
	end := time.Since(begin)
	return end
}
