package reflectw

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Class string
}

//range遍历获取所有属性
func GetAllFields(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("Type:", t.Name())

	// 类型判断，不为结构体则报错，且直接返回
	if kind := t.Kind(); kind != reflect.Struct {
		fmt.Println("----->ERROR!")
		return
	}
	v := reflect.ValueOf(a)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v=%v \n", field.Name, field.Type, val)
	}
}


//反射修改属性的值
func SetValueByReflection(a interface{},valArgs int64){
	value:=reflect.ValueOf(a).Elem()
	value.Field(1).SetInt(valArgs)

}
