package main

import "fmt"

func main() {
	// just to explicitly make a slice and not an array
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range nums {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
