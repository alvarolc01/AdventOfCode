package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	hashDoorsRelation = "UDLR"
	validDoorKeys     = "bcdef"
)

var doorDirections = map[string]Point{
	"U": {row: -1, col: 0},
	"D": {row: 1, col: 0},
	"L": {row: 0, col: -1},
	"R": {row: 0, col: 1},
}

type Point struct {
	row, col int
}

type SearchNode struct {
	position  Point
	pathTaken string
}

func (s *SearchNode) isWithinEdges() bool {
	return s.position.col >= 0 && s.position.row >= 0 && s.position.row <= 3 && s.position.col <= 3
}

func isDoorOpen(doorKey string) bool {
	return strings.Contains(validDoorKeys, doorKey)
}

func (s *SearchNode) neighbourAtDirection(direction string) SearchNode {
	directionMovement := doorDirections[direction]
	return SearchNode{
		position: Point{
			row: s.position.row + directionMovement.row,
			col: s.position.col + directionMovement.col,
		},
		pathTaken: s.pathTaken + string(direction),
	}
}

func (s *SearchNode) GetAccessibleCells(mazePassword string) []SearchNode {
	hashKey := mazePassword + s.pathTaken
	hash := md5Hex(hashKey)

	result := []SearchNode{}
	for idx, doorDirection := range hashDoorsRelation {
		if isDoorOpen(string(hash[idx])) {
			result = append(
				result,
				s.neighbourAtDirection(string(doorDirection)),
			)
		}
	}

	return result
}

func md5Hex(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func exitMaze(mazePassword string, isCompleted func(SearchNode) bool) {
	queue := []SearchNode{{position: Point{0, 0}, pathTaken: ""}}
	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if !currentNode.isWithinEdges() {
			continue
		}

		if currentNode.position.col == 3 && currentNode.position.row == 3 {
			if isCompleted(currentNode) {
				break
			}
			continue
		}
		surroundingPositions := currentNode.GetAccessibleCells(mazePassword)
		queue = append(queue, surroundingPositions...)
	}
}

func part1(mazePassword string) {
	shortestPath := ""
	exitMaze(mazePassword, func(sn SearchNode) bool {
		shortestPath = sn.pathTaken
		return true
	})
	fmt.Printf("Part 1: %s\n", shortestPath)
}

func part2(mazePassword string) {
	longestPath := ""
	exitMaze(mazePassword, func(sn SearchNode) bool {
		if len(longestPath) < len(sn.pathTaken) {
			longestPath = sn.pathTaken
		}

		return false
	})
	fmt.Printf("Part 2: %d\n", len(longestPath))
}

func main() {

	mazePassword := "ihgpwlah"

	part1(mazePassword)
	part2(mazePassword)
}
