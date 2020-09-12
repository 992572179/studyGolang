package main

import (
	"github.com/study-golang/functional"
	"fmt"
)

func main() {
	f := functional.FibGen()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
