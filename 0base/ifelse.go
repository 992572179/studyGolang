package main

import (
	"io/ioutil"
	"fmt"
)

func readFile() {
	filename := "123.log"
	if b, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", b)
	}
}

func score(score int) string {
	msg := ""
	switch {
	case score < 0 || score > 100:
		panic("error occurred.")
	case score < 60:
		msg = "C"
	case score < 80:
		msg = "B"
	case score < 90:
		msg = "A"
	case score < 95:
		msg = "S"
	case score <= 100:
		msg = "S+"
	}
	return msg
}

func main() {

	readFile()

	fmt.Println(score(75))
	//fmt.Println(score(-1))
	//fmt.Println(score(101))
	fmt.Println(score(96))

}
