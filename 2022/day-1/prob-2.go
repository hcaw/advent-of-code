package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input.txt")
	elves := strings.Split(string(input), "\n\n")

	var calsPerElf []int
	for _, elf := range elves {
		var total int
		for _, calories := range strings.Split(elf, "\n") {
			amount, _ := strconv.Atoi(calories)
			total += amount
		}
		calsPerElf = append(calsPerElf, total)
	}

	sort.Slice(calsPerElf, func(i, j int) bool {
		return calsPerElf[i] > calsPerElf[j]
	})
	bigBois := calsPerElf[:3]
	var totalCals int = 0
	for _, cals := range bigBois {
		totalCals += cals
	}
	fmt.Println(totalCals)
}
