package main

type Human struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// x := Human{firstName: "John", lastName: "Doe", age: 31}

	// println(x.age)

	x := new(Human)
	x.firstName = "John"
	x.lastName = "Doe"

	println(x.age)

	// consturctor method kullanımı
	y := NewHuman()
	y.age = 28
	println(y.age)

	z := NewHumanWithParam("name", "firstname", 23)
	println(z.age, z.firstName, z.lastName)

}

// varsayılan ve boş constructor
func NewHuman() *Human {
	x := new(Human)
	return x
}

// parametre alan constructor
func NewHumanWithParam(firstName, lastName string, age int) *Human {
	x := new(Human)
	x.age = age
	x.firstName = firstName
	x.lastName = lastName
	return x
}
