package main

import (
	"fmt"
)

func calculateAverageScore(scores []int) float64 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

func main() {
	// Data 1
	dolphinScores := []int{96, 108, 89}
	koalaScores := []int{88, 91, 110}

	dolphinAverage := calculateAverageScore(dolphinScores)
	koalaAverage := calculateAverageScore(koalaScores)

	fmt.Println("Data 1:")
	fmt.Printf("Average Score Dolphin: %.2f\n", dolphinAverage)
	fmt.Printf("Average Score Koala: %.2f\n", koalaAverage)

	if dolphinAverage > koalaAverage {
		fmt.Println("Dolphin wins the competition!")
	} else if dolphinAverage < koalaAverage {
		fmt.Println("Koala wins the competition!")
	} else {
		fmt.Println("The competition ends in a draw!")
	}
	fmt.Println()

	// Bonus 1
	dolphinScores = []int{97, 112, 101}
	koalaScores = []int{109, 95, 123}

	dolphinAverage = calculateAverageScore(dolphinScores)
	koalaAverage = calculateAverageScore(koalaScores)

	fmt.Println("Bonus 1:")
	fmt.Printf("Average Score Dolphin: %.2f\n", dolphinAverage)
	fmt.Printf("Average Score Koala: %.2f\n", koalaAverage)

	if dolphinAverage > koalaAverage && dolphinAverage >= 100 {
		fmt.Println("Dolphin wins the competition!")
	} else if koalaAverage > dolphinAverage && koalaAverage >= 100 {
		fmt.Println("Koala wins the competition!")
	} else {
		fmt.Println("No team wins the competition!")
	}
	fmt.Println()

	// Bonus 2
	dolphinScores = []int{97, 112, 101}
	koalaScores = []int{109, 95, 106}

	dolphinAverage = calculateAverageScore(dolphinScores)
	koalaAverage = calculateAverageScore(koalaScores)

	fmt.Println("Bonus 2:")
	fmt.Printf("Average Score Dolphin: %.2f\n", dolphinAverage)
	fmt.Printf("Average Score Koala: %.2f\n", koalaAverage)

	if dolphinAverage > koalaAverage && dolphinAverage >= 100 {
		fmt.Println("Dolphin wins the competition!")
	} else if koalaAverage > dolphinAverage && koalaAverage >= 100 {
		fmt.Println("Koala wins the competition!")
	} else if dolphinAverage == koalaAverage && dolphinAverage >= 100 && koalaAverage >= 100 {
		fmt.Println("The competition ends in a draw!")
	} else {
		fmt.Println("No team wins the competition!")
	}
}