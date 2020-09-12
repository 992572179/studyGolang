package main

import (
	"github.com/study-golang/interfaces"
	 "fmt"
)

func main()  {

	stu:=new(interfaces.Student)
	stu.Name = "su"
	stu.Age = 18
	fmt.Println(stu)

	stu.String()

}