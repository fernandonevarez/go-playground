package main

import (
	"fmt"
)

func main() {
	age := 12

	if age >= 20 {
		fmt.Println("you are an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Friday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's midweek")
	case "Friday":
		fmt.Println("TGIF")
		fallthrough
	default:
		fmt.Println("It's the weekend")
	}

	// loops

	// for loops

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("this is i: %d\n", i)
	// }

	// for i := range 5 {
	// 	fmt.Printf("this is i: %d\n", i)
	// }

	// while loops

	// i := 0

	// for i < 5 {
	// 	// fmt.Printf("this is i: %d\n", i)
	// 	fmt.Println("this is the counter", i)
	// 	i++
	// }

	// i := 0

	// for {
	// 	if i > 5 {
	// 		break
	// 	}
	// 	fmt.Println("The counter is", i)
	// 	i++
	// }

	// Arrays

	numbers := [5]int{10, 20, 30, 40, 50}
	// numbers[5] = 60     // can't do this

	// fmt.Printf("this is our array: %v\n", numbers)
	// // fmt.Println("this is our array", numbers)

	// fmt.Printf("The length of the array is: %d\n", len(numbers))

	// fmt.Println("The last number in the array is", numbers[len(numbers)-1])

	// matrix := [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }

	// fmt.Printf("%v\n", matrix)
	// fmt.Println(matrix)

	// Slices - dynamtic arrays

	allNumbers := numbers[:] // a copy of numbers, but now we can add more to it(aka we can now append to arrays, well to the copied array - allNumbers)

	firstThree := numbers[0:3]

	fmt.Println("all the numbers", allNumbers)
	fmt.Println("the first three numbers", firstThree)

	fmt.Println("i added some numbers")

	fruits := []string{"apple", "banana", "straberry"}
	fmt.Println("these are my fruits", fruits)

	fruits = append(fruits, "kiwi")

	fmt.Println("these are my fruits", fruits)

	fruits = append(fruits, "kiwi", "mango", "pineapple")

	fmt.Println("these are my fruits", fruits)

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// Maps - hash

	capitalCities := map[string]string{
		"USA":   "Washington D.C",
		"India": "New Delhi",
		"UK":    "London",
	}

	// fmt.Println(capitalCities["USA"])
	// fmt.Println(capitalCities["Germany"])
	capital, exists := capitalCities["Germany"]
	// for value := range capitalCities {
	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Printf("Does not exists\n")
	}

	delete(capitalCities, "UK")
	fmt.Println(capitalCities)
	// }

}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
