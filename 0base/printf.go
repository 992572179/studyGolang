package main

import "fmt"

func main() {
	a, b, c, d := 10, 3.1415926, true, 'a'
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%t\n", c)
	fmt.Printf("%c\n", d)
	fmt.Printf("%p\n", &d)
	fmt.Printf("%p\n", &a)
}
