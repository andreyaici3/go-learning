package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for k := 1; k < (i*2)-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 5; i >= 2; i-- {
		for j := i; j >= (i*2)-5; j-- {
			fmt.Print(" ")
		}

		for k := (i * 2) - 2; k > 1; k-- {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
