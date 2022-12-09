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

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	transforms := map[string]Coords{"L": {-1, 0}, "R": {1, 0}, "U": {0, 1}, "D": {0, -1}}
	h, t := Coords{}, Coords{}
	visited := make(map[Coords]struct{})
	visited[t] = struct{}{}
	for _, line := range lines {
		fmt.Println("Line is", line)
		direction := strings.Fields(line)[0]
		steps, _ := strconv.Atoi(strings.Fields(line)[1])
		fmt.Println("start h is", h)
		for i := 0; i < steps; i++ {
			h.x = h.x + transforms[direction].x
			h.y = h.y + transforms[direction].y
			fmt.Println("New head is", h)
			diff := Coords{h.x - t.x, h.y - t.y}
			fmt.Println("Difference is", diff)
			if (abs(diff.x) == 2 && diff.y == 0) || (abs(diff.y) == 2 && diff.x == 0) {
				t.x = t.x + transforms[direction].x
				t.y = t.y + transforms[direction].y
				fmt.Println("Found two away, tail is now", t)
			} else if (abs(diff.x) == 2 && abs(diff.y) == 1) || (abs(diff.y) == 2 && abs(diff.x) == 1) {
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
				fmt.Println("Found diagonal, t is now", t)
			}
			visited[t] = struct{}{}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("Solution to problem 1 is", len(visited))
}
