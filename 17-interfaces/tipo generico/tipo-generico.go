package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(15)
	generica(true)

	fmt.Println(1, 2, "string", true, 123.5)

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     69,
	}
	fmt.Println(mapa)
}
