package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func tryChan() {
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)

}

//创建缓冲区大小为i的chan
func createBuffedChan(i int) chan int {
	return make(chan int, i)
}

func main() {
	//tryChan()

	c := createBuffedChan(2)
	c <- 1
	c <- 2
	close(c)
	//c<-3
}
