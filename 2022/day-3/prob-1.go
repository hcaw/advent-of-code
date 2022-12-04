package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func splitAllLines(lines []string) [][2]string {
	var newLines [][2]string
	for _, line := range lines {
		length := len(line)
		half := length / 2
		a := line[:half]
		b := line[half:]
		newLine := [2]string{a, b}
		newLines = append(newLines, newLine)
	}
	return newLines
}

func getCommonChar(a string, b string) (rune, error) {
	for _, charA := range a {
		if strings.ContainsRune(b, charA) {
			return charA, nil
		}
	}
	return -1, errors.New("Unable to find matching rune")
}

func getPriority(item rune) (int, error) {
	intItem := int(item)
	// If uppercase
	if intItem >= 65 && intItem <= 90 {
		return intItem - 38, nil
	}
	// If lowercase
	if intItem >= 97 && intItem <= 122 {
		return intItem - 96, nil
	}
	return -1, errors.New("Rune was outside the expected scope")
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	splitLines := splitAllLines(lines)
	totalPriorities := 0
	for _, splitLine := range splitLines {
		commonChar, _ := getCommonChar(splitLine[0], splitLine[1])
		priority, _ := getPriority(commonChar)
		totalPriorities += priority
	}
	fmt.Println(totalPriorities)
}
