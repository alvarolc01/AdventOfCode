package main

import (
	"container/heap"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	totalRisk int
	x, y      int
	index     int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].totalRisk < pq[j].totalRisk }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getMinRisk(initialMap []string, expansion int) int {
	baseRows, baseCols := len(initialMap), len(initialMap[0])
	rows, cols := baseRows*expansion, baseCols*expansion

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{totalRisk: 0, x: 0, y: 0})

	found := make(map[[2]int]bool)

	for pq.Len() > 0 {
		top := heap.Pop(pq).(*Node)
		if found[[2]int{top.x, top.y}] {
			continue
		}
		found[[2]int{top.x, top.y}] = true

		if top.x == rows-1 && top.y == cols-1 {
			return top.totalRisk
		}

		for _, dir := range directions {
			nx, ny := top.x+dir[0], top.y+dir[1]
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols {
				continue
			}

			squaresDistance := nx/baseRows + ny/baseCols
			riskIncrease := (int(initialMap[nx%baseRows][ny%baseRows]-'0')+squaresDistance-1)%9 + 1
			heap.Push(pq, &Node{
				totalRisk: top.totalRisk + riskIncrease,
				x:         nx,
				y:         ny,
			})

		}

	}

	return -1
}

func part1(initialMap []string) {
	fmt.Printf("Part 1: %d\n", getMinRisk(initialMap, 1))
}

func part2(initialMap []string) {
	fmt.Printf("Part 2: %d\n", getMinRisk(initialMap, 5))
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
	lines := strings.Split(string(fileContent), "\n")

	part1(lines)
	part2(lines)
}
