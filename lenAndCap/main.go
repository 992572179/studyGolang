package main

import "fmt"

func main() {
	printLenAndCap()
}

func printLenAndCap() {
	s := make([]int, 2, 5)
	s[0] = 5
	s[1] = 12
	fmt.Println("cap : ", len(s))
	fmt.Println("len : ", cap(s))

}
