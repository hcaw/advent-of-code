package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func calcVisitedPositions(lines []string, length int) int {
	knots := make([]Coords, length)
	visited := make(map[Coords]struct{})
	visited[Coords{}] = struct{}{}
	for _, line := range lines {
		direction := strings.Fields(line)[0]
		steps, _ := strconv.Atoi(strings.Fields(line)[1])
		for i := 0; i < steps; i++ {
			knots[0].x = knots[0].x + transforms[direction].x
			knots[0].y = knots[0].y + transforms[direction].y
			for j := 1; j < length; j++ {
				h, t := &knots[j-1], &knots[j]
				diff := Coords{h.x - t.x, h.y - t.y}
				if abs(diff.x) == 2 || abs(diff.y) == 2 {
					if diff.x == 0 || diff.y == 0 {
						t.x = t.x + (diff.x / 2)
						t.y = t.y + (diff.y / 2)
					} else {
						if diff.x > 0 {
							t.x = t.x + 1
						} else {
							t.x = t.x - 1
						}
						if diff.y > 0 {
							t.y = t.y + 1
						} else {
							t.y = t.y - 1
						}
					}
				} else {
					break
				}
			}
			visited[knots[len(knots)-1]] = struct{}{}
		}
	}
	return len(visited)
}

var transforms = map[string]Coords{"L": {-1, 0}, "R": {1, 0}, "U": {0, 1}, "D": {0, -1}}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	prob1 := calcVisitedPositions(lines, 2)
	fmt.Println("Solution to problem 2 is", prob1)
	prob2 := calcVisitedPositions(lines, 10)
	fmt.Println("Solution to problem 2 is", prob2)
}
