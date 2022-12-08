package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createGrid(lines []string) [][]int {
	rows := len(lines)
	grid := make([][]int, rows)
	for i, line := range lines {
		for _, char := range line {
			val, _ := strconv.Atoi(string(char))
			grid[i] = append(grid[i], val)
		}
	}
	return grid
}

func findTotalVisible(grid [][]int) int {
	// Length of sides
	visible := len(grid) * 2
	// Length of remaining top and bottom
	visible += (len(grid[0]) - 2) * 2
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			num := grid[r][c]
			upVisible := true
			for k := r - 1; k >= 0; k-- {
				if num <= grid[k][c] {
					upVisible = false
					break
				}
			}
			leftVisible := true
			for k := c - 1; k >= 0; k-- {
				if num <= grid[r][k] {
					leftVisible = false
					break
				}
			}
			downVisible := true
			for k := r + 1; k < len(grid); k++ {
				if num <= grid[k][c] {
					downVisible = false
					break
				}
			}
			rightVisible := true
			for k := c + 1; k < len(grid[0]); k++ {
				if num <= grid[r][k] {
					rightVisible = false
					break
				}
			}
			if upVisible || leftVisible || downVisible || rightVisible {
				visible += 1
			}
		}
	}
	return visible
}

func findHighestScenicScore(grid [][]int) int {
	var highest int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			num := grid[r][c]
			var upScore, leftScore, downScore, rightScore int
			for i := r - 1; i >= 0; i-- {
				upScore += 1
				if num <= grid[i][c] {
					break
				}
			}
			for i := c - 1; i >= 0; i-- {
				leftScore += 1
				if num <= grid[r][i] {
					break
				}
			}
			for i := r + 1; i < len(grid); i++ {
				downScore += 1
				if num <= grid[i][c] {
					break
				}
			}
			for i := c + 1; i < len(grid[0]); i++ {
				rightScore += 1
				if num <= grid[r][i] {
					break
				}
			}
			totalScore := upScore * leftScore * downScore * rightScore
			if (totalScore > highest) {
				highest = totalScore
			}
		}
	}
	return highest
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")
	grid := createGrid(lines)
	visible := findTotalVisible(grid)
	fmt.Println("Solution to problem 1", visible)
	highest := findHighestScenicScore(grid)
	fmt.Println("Solution to problem 2", highest)
}
