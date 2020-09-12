package main

import (
	"math/cmplx"
	"math"
	"fmt"
)

func euler() {
	res := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(res)
	fmt.Printf("%T", res)
}


//勾股定理
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

func main() {

	euler()
	fmt.Println()
	triangle()

}
