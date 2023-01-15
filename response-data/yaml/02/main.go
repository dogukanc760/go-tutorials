package main

import (
	"io/ioutil"

	m "./models"
)

func main(){
	fileName:= "./config.yaml" 

	var config m.Config 
	source, err := ioutil.ReadFile(fileName)
	if err !=nil {
		panic(err)
	}
	println(string(source))
}