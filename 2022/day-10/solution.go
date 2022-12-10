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

	check, x, processing := 20, 1, false
	sigs := make([]int, 0)
	for i, cycle := 0, 1; i < len(lines); cycle++ {
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
			if instr := strings.Fields(line)[0]; instr == "addx" {
				processing = true
			} else {
				i += 1
			}
		}
	}
	sum := 0
	for _, sig := range sigs {
		sum += sig
	}
	fmt.Println("Solution to problem 1 is", sum)
}
