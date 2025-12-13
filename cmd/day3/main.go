package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	b, err := os.ReadFile("input")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Total Outputs Joltage: ", totalOutputJoltage(string(b)))

}

func totalOutputJoltage(input string) int {

	banks := strings.Split(input, "\n")

	totalJoltage := 0
	for _, bank := range banks {
		maxJoltage := findMaxJoltage(bank)
		totalJoltage += maxJoltage
	}
	return totalJoltage
}

func findMaxJoltage(bank string) int {
	firstDigit := 0
	secondDigit := 0
	for i, c := range bank {
		digit, _ := strconv.Atoi(string(c))
		// check for first digit on all but last position
		if i < len(bank)-1 {
			if digit > firstDigit {
				firstDigit = digit				
				secondDigit = 0
				continue
			}
		}

		if digit > secondDigit {
			secondDigit = digit
		}
	}

	return (firstDigit * 10) + secondDigit
}
