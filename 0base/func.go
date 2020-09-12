package main

import (
	"fmt"
	"math"
)

//函数作为入参
func calc(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {

	fmt.Println(calc(pow, 3, 4))

	fmt.Println(calc(max, 2020, 100))

}
