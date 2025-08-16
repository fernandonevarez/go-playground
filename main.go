package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{"Ian", 30}

	// fmt.Println(person.name)

	fmt.Printf("This is our person: %+v\n", person)

	employee := struct {
		name string
		id   int
		age  int
	}{
		name: "David",
		id:   123,
		age:  21,
	}

	fmt.Println(employee)

	type Address struct {
		street string
		city   string
	}

	type Contact struct {
		name    string
		address Address
		phone   string
	}

	contact := Contact{
		name: "David",
		address: Address{
			street: "123 main street",
			city:   "anytown",
		},
		phone: "123-123-1234",
	}

	fmt.Println(contact)

	// pointers and struct methods

	// fmt.Println("name before:", person.name)

	// modifyPersonName(&person)

	// fmt.Println("name after", person.name)

	x := 20

	ptrX := &x

	// Pointer arithmetic is not allowed in Go, so we remove ptrX + 1 usage.
	// If you want to increment the value pointed to by ptrX:
	*ptrX += 1

	fmt.Println(*ptrX)

	fmt.Println(person)

	person.modifyPersonName()
	person.modifyPersonAge()

	fmt.Println(person)

}

func (person *Person) modifyPersonAge() {
	person.age = 21
	// fmt.Println("inside scoop:", person.age)
}

// makes it so that modifyPersonName() is now a method in person
func (person *Person) modifyPersonName() {
	person.name = "David"
	// fmt.Println("inside scoop: new name", person.name)
}
z
func (p *Person) nameChange(name string) {
	p.name = name
}
