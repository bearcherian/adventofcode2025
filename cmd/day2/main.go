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
			strLength := len(intStr)

			for splits := strLength; splits > 0; splits-- {

				intParts := splitStringByN(intStr, splits)
				if len(intParts) == 0 {
					continue
				}

				if len(intParts) > 1 && isAllSameChars(intParts) {
					invalidRanges = append(invalidRanges, int64(i))
					break
				}
			}
			
		}
	}

	return invalidRanges
}

func splitStringByN(s string, n int) []string {
	var parts []string
	if len(s)%n != 0 {
		return parts
	}
	partsCount := len(s) / n
	for i := range partsCount {
		parts = append(parts, s[i*n:(i+1)*n])
	}
	return parts
}

func isAllSameChars(parts []string) bool {
	first := parts[0]
	for _, part := range parts[1:] {
		if part != first {
			return false
		}
	}
	return true
}