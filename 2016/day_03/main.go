package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName *string) []string {
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file content: %s", err)
		os.Exit(1)
	}

	return strings.Split(string(content), "\n")

}

func parseTriangle(triangleString string) ([]int, error) {
	triangleString = strings.TrimSpace(triangleString)
	vertexes := strings.Fields(triangleString)

	var vertices []int

	for _, vertex := range vertexes {
		vertexInt, err := strconv.Atoi(vertex)
		if err != nil {
			return nil, err
		}
		vertices = append(vertices, vertexInt)
	}

	if len(vertexes) != 3 {
		return nil, fmt.Errorf("expected 3 vertexes, got %d", len(vertexes))
	}

	return vertices, nil
}

func parseTriangles(inputStrings []string) [][]int {
	triangles := make([][]int, len(inputStrings))

	for idx, triangleString := range inputStrings {
		parsedTriangle, valid := parseTriangle(triangleString)
		if valid == nil {
			triangles[idx] = parsedTriangle
		}
	}

	return triangles
}

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

		}
	}
}

func rotateTriangles(triangles [][]int) {
	for i := 0; i < len(triangles); i += 3 {
		rotate(triangles[i : i+3])
	}
}

func isValidTriangle(triangle []int) bool {
	if len(triangle) != 3 {
		return false
	}

	validFirstSide := triangle[0] < triangle[1]+triangle[2]
	validSecondSide := triangle[1] < triangle[0]+triangle[2]
	validThirdSide := triangle[2] < triangle[0]+triangle[1]

	return validFirstSide && validSecondSide && validThirdSide
}

func countValidTriangles(listTriangles [][]int) int {
	validTriangles := 0

	for _, triangle := range listTriangles {
		if isValidTriangle(triangle) {
			validTriangles++
		}
	}

	return validTriangles
}

func part1(listTriangles [][]int) {
	validTriangles := countValidTriangles(listTriangles)

	fmt.Printf("Part 1: %d\n", validTriangles)
}

func part2(listTriangles [][]int) {
	rotateTriangles(listTriangles)
	validTriangles := countValidTriangles(listTriangles)

	fmt.Printf("Part 2: %d\n", validTriangles)
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	inputStrings := readFile(fileName)
	listTriangles := parseTriangles(inputStrings)

	part1(listTriangles)
	part2(listTriangles)

}
