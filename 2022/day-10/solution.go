package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	cycle, i, check, x, processing := 1, 0, 20, 1, false
	sigs := make([]int, 0)
	for i < len(lines) {
		// START OF CYCLE
		if cycle == check {
			check += 40
			sigs = append(sigs, cycle*x)
		}
		line := lines[i]
		if processing {
			processing = false
			// END OF CYCLE
			i += 1
			amountToAdd, _ := strconv.Atoi(strings.Fields(line)[1])
			x += amountToAdd
			} else {
				instr := strings.Fields(line)[0]
				if instr == "addx" {
					processing = true
				} else {
					i += 1
				}
		}
		cycle += 1
	}
	sum := 0
	for _, sig := range sigs {
		sum += sig
	}
	fmt.Println("Solution to problem 1 is", sum)
}
