package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	fmt.Println(array2)

	var array3 [5]int
	array3[0] = 999
	fmt.Println(array3)

	array4 := [5]int{1, 2, 3, 4, 6}
	fmt.Println(array4)

	//tamanho do array de acordo com a quantidade de parametros
	array5 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array5)

	//slice = fatia de array
	slice := []int{10, 11, 12}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array5))

	//adicionar um item ao slice
	slice = append(slice, 5555)
	fmt.Println(slice)

	slice2 := array5[1:3]
	fmt.Println(slice2)

	// arrays internos
	fmt.Println("----------------------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 45)
	slice3 = append(slice3, 46)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
