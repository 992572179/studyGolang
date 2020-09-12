package main

import (
	"time"
	"fmt"
)

func main() {
	c := make(chan int)
	timeout := make(chan bool)
	go Send(c,timeout)
	go Receive(c,timeout)
	time.Sleep(time.Second * 3)

}

func Send(c chan int, timeout chan bool) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Millisecond * 175)
	}
	timeout <- false
}

func Receive(c chan int, timeout chan bool) {
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case <-timeout:
			fmt.Println("timeout! program exit...")
		}

	}
}
