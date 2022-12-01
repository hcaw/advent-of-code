package main

import (
	"fmt"
	"os"
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

	bigBoi := 0
	for _, cals := range calsPerElf {
		if cals > bigBoi {
			bigBoi = cals
		}
	}
	fmt.Println(bigBoi)
}
