package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, world!")

	// variables
	// var name string = "David"

	// fmt.Printf("This is my name: %s\n", name)

	// age := 27

	// fmt.Printf("my age is: %d\n", age)

	// var city string

	// city = "Tempe"
	// fmt.Printf("I live in $s\n", city)

	// var name, city string = "David", "Tempe"

	// fmt.Printf("Hello, my name is %s and I live in %s\n", name, city)

	var (
		name       string = "David"
		age        int    = 21
		city       string = "Tempe"
		isEmployed bool   = true
	)

	fmt.Printf("Hello! My name is %s and i'm from %s. I'm %d year's old.\n", name, city, age)
	fmt.Printf("It's %t that i got an internship this fall.\n", isEmployed)

	// Zero values

	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("Int: %d\nFloat: %f\nString: %s\nBool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// const - can't change this variable
	const pi = 3.14

	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
	)

	fmt.Printf("Monday: %d - Tuesday: %d - Wednesday: %d\n", monday, tuesday, wednesday)

	const typedAge int = 25

	const unTypedAge = 25

	fmt.Println(typedAge == unTypedAge)

	// emuns

	const (
		Jan int = iota + 1 // 1
		Feb                // 2
		Mar                // 3
		Apr                // 4
	)

	fmt.Printf("jan: %d - feb: %d - mar: %d - apr: %d\n", Jan, Feb, Mar, Apr)

	const (
		num1 int = 1
		num2 int = 2
	)

	fmt.Printf("Total: %d\n", add(num1, num2))
	// add(3, 4)

	sum, _ := calculateSumAndProduct(1, 1)
	// sum, product := calculateSumAndProduct(1, 1)

	fmt.Printf("The sum is: %d\nThe product is: %d\n", sum, 0)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
