package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	LayerWidth  = 3
	LayerHeight = 2
)

type Layer [LayerHeight][LayerWidth]int

func CountNum(layer Layer, num int) int {
	count := 0

	for _, row := range layer {
		for _, val := range row {
			if val == num {
				count++
			}
		}
	}

	return count
}

func NewLayer(nums string) (Layer, error) {
	var output Layer
	for i := 0; i < LayerHeight; i++ {
		for j := 0; j < LayerWidth; j++ {
			idx := i*LayerWidth + j
			num, err := strconv.Atoi(string(nums[idx]))
			if err != nil {
				return Layer{}, err
			}
			output[i][j] = num
		}
	}
	return output, nil
}

func part1(layers []Layer) {
	minZeroes := math.MaxInt
	result := 0

	for _, l := range layers {
		if countZero := CountNum(l, 0); countZero < minZeroes {
			minZeroes = countZero
			result = CountNum(l, 1) * CountNum(l, 2)
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func composeImage(layers []Layer) []int {
	const transparent int = 2
	img := make([]int, LayerHeight*LayerWidth)
	for i := range LayerHeight * LayerWidth {
		img[i] = -1
	}

	for _, l := range layers {
		for idxRow, row := range l {
			for idxCol, char := range row {
				if img[idxRow*LayerWidth+idxCol] == -1 && char != transparent {
					img[idxRow*LayerWidth+idxCol] = char
				}
			}
		}
	}

	return img
}

func renderImage(img []int) {
	for i := range LayerHeight {
		for j := range LayerWidth {
			if img[i*LayerWidth+j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

}

func part2(layers []Layer) {
	img := composeImage(layers)

	fmt.Println("Part 2:")
	renderImage(img)
}

func parseInput(input string) ([]Layer, error) {
	output := make([]Layer, 0, len(input)/(LayerHeight*LayerWidth))
	for i := 0; i < len(input); i += LayerWidth * LayerHeight {
		currLayer, err := NewLayer(input[i : i+LayerWidth*LayerHeight])
		if err != nil {
			return nil, err
		}
		output = append(output, currLayer)
	}
	return output, nil
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

	layers, err := parseInput(input)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}
	part1(layers)
	part2(layers)
}
