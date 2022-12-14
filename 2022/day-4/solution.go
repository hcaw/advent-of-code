package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

type Pair struct {
	a Section
	b Section
}

func getSection(secStr string) Section {
	arr := strings.Split(secStr, "-")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])
	return Section{start, end}
}

func getPairs(lines []string) []Pair {
	var pairs []Pair
	for _, line := range lines {
		splitLine := strings.SplitN(line, ",", 2)
		pair := Pair{getSection(splitLine[0]), getSection(splitLine[1])}
		pairs = append(pairs, pair)
	}
	return pairs
}

func isContained(pair Pair) bool {
	return (pair.a.start >= pair.b.start && pair.a.end <= pair.b.end) ||
		(pair.b.start >= pair.a.start && pair.b.end <= pair.a.end)
}

func overlaps(pair Pair) bool {
	return (pair.a.start >= pair.b.start && pair.a.start <= pair.b.end) ||
		(pair.b.start >= pair.a.start && pair.b.start <= pair.a.end)
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	pairs := getPairs(lines)
	totalContained := 0
	totalOverlapping := 0
	for _, pair := range pairs {
		if isContained(pair) {
			totalOverlapping += 1
		}
		if overlaps(pair) {
			totalContained += 1
		}
	}
	fmt.Println("Solution to problem 1:",totalOverlapping)
	fmt.Println("Solution to problem 2:",totalContained)
}
