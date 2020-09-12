package main

import (
	"fmt"
	"reflect"
)

func main()  {
	printTypeOf()
}


//make函数创建的是值类型，而new则是引用类型
func printTypeOf(){
	map1:=make(map[int]string)
	map2:=new(map[string]int)
	fmt.Println(reflect.TypeOf(map1))
	fmt.Println(reflect.TypeOf(map2))
}
