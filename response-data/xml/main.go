package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sites.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	v := ObjectSites{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

	// Employees
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("Dosya açılırken bir hata oluştu", err)
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Company
	xml.Unmarshal(xmlData, &c)
	fmt.Println(c.Persons)

}

type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site   `xml:"site"`
	Description string   `xml:", innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
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

func (p Person) String() string {
	return fmt.Sprintf("\t ID: %d - First Name: %s - Last Name: %s - User Name: %s\n", p.ID, p.FirstName, p.Lastname, p.UserName)
}
