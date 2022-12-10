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

	pixels := make([]string, 240)
	check, x, processing := 20, 1, false
	sigs := make([]int, 0)
	for i := 0; i < len(pixels); i++ {
		pixels[i] = "."
	}
	for i, cycle := 0, 1; i < len(lines); cycle++ {
		// START OF CYCLE
		if cycle == check {
			check += 40
			sigs = append(sigs, cycle*x)
		}
		pixelHorizontal := (cycle % 40) - 1
		if pixelHorizontal == -1 {
			pixelHorizontal = 39
		}
		if pixelHorizontal == x || pixelHorizontal == x-1 || pixelHorizontal == x+1 {
			pixels[cycle-1] = "#"
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
	for i, pixel := range pixels {
		fmt.Print(pixel)
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	sum := 0
	for _, sig := range sigs {
		sum += sig
	}
	fmt.Println("Solution to problem 1 is", sum)
}
