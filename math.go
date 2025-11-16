package main

import "fmt"

func main() {
	s1 := 15
	s2 := 4

	sum := s1 + s2
	subtraction := s1 - s2
	multiplication := s1 * s2
	division := float64(s1) / float64(s2)
	modulo := s1 % s2

	fmt.Println(sum)
	fmt.Println(subtraction)
	fmt.Println(multiplication)
	fmt.Println(division)
	fmt.Println(modulo)

	if s1 > 10 {
		fmt.Printf("%d is greater than 10 \n", s1)
	} else if s1 >= 5 && s1 <= 10 {
		fmt.Printf("%d is less than 10 but greater than 5\n", s1)
	} else {
		fmt.Printf("%d is less than 5\n", s1)
	}

	number := 11

	switch number {
	case 10:
		fmt.Printf("%d is less than 10 \n", number)
	case 11:
		fmt.Printf("%d is greater than 10 \n", number)
	default:
		fmt.Printf("%d is not less or greater than 10 \n", number)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			// break
			continue
		}
		fmt.Println(i)
	}
	fmt.Printf("End!")
}
