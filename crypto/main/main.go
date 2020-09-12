package main

import (
	"github.com/study-golang/crypto"
	"fmt"
)

func main() {

	err := crypto.RsaKeyGen(4096)
	if err != nil {
		fmt.Println("error occurred...")
		fmt.Printf("%v", err)
	}

}
