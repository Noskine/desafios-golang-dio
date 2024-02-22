package main

import "fmt"

func main() {
	fmt.Println("Desafio 1")
	for i := 1; i <= 100; i++ {
		x := i % 3

		if x == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\n\nDesafio 2")
	for i := 1; i <= 100; i++ {
		x := i % 3
		z := i % 5

		if x == 0 && z == 0 {
			fmt.Println("pin/pan")
		} else {
			if x == 0 {
				fmt.Println("pin")
			} else if z == 0 {
				fmt.Println("pan")
			} else {
				println(i)
			}
		}
	}
}
