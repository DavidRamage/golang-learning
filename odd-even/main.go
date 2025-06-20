package main

import "fmt"

func main() {
	var n []int
	for i := range 11 {
		n = append(n, i)
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
