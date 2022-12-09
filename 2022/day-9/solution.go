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
		direction := strings.Fields(line)[0]
		steps, _ := strconv.Atoi(strings.Fields(line)[1])
		for i := 0; i < steps; i++ {
			h.x = h.x + transforms[direction].x
			h.y = h.y + transforms[direction].y
			diff := Coords{h.x - t.x, h.y - t.y}
			if abs(diff.x) == 2 || abs(diff.y) == 2 {
				if diff.x == 0 || diff.y == 0 {
					t.x = t.x + transforms[direction].x
					t.y = t.y + transforms[direction].y
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
			}
			visited[t] = struct{}{}
		}
	}
	fmt.Println("Solution to problem 1 is", len(visited))
}
