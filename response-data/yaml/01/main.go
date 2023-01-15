package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// import (
// 	"github.com/go-yaml/yaml"
// )

func main() {
	// Object to yaml
	p := Person{1, "John", "Doe", "jodo", 31}
	fmt.Println(p)
	y, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(y))

}

type Person struct {
	ID        int    `yaml:"id"`
	FirstName string `yaml:"first_name"`
	Lastname  string `yaml:"last_name"`
	UserName  string `yaml:"user_name"`
	Age       int    `yaml:"age"`
}
