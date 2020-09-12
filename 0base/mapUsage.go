package main

import "fmt"

var (
	m = map[string]string{
		"city":    "wuhan",
		"weather": "rain",
		"pm2.5":   "20",
	}
)

func makeMap2() {
	m := make(map[string]string)
	m["city"] = "wuhan"
	m["weather"] = "rain"
	m["pm2.5"] = "50"
	fmt.Println(m)
}

func rangeMap() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func getKeyIfExists(key string) {
	if k, ok := m[key]; ok {
		fmt.Println(k)
	} else {
		fmt.Println("ERROR-->  key doesn't exists!")
	}
}

func deleteKey(key string) {
	k := m[key]
	fmt.Println("del key:", k)
	delete(m, key)
}

func main() {

	//fmt.Println()
	//makeMap2()

	rangeMap()
	fmt.Println()
	getKeyIfExists("city")
	deleteKey("city")
	fmt.Println("after del")
	getKeyIfExists("city")

}
