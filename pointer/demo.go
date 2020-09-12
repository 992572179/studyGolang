package main

import "fmt"

func main() {
	printPointerMessage()
}

func printPointerMessage() {
	count := 20
	countPointer := &count
	pointerValue := *countPointer

	fmt.Printf("countPointer: %x\n", countPointer)
	fmt.Printf("countPointerValue: %d", pointerValue)
}
