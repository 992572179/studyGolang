package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	//tryGoRoutine()
	printCpuNum()

	cpuLimit()
}

func tryGoRoutine() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("go routine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}

func printCpuNum(){
	num:=runtime.NumCPU()
	fmt.Printf("%d\n",num)
}

func cpuLimit(){
	processors:=runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	fmt.Println("runtime cpus: ",processors)
}
