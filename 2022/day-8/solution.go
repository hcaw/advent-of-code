package main

import (
	"fmt"
	"os"
	"sort"
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
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			num := grid[i][j]
			upVisible := true
			for k := i - 1; k >= 0; k-- {
				height := grid[k][j]
				if num <= height {
					upVisible = false
					break
				}
			}
			leftVisible := true
			for k := j - 1; k >= 0; k-- {
				height := grid[i][k]
				if num <= height {
					leftVisible = false
					break
				}
			}
			downVisible := true
			for k := i + 1; k < len(grid); k++ {
				height := grid[k][j]
				if num <= height {
					downVisible = false
					break
				}
			}
			rightVisible := true
			for k := j + 1; k < len(grid[0]); k++ {
				height := grid[i][k]
				if num <= height {
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
	var scores []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			num := grid[i][j]
			upScore := 0
			for k := i - 1; k >= 0; k-- {
				height := grid[k][j]
				upScore += 1
				if num <= height {
					break
				}
			}
			leftScore := 0
			for k := j - 1; k >= 0; k-- {
				height := grid[i][k]
				leftScore += 1
				if num <= height {
					break
				}
			}
			downScore := 0
			for k := i + 1; k < len(grid); k++ {
				height := grid[k][j]
				downScore += 1
				if num <= height {
					break
				}
			}
			rightScore := 0
			for k := j + 1; k < len(grid[0]); k++ {
				height := grid[i][k]
				rightScore += 1
				if num <= height {
					break
				}
			}
			totalScore := upScore * leftScore * downScore * rightScore
			scores = append(scores, totalScore)
		}
	}
	sort.Ints(scores)
	return scores[len(scores) - 1]
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
