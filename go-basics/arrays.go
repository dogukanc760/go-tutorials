package main

import "fmt"

func main() {

	// simple array
	simpleArr := [3]int{}
	simpleArr[0] = 0
	simpleArr[1] = 1
	simpleArr[2] = 2
	fmt.Println(simpleArr)

	// another simple array
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Println(colors)
	fmt.Println("Length of simple array=>", len(simpleArr))

	// auto size array
	autoSizeArr := [...]int{1, 2, 3}
	fmt.Println(autoSizeArr)
	fmt.Println("Dizi kapasitesi, kaç eleman alabilir?", cap(simpleArr))


	// Slice

}
