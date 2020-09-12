package interfaces

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

//接口的组合
type ReaderAndWriter interface {
	Reader
	Writer
}

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	return fmt.Sprintf("%d:%s",stu.Age,stu.Name)
}



