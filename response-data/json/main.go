package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	jsonStr := `
	  {
			"data":{
			"object":"card",
			"id":"card_123",
			"first_name":"John",
			"last_name":"Doe",
			"balance":"5432.32"
		}
		}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println("--------------------------------")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Convert XML to JSON

	var c Company

	var person jsonPerson
	var persons []jsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		//person.Lastname = value.LastName
		person.UserName = value.UserName

		persons = append(persons, person)
	}

	jsonData2, err := json.Marshal(persons)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData2))

	jsonFile, err := os.Create("./Employees.json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData2)
	jsonFile.Close()
}

type jsonPerson struct {
	ID        int
	FirstName string
	Lastname  string
	UserName  string
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	Lastname  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLName xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}
