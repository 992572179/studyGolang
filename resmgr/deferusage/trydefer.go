package deferusage

import (
	"fmt"
	"os"
	"bufio"
)

func TryDefer() {
	fmt.Println(1)

	defer fmt.Println(2)
	defer fmt.Println(3)
	return
}

func FibGen() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func WriteFile(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()
	f := FibGen()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(w, f())
	}

}
