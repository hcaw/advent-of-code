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

func splitIntoGroups(lines []string) [][3]string {
	var newLines [][3]string
	for i := 0; i < len(lines); i += 3 {
		group := [3]string{lines[i], lines[i+1], lines[i+2]}
		newLines = append(newLines, group)
	}
	return newLines
}

func getCommonCharOf2(a string, b string) (rune, error) {
	for _, charA := range a {
		if strings.ContainsRune(b, charA) {
			return charA, nil
		}
	}
	return -1, errors.New("unable to find matching rune")
}

func getCommonCharOf3(group [3]string) (rune, error) {
	for _, char := range group[0] {
		if strings.ContainsRune(group[1], char) && strings.ContainsRune(group[2], char) {
			return char, nil
		}
	}
	return -1, errors.New("unable to find matching run")
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
	return -1, errors.New("rune was outside the expected scope")
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	splitLines := splitAllLines(lines)
	totalPrioritiesA := 0
	for _, splitLine := range splitLines {
		commonChar, _ := getCommonCharOf2(splitLine[0], splitLine[1])
		priority, _ := getPriority(commonChar)
		totalPrioritiesA += priority
	}
	fmt.Println("Solution to problem 1:", totalPrioritiesA)

	groups := splitIntoGroups(lines)
	totalPrioritiesB := 0
	for _, group := range groups {
		commonChar, _ := getCommonCharOf3(group)
		priority, _ := getPriority(commonChar)
		totalPrioritiesB += priority
	}
	fmt.Println("Solution to problem 2:", totalPrioritiesB)
}
