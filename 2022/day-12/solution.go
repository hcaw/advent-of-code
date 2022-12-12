package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	i int
	j int
}

func createGrid(lines []string) ([][]rune, Pos, Pos) {
	start, end := Pos{}, Pos{}
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		for j, char := range line {
			grid[i] = append(grid[i], char)
			if char == 'S' {
				start = Pos{i, j}
				grid[i][j] = 'a'
			} else if char == 'E' {
				end = Pos{i, j}
				grid[i][j] = 'z'
			}
		}
	}
	return grid, start, end
}

func isInBounds(pos Pos, grid [][]rune) bool {
	return pos.i >= 0 &&
		pos.i < len(grid) &&
		pos.j >= 0 &&
		pos.j < len(grid[0])
}

func canBeVisited(curr, next Pos, grid [][]rune) bool {
	return grid[curr.i][curr.j]+1 >= grid[next.i][next.j]
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	grid, start, end := createGrid(lines)
	transforms := []Pos{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	frontier := []Pos{}
	frontier = append(frontier, start)
	cameFrom := make(map[Pos]*Pos)
	cameFrom[start] = nil

	for len(frontier) != 0 {
		curr := frontier[0]
		frontier = frontier[1:]
		for _, transform := range transforms {
			next := Pos{curr.i + transform.i, curr.j + transform.j}
			_, beenVisited := cameFrom[next]
			if isInBounds(next, grid) &&
				canBeVisited(curr, next, grid) &&
				!beenVisited {
					frontier = append(frontier, next)
					cameFrom[next] = &curr
			}
		}
	}

	curr := end
	path := []Pos{}
	for curr != start {
		path = append(path, curr)
		curr = *cameFrom[curr]
	}

	fmt.Println("Solution to problem 1 is", len(path))
}
