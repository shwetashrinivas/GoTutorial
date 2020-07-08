package main

import "fmt"

func main() {
	eo := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range eo {
		if eo[i]%2 == 0 {
			fmt.Println(eo[i], " is even")
		}
		if eo[i]%2 != 0 {
			fmt.Println(eo[i], " is odd")
		}
	}
}

//Or
func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		}
		if number%2 != 0 {
			fmt.Println(number, " is odd")
		}
	}
}
