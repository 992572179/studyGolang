package main

import (
	"fmt"
)

func main() {

	panicOccur()

}

func panicOccur() {
	defer recoverPanic()
	//panic("file not found,panic")
	//panic(errors.New("error"))
	panic(1)
}

func recoverPanic() {
	msg := recover()
	switch msg.(type) {
	case string:
		fmt.Println("cover string panic:", msg)
	case error:
		fmt.Println("cover error panic:", msg)
	default:
		fmt.Println("unknow panic:",msg)

	}
}
