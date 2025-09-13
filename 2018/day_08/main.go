package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []*Node
	metadata []int
}

func (n *Node) SumMetadata() int {
	result := 0
	for _, child := range n.children {
		result += child.SumMetadata()
	}

	for _, metadataNum := range n.metadata {
		result += metadataNum
	}

	return result
}

func (n *Node) leafValue() int {
	sum := 0
	for _, m := range n.metadata {
		sum += m
	}
	return sum
}

func (n *Node) nonLeafValue() int {
	sum := 0
	for _, m := range n.metadata {
		idx := m - 1
		if idx >= 0 && idx < len(n.children) {
			sum += n.children[idx].RootValue()
		}
	}
	return sum
}

func (n *Node) RootValue() int {
	if len(n.children) == 0 {
		return n.leafValue()
	}
	return n.nonLeafValue()
}

func NewNode(nums []int, idx int) (*Node, int, error) {
	numChildren, numMetadata := nums[idx], nums[idx+1]
	currentIdx := idx + 2
	result := Node{}

	for i := 0; i < numChildren; i++ {
		child, nextIdx, err := NewNode(nums, currentIdx)
		if err != nil {
			return nil, 0, err
		}
		result.children = append(result.children, child)
		currentIdx = nextIdx
	}

	result.metadata = nums[currentIdx : currentIdx+numMetadata]
	nextStart := currentIdx + numMetadata
	return &result, nextStart, nil
}

func part1(root *Node) {

	fmt.Printf("Part 1: %d\n", root.SumMetadata())
}

func part2(root *Node) {

	fmt.Printf("Part 2: %d\n", root.RootValue())
}

func parseInput(line string) (*Node, error) {
	fields := strings.Fields(line)
	nums := make([]int, 0, len(fields))

	for _, f := range fields {
		convN, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		nums = append(nums, convN)
	}

	root, _, err := NewNode(nums, 0)
	return root, err
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}
	input := string(fileContent)
	root, err := parseInput(input)
	if err != nil {
		fmt.Println("error creating graph:", err)
		os.Exit(1)
	}

	part1(root)
	part2(root)
}
