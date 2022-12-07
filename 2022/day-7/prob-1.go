package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	name      string
	size      int
	children  []*node
	parent    *node
	totalSize *int
}

func isListCommand(str string) bool {
	return str == "$ ls"
}

func isCdCommand(str string) bool {
	return strings.HasPrefix(str, "$ cd")
}

func containsPrompt(str string) bool {
	return strings.HasPrefix(str, "$")
}

func isDirectoryListing(str string) bool {
	return strings.HasPrefix(str, "dir")
}

// not my proudest moment...
func createFileTree(lines []string) *node {
	root := &node{name: "/"}
	root.parent = root
	curr := root
	i := 1
	for i < len(lines)-1 {
		line := lines[i]
		if isListCommand(line) {
			for i < len(lines)-1 {
				i += 1
				line := lines[i]
				if containsPrompt(line) {
					i -= 1
					break
				}
				var newNode *node
				name := strings.Fields(line)[1]
				if isDirectoryListing(line) {
					newNode = &node{name: name, parent: curr}
				} else {
					size, _ := strconv.Atoi(strings.Fields(line)[0])
					newNode = &node{name: name, size: size}
				}
				curr.children = append(curr.children, newNode)
			}
		} else if isCdCommand(line) {
			destination := strings.Fields(line)[2]
			if destination == ".." {
				curr = curr.parent
			} else {
				var newNode *node
				for _, n := range curr.children {
					if n.name == destination {
						newNode = n
						break
					}
				}
				curr = newNode
			}
		}
		i += 1
	}
	return root
}

func calcSizes(n *node) int {
	total := 0
	for _, child := range n.children {
		if child.parent == nil {
			total += child.size
		} else {
			total += calcSizes(child)
		}
	}
	n.totalSize = &total
	return total
}

func getDirsBelowSize(n *node, size int) int {
	total := 0
	for _, child := range n.children {
		if child.parent != nil {
			total += getDirsBelowSize(child, size)
		}
	}
	if n.totalSize != nil && *n.totalSize <= size {
		total += *n.totalSize
	}
	return total
}

func getAllDirSizes(n *node) []int {
	var sizes []int
	for _, child := range n.children {
		if child.parent != nil {
			childSizes := getAllDirSizes(child)
			sizes = append(sizes, childSizes...)
		}
	}
	if n.totalSize != nil {
		sizes = append(sizes, *n.totalSize)
	}
	return sizes
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	tree := createFileTree(lines)
	calcSizes(tree)
	total := getDirsBelowSize(tree, 100000)
	fmt.Println("Problem 1 answer:", total)

	const requiredAvailable = 30000000
	const diskSpace = 70000000
	diskUsed := *tree.totalSize
	diskAvailable := diskSpace - diskUsed
	sizeToDelete := requiredAvailable - diskAvailable

	dirSizes := getAllDirSizes(tree)
	sort.Ints(dirSizes)

	for _, size := range dirSizes {
		if size >= sizeToDelete {
			fmt.Println("Problem 2 answer:", size)
			return
		}
	}
}
