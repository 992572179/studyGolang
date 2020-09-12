package main

import (
	"encoding/json"
	"fmt"
	"time"
	"reflect"
)

func main() {
	Serialize()
	Unmarshal()
}

type Student struct {
	StuId   int       `json:"id"`
	Name    string    `json:"name"`
	Class   string    `json:"class"`
	RegTime time.Time `json:"reg_time"`
}

func Serialize() {
	stu := new(Student)
	stu.StuId = 1001
	stu.Name = "zs"
	stu.Class = "Class F"
	stu.RegTime = time.Now()
	b, err := json.Marshal(&stu) //注意传指针，尤其是使用map的时候
	if err != nil {
		fmt.Println("Marshal error:", err.Error())
		return
	}
	fmt.Println("After Marshal stu:", string(b))
	fmt.Println(reflect.TypeOf(b))
}

func Unmarshal() {
	jsonString := `{"id":1001,"name":"zs","class":"Class F","reg_time":"2020-04-30T22:32:21.7691063+08:00"}`
	stu := new(Student)
	err := json.Unmarshal([]byte(jsonString), &stu)
	if err != nil {
		fmt.Println("Unmarshal error:", err.Error())
		return
	}
	fmt.Println("After Unmarshal from json:", stu)
	fmt.Println(reflect.TypeOf(stu))
}
