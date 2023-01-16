package main

import (
	"fmt"
	"strconv"
)

const aciklama = "Go"
const pi = 3.14

func main() {
	// var message string
	// message = "adsfadsf"
	// Değişken atamaları
	var a, b, c int
	a = 1
	b = 2
	c = 3

	fmt.Println(a, b, c)

	var p, k = 42, "asdf"

	fmt.Println(p, k)

	// tür ve tip belirtmeksizin atama yaparken := kullanılır
	u := 55

	fmt.Println("u=>", u)

	myComplex := complex(3, 4)
	println(myComplex)

	println(aciklama, pi)

	// type casting
	str := "32"
	i := 42
	f := float64(42)
	ye := uint(i)

	// str to int
	number, _ := strconv.Atoi(str)

	println(i, f, ye, number)

	result := number + 2
	println(result)

	// int to str
	fmt.Println("Sonuç" + strconv.Itoa(result))

	message := "message"
	// call func
	sayHello(&message)
	println(message)
	v := Vertex{1, 2}
	ke := &v
	ke.X = 1e9
	fmt.Println(v)

}

type Vertex struct {
	X int
	Y int
}

func sayHello(message *string) {
	println(*message)
	*message = "heil hitler"
}
