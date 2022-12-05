package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Instruction struct {
	quantity int
	from     int
	to       int
}

func reverseSlice[T any](slice []T) []T {
	newSlice := make([]T, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		newSlice[i] = slice[j]
	}
	return newSlice
}

func initStacks(str string) [][]rune {
	lines := strings.Split(str, "\n")
	lastLine := strings.Fields(lines[len(lines)-1])
	numStacks, _ := strconv.Atoi(lastLine[len(lastLine)-1])
	stacks := make([][]rune, numStacks)
	reversedLines := reverseSlice(lines)[1:]
	for _, line := range reversedLines {
		for j, p := 0, 1; j < numStacks; j++ {
			char := line[p]
			if unicode.IsLetter(rune(char)) {
				stack := stacks[j]
				stack = append(stack, rune(char))
				stacks[j] = stack
			}
			p += 4
		}
	}
	return stacks
}

func getInstructions(section string) []Instruction {
	lines := strings.Split(section, "\n")
	var instructions []Instruction
	for _, line := range lines {
		fields := strings.Fields(line)
		quantity, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		instruction := Instruction{quantity, from, to}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func copy2DSlice[T any](slice [][]T) [][]T {
	newSlice := make([][]T, len(slice))
	for i := range slice {
		newSlice[i] = make([]T, len(slice[i]))
		copy(newSlice[i], slice[i])
	}
	return newSlice
}

func processInstructions(stacks [][]rune, instr []Instruction) [][]rune {
	newStacks := copy2DSlice(stacks)
	for _, instruction := range instr {
		from := newStacks[instruction.from-1]
		to := newStacks[instruction.to-1]
		for i := 0; i < instruction.quantity; i++ {
			char := from[len(from)-1]
			from = from[:len(from)-1]
			to = append(to, char)
		}
		newStacks[instruction.from-1] = from
		newStacks[instruction.to-1] = to
	}

	return newStacks
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	sections := strings.Split(string(input), "\n\n")

	stacks := initStacks(sections[0])
	instructions := getInstructions(sections[1])
	newStacks := processInstructions(stacks, instructions)
	for _, stack := range newStacks {
		fmt.Print(string(stack[len(stack) - 1]))
	}
}
