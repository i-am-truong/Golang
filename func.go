package main

import "fmt"

func sum(n1, n2 int) (int, int, int, float64) {
	if n2 == 0 {
		n2 = 5
	}
	sum := n1 + n2
	subtraction := n1 - n2
	multiplication := n1 * n2
	division := float64(n1) / float64(n2)
	return sum, subtraction, multiplication, division
}

func countdown(number int) {
	fmt.Println(number)
	if number > 0 {
		countdown(number - 1)
	}
}

func main() {
	sum, subtraction, multiplication, division := sum(5, 10)
	fmt.Println(sum)
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(division)
	countdown(5)
}
