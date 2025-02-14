package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

func part1(memoryLocation int) {
	circle := 0
	for int(math.Pow(float64(2*circle+1), 2)) < memoryLocation {
		circle++
	}

	circleBase := int(math.Pow(float64(2*circle-1), 2))

	distance := memoryLocation
	for side := 0; side < 4; side++ {
		currDistance := int(math.Abs(float64(circleBase + (circle * (2*side + 1)) - memoryLocation)))

		if currDistance < distance {
			distance = currDistance
		}
	}

	fmt.Printf("Part 1: %d\n", distance+circle)
}

type Point struct {
	x int
	y int
}

func getNextPosition(currentPos Point) Point {
	x, y := currentPos.x, currentPos.y

	if x == 0 && y == 0 {
		return Point{1, 0}
	} else if y >= x && y > -x {
		return Point{x - 1, y}
	} else if -x >= y && -x > -y {
		return Point{x, y - 1}
	} else if x > -y && x > y {
		return Point{x, y + 1}
	} else if x >= y && -x >= y {
		return Point{x + 1, y}
	}

	return currentPos
}

func getSumSurrounding(point Point, foundPoints map[Point]int) int {
	sumSurrounding := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if val, found := foundPoints[Point{point.x + i, point.y + j}]; found {
				sumSurrounding += val
			}
		}
	}

	return sumSurrounding
}

func part2(memoryLocation int) {
	currentPoint := Point{0, 0}
	foundPoints := map[Point]int{currentPoint: 1}

	for val := 1; val <= memoryLocation; {
		nextPosition := getNextPosition(currentPoint)
		sumSurrounding := getSumSurrounding(nextPosition, foundPoints)

		foundPoints[nextPosition] = sumSurrounding
		val = sumSurrounding
		currentPoint = nextPosition
	}

	fmt.Printf("Part 2: %d\n", foundPoints[currentPoint])
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := string(fileContent)

	inputInt, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Failed to parse input string: %s\n", err)
		os.Exit(1)
	}

	part1(inputInt)
	part2(inputInt)
}
