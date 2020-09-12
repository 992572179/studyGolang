package main

import (
	"fmt"
	"github.com/study-golang/reflectw"
	"reflect"
)

func main() {
	stu := reflectw.Student{Name: "zs", Age: 18, Class: "class F"}
	//fmt.Println(stu)

	fmt.Println(reflect.TypeOf(stu))
	fmt.Println(reflect.ValueOf(stu))

	reflectw.GetAllFields(stu)

	//b:=10
	//reflectw.GetAllFields(b)

	reflectw.SetValueByReflection(&stu, 22)
	fmt.Println(stu)

}
