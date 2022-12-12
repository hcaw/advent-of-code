package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	i int
	j int
}

func createGrid(lines []string) ([][]rune, Pos, Pos, []Pos) {
	start, end := Pos{}, Pos{}
	grid := make([][]rune, len(lines))
	aPos := []Pos{}
	for i, line := range lines {
		for j, char := range line {
			grid[i] = append(grid[i], char)
			if char == 'S' {
				start = Pos{i, j}
				grid[i][j] = 'a'
			} else if char == 'E' {
				end = Pos{i, j}
				grid[i][j] = 'z'
			} else if char == 'a' {
				aPos = append(aPos, Pos{i, j})
			}
		}
	}
	return grid, start, end, aPos
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

func breadthFirst(grid [][]rune, start, end Pos) (int, error) {
	frontier := []Pos{}
	frontier = append(frontier, start)
	cameFrom := make(map[Pos]*Pos)
	cameFrom[start] = nil
	endFound := false
	for len(frontier) != 0 {
		curr := frontier[0]
		frontier = frontier[1:]
		if curr == end {
			break
		}
		for _, transform := range transforms {
			next := Pos{curr.i + transform.i, curr.j + transform.j}
			_, beenVisited := cameFrom[next]
			if isInBounds(next, grid) &&
				canBeVisited(curr, next, grid) &&
				!beenVisited {
				frontier = append(frontier, next)
				cameFrom[next] = &curr
				if next == end {
					endFound = true
				}
			}
		}
	}
	if !endFound {
		return -1, errors.New("no solution possible")
	}
	curr := end
	path := []Pos{}
	for curr != start {
		path = append(path, curr)
		curr = *cameFrom[curr]
	}
	return len(path), nil
}

var transforms = []Pos{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	grid, start, end, aPos := createGrid(lines)

	length, _ := breadthFirst(grid, start, end)
	fmt.Println("Solution to problem 1 is", length)

	shortest := length
	for _, a := range aPos {
		length, err := breadthFirst(grid, a, end)
		if err == nil && length < shortest {
			shortest = length
		}
	}
	fmt.Println("Solution to problem 2 is", shortest)
}
