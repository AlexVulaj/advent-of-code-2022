package main

import (
	"advent-of-code-2022/util"
	_ "embed"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	size     int
	parent   *Node
	children []*Node
}

//go:embed input.txt
var input string

func main() {
	if util.ParsePartFlag() == 1 {
		util.PrintResult(p1(input))
	} else {
		util.PrintResult(p2(input))
	}
}

func p1(input string) int {
	root, dirs := createTree(input)
	calculateSizes(&root)

	total := 0
	for _, dir := range dirs {
		if dir.size < 100000 {
			total += dir.size
		}
	}
	return total
}

func p2(input string) int {
	root, dirs := createTree(input)
	calculateSizes(&root)

	fileSystemSize := 70000000
	currentFreeSpace := fileSystemSize - root.size
	updateSize := 30000000
	minSize := fileSystemSize
	
	for _, dir := range dirs {
		if dir.size+currentFreeSpace >= updateSize {
			if dir.size < minSize {
				minSize = dir.size
			}
		}
	}
	return minSize
}

func createTree(input string) (Node, []*Node) {
	root := Node{
		name:     "/",
		size:     0,
		parent:   nil,
		children: []*Node{},
	}
	dirs := []*Node{&root}

	var current *Node

	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "$" {
			if splitLine[1] == "cd" {
				if splitLine[2] == "/" {
					current = &root
				} else if splitLine[2] == ".." {
					current = current.parent
				} else {
					child, found := searchChildren(current, splitLine[2])
					if !found {
						child = &Node{
							name:     splitLine[2],
							size:     0,
							parent:   current,
							children: []*Node{},
						}
						dirs = append(dirs, child)
						current.children = append(current.children, child)
					}
					current = child
				}
			}
		} else {
			var newNode Node
			if splitLine[0] == "dir" {
				newNode = Node{
					name:     splitLine[1],
					size:     0,
					parent:   current,
					children: []*Node{},
				}
				dirs = append(dirs, &newNode)
			} else {
				size, _ := strconv.Atoi(splitLine[0])
				newNode = Node{
					name:     splitLine[1],
					size:     size,
					parent:   current,
					children: nil,
				}
			}
			current.children = append(current.children, &newNode)
		}
	}
	return root, dirs
}

func searchChildren(parent *Node, targetName string) (*Node, bool) {
	for _, node := range parent.children {
		if node.name == targetName {
			return node, true
		}
	}
	return nil, false
}

func calculateSizes(root *Node) int {
	if root.size != 0 {
		return root.size
	}
	size := 0
	for _, child := range root.children {
		size += calculateSizes(child)
	}
	root.size = size
	return size
}
