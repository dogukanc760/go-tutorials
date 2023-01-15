package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("dosyam.txt")
	if err != nil {
		errors.New("Bu bir hata=>" + err.Error())
	}
	fmt.Println(file)

}
