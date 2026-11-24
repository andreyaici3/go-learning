package main

import "fmt"

func main() {

	scores := []int{80, 85, 90, 75, 95}
	var totalScore int

	for _, score := range scores {
		fmt.Println("Score:", score)
		totalScore += score
	}

	fmt.Println()

	average := totalScore / len(scores)

	fmt.Println("Total:", totalScore)
	fmt.Println("Average:", average)
}
