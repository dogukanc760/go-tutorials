package main

import "fmt"

func main() {

	myArray1 := [...]int{1, 2, 3}
	mySlice := myArray1[:]
	fmt.Println(mySlice)
	mySlice[0] = 11
	fmt.Println(mySlice)
	fmt.Println(myArray1)

	fmt.Print("Farklı Bir Şey \n")

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Print(colors)

	fmt.Print("Farklı Bir Slice Yöntemi \n")

	numbers := make([]int, 5, 5)
	numbers[0] = 123
	numbers[1] = 456
	numbers[2] = 789

	numbers = append(numbers, 331)
	fmt.Println(numbers, cap(numbers), len(numbers))

}
