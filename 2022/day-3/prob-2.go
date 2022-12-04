package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func splitIntoGroups(lines []string) [][3]string {
	var newLines [][3]string
	for i := 0; i < len(lines); i += 3 {
		group := [3]string{lines[i], lines[i+1], lines[i+2]}
		newLines = append(newLines, group)
	}
	return newLines
}

func getCommonChar(group [3]string) (rune, error) {
	for _, charA := range group[0] {
		for _, charB := range group[1] {
			for _, charC := range group[2] {
				if charA == charB && charB == charC {
					return charA, nil
				}
			}
		}
	}
	return -1, errors.New("Unable to find matching run")
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

	groups := splitIntoGroups(lines)
	totalPriorities := 0
	for _, group := range groups {
		commonChar, _ := getCommonChar(group)
		priority, _ := getPriority(commonChar)
		totalPriorities += priority
	}
	fmt.Println(totalPriorities)
}
