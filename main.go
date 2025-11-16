package main

import "fmt"

func main() {
	//fullName := "Full Name"
	//age := 99
	//fmt.Printf("Hello %s and your age is %d\n", fullName, age)
	//message := fmt.Sprintln("My name is: ", fullName)
	//fmt.Println(message)

	name := "Name"
	age := 99
	height := 9.9
	isGraduated := true
	percent := 100
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of height: %T\n", height)
	fmt.Printf("Type of isGraduated: %T\n", isGraduated)

	fmt.Printf("Type of height (2 decimal places): %.2f\n", height)
	fmt.Printf("Percent: %d%% \n", percent)
	fmt.Printf("I'm: %v", name)

}
