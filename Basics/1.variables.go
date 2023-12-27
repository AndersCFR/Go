package main

import "fmt"

const secondInHour = 3600

func main() {

	// Simple output
	fmt.Println("Hello Go World!")
	distance := 60.9
	fmt.Printf("The distance in miles is %f", distance*0.621)

	// 1. Explicit declaration
	var age int = 30
	fmt.Println("\nAge : ", age)

	// 2. Implicit declaration
	var name = "Anderson"
	// To unmute unsued variables
	_ = name

	// 3.Short declaration
	profesion := "engineer"
	fmt.Println("Profesion: ", profesion)

	// 4.Multiple declarations
	car, cost, working := "Polo", 6, true
	fmt.Printf("Car: %s - Cost per hour: %d - Is working? %t \n", car, cost, working)

	// Empty multiple declarations
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	//Empty same type declaration
	var a, b, c int
	fmt.Println(a, b, c)

	// Post assignmend and swaping variables
	var i, j int
	_, _ = i, j
	i, j = 5, 8
	j, i = i, j
	fmt.Println(i, j)

	// Expresion in short  declaration
	sum := 5 + 2.3
	fmt.Println(sum)
}
