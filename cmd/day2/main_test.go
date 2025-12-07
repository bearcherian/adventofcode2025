package main

import (
	"slices"
	"testing"
)

func TestExampleInputInvalidRangesSum(t *testing.T) {

	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	expectedInvalidRanges := []int64{
		11, 22, 99, 1010, 1188511885, 222222, 446446, 38593859,
	}

	result := gatherInvalidRanges(input)
	
	if !slices.Equal(expectedInvalidRanges, result) {
		t.Errorf("Expected invalid ranges \n\t%v, but got \n\t%v", expectedInvalidRanges, result)
	}

	expectedSum := int64(1227775554)
	resultSum := sumOfInvalidRanges(input)
	if expectedSum != resultSum {
		t.Errorf("Expected sum of invalid ranges %d, but got %d", expectedSum, resultSum)
	}
}