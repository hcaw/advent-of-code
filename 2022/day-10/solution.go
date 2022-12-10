package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcSigs(cycle int, cycleToCheck int, x int, sigs *[]int) int {
	addToCycle := 0
	if cycle == cycleToCheck {
		addToCycle = 40
		*sigs = append(*sigs, cycle*x)
	}
	return addToCycle
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	cycleToCheck := 20
	var sigs []int
	cycle, x := 0, 1
	for _, line := range lines {
		instr := strings.Fields(line)[0]
		switch instr {
		case "noop":
			cycle += 1
			cycleToCheck += calcSigs(cycle, cycleToCheck, x, &sigs)
		case "addx":
			n, _ := strconv.Atoi(strings.Fields(line)[1])
			cycle += 1
			cycleToCheck += calcSigs(cycle, cycleToCheck, x, &sigs)
			cycle += 1
			cycleToCheck += calcSigs(cycle, cycleToCheck, x, &sigs)
			x += n
		}
	}
	sum := 0
	for _, sig := range sigs {
		sum += sig
	}
	fmt.Println("Solution to problem 1 is", sum)
}
