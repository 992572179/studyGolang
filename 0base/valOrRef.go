package main

import "fmt"

//默认是值传递，变量改变只存在于该函数中
func swap(a, b int) {
	a, b = b, a
	//fmt.Println(a,b)
}

func swapWithRef(a, b *int) {
	*a, *b = *b, *a
}

func swapByReturn(a, b int) (int, int) {
	return b, a
}

func main() {

	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	c, d := 1, 2
	swapWithRef(&c, &d)
	fmt.Println(c, d)

	e, f := swapByReturn(12, 13)
	fmt.Println(e, f)

}
