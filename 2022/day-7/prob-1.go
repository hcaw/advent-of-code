package main

import (
	"fmt"
	"os"
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

// not my proudest moment...
func createFileTree(lines []string) *node {
	root := &node{name: "/"}
	root.parent = root
	var curr *node = root
	i := 1
	for i < len(lines)-1 {
		line := lines[i]
		if strings.Contains(line, "$ ls") {
			fmt.Println("Found ls on line", i+1)
			for i < len(lines)-1 {
				i += 1
				line := lines[i]
				if strings.Contains(line, "$") {
					fmt.Println("Found $ on line", i+1)
					i -= 1
					break
				}
				var newNode *node
				name := strings.Fields(line)[1]
				if strings.Contains(line, "dir") {
					fmt.Println("Found dir on line", i+1)
					newNode = &node{name: name, parent: curr}
				} else {
					fmt.Println("Found file on line", i+1)
					size, _ := strconv.Atoi(strings.Fields(line)[0])
					newNode = &node{name: name, size: size}
				}
				curr.children = append(curr.children, newNode)
			}
		} else if strings.Contains(line, "$ cd") {
			destination := strings.Fields(line)[2]
			fmt.Println("Found cd on line", i+1, ", destination is", destination)
			if destination == ".." {
				curr = curr.parent
			} else {
				var newNode *node
				for _, n := range curr.children {
					if n.name == destination {
						fmt.Println("Found matching destination of", n.name)
						newNode = n
						break
					}
				}
				curr = newNode
				fmt.Println("New dir is now", curr.name)
				fmt.Println("New parent is now", curr.parent.name)
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
	if n.totalSize != nil && *n.totalSize <= size{
		total += *n.totalSize
	}
	return total
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(input), "\n")

	tree := createFileTree(lines)
	calcSizes(tree)
	total := getDirsBelowSize(tree, 100000)
	fmt.Println(total)
}
