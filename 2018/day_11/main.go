package main

import (
	"fmt"
	"math"
)

const (
	side             int = 300
	gridSerialNumber int = 18
)

func calculateFuelCell(x, y int) int {
	rackID := x + 10
	powerLevel := (rackID*y + gridSerialNumber) * rackID
	return (powerLevel/100)%10 - 5
}

func generateGrid() [side + 1][side + 1]int {
	var grid, sum [side + 1][side + 1]int
	for y := 1; y <= side; y++ {
		for x := 1; x <= side; x++ {
			grid[y][x] = calculateFuelCell(x, y)
			sum[y][x] = grid[y][x] + sum[y-1][x] + sum[y][x-1] - sum[y-1][x-1]
		}
	}
	return sum
}

func squareSum(sum *[side + 1][side + 1]int, x, y, size int) int {
	x2, y2 := x+size-1, y+size-1
	return sum[y2][x2] - sum[y-1][x2] - sum[y2][x-1] + sum[y-1][x-1]
}

func part1(sum *[side + 1][side + 1]int) {
	maxPower := math.MinInt
	var ansX, ansY int
	for x := 1; x <= side-2; x++ {
		for y := 1; y <= side-2; y++ {
			total := squareSum(sum, x, y, 3)
			if total > maxPower {
				maxPower = total
				ansX, ansY = x, y
			}
		}
	}
	fmt.Printf("Part 1: %d,%d\n", ansX, ansY)
}

func part2(sum *[side + 1][side + 1]int) {
	maxPower := math.MinInt
	var ansX, ansY, ansSize int
	for size := 1; size <= side; size++ {
		for x := 1; x <= side-size+1; x++ {
			for y := 1; y <= side-size+1; y++ {
				total := squareSum(sum, x, y, size)
				if total > maxPower {
					maxPower = total
					ansX, ansY, ansSize = x, y, size
				}
			}
		}
	}
	fmt.Printf("Part 2: %d,%d,%d\n", ansX, ansY, ansSize)
}

func main() {
	sum := generateGrid()
	part1(&sum)
	part2(&sum)
}
