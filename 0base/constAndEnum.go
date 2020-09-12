package main

import "fmt"

func defConst(){
	const(
		str= "strings"
		command = 5
	)
	fmt.Println(str,command)
}

func defNumbers()  {
	const(
		init = iota
		start
		pause
		resume
		stop
		destroy
	)
	fmt.Println(init,start,destroy)
}

func main() {

	defConst()

	defNumbers()


}
