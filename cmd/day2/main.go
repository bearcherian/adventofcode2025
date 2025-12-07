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

	fmt.Println("Sum of invalid ranges: ", sumOfInvalidRanges(string(b)))

}

func sumOfInvalidRanges(input string) int64 {
	invalidRanges := gatherInvalidRanges(input)
	var sum int64 = 0
	for _, val := range invalidRanges {
		sum += val
	}
	return sum
}

func gatherInvalidRanges(input string) []int64 {
	invalidRanges := []int64{}
	ranges := strings.Split(input, ",")

	for _, ranges := range ranges {
		bounds := strings.Split(ranges, "-")
		if len(bounds) != 2 {
			continue
		}
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])

		for i := start; i <= end; i++ {
			intStr := strconv.Itoa(i)
			// fmt.Println("Checking integer: ", intStr, " ", intStr[:len(intStr)/2], ":", intStr[len(intStr)/2:])
			if len(intStr) % 2 == 0 && intStr[:len(intStr)/2] == intStr[len(intStr)/2:] {
				invalidRanges = append(invalidRanges, int64(i))
			}
		}
	}

	return invalidRanges
}