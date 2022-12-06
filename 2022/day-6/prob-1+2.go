package main

import (
	"fmt"
	"os"
)

func getUniqueSequenceEnd(input string, uniqLen int) int {
	for i := uniqLen - 1; i < len(input); i += 1 {
		set := make(map[byte]struct{})

		for j := 0; j < uniqLen; j++ {
			set[input[i-j]] = struct{}{}
		}
		if len(set) == uniqLen {
			return i + 1
		}
	}
	return -1
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	datastream := string(input)

	prob1Solution := getUniqueSequenceEnd(datastream, 4)
	prob2Solution := getUniqueSequenceEnd(datastream, 14)
	fmt.Println(prob1Solution)
	fmt.Println(prob2Solution)
}
