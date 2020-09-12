package main

import "fmt"

//slice有指向第一个元素的ptr、数组长度len、整个数组cap，但不能通过数组索引访问cap中除len之外的元素
func sliceUpgrade() {
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6}
	s1 := arr2[2:6] //{2,3,4,5}

	//panic: runtime error: index out of range 不能通过索引访问
	//value := s1[4]
	//fmt.Println(value)
	s2 := s1[3:5] //{5,6}
	fmt.Println("s1-->", s1)
	fmt.Printf("s1%v,len(s1)=%d,cap(s1)=%d", s1, len(s1), cap(s1))
	fmt.Println()
	fmt.Printf("s2%v,len(s2)=%d,cap(s2)=%d", s2, len(s2), cap(s2))
}

func main() {

	arr := [...]int{1, 3, 5, 7, 9}
	s := arr[2:4]
	fmt.Println(s)
	fmt.Printf("%T", s)
	fmt.Println()

	sliceUpgrade()

}
