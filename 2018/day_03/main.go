package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Claim struct {
	id            int
	start         Point
	width, height int
}

func NewClaim(line string) (*Claim, error) {
	var id, width, height, startX, startY int
	_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &startX, &startY, &width, &height)
	if err != nil {
		return nil, err
	}

	return &Claim{
		id: id,
		start: Point{
			x: startX,
			y: startY,
		},
		width:  width,
		height: height,
	}, nil
}

func fillClaimMap(claimMap map[Point][]int, claim *Claim) {
	for i := 0; i < claim.width; i++ {
		for j := 0; j < claim.height; j++ {
			p := Point{claim.start.x + i, claim.start.y + j}
			claimMap[p] = append(claimMap[p], claim.id)
		}
	}
}

func generateClaimsMap(claims []*Claim) (map[Point][]int, map[int]bool) {
	claimsMap := make(map[Point][]int)
	idsInMap := make(map[int]bool)
	for _, claim := range claims {
		fillClaimMap(claimsMap, claim)
		idsInMap[claim.id] = true
	}
	return claimsMap, idsInMap
}

func part1(claimsMap map[Point][]int) {

	countPointsWithMultipleClaims := 0
	for _, val := range claimsMap {
		if len(val) >= 2 {
			countPointsWithMultipleClaims++
		}
	}

	fmt.Printf("Part 1: %d\n", countPointsWithMultipleClaims)

}

func part2(claimsMap map[Point][]int, idsInMap map[int]bool) {

	for _, val := range claimsMap {
		if len(val) >= 2 {
			for _, id := range val {
				delete(idsInMap, id)
			}
		}
	}

	nonOverlappingID := -1
	for key := range idsInMap {
		nonOverlappingID = key
	}

	fmt.Printf("Part 2: %d\n", nonOverlappingID)

}

func parseInput(lines []string) ([]*Claim, error) {
	output := make([]*Claim, len(lines))

	for idx, line := range lines {
		claim, err := NewClaim(line)
		if err != nil {
			return nil, err
		}

		output[idx] = claim
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
	input := strings.Split(string(fileContent), "\n")
	claims, err := parseInput(input)
	if err != nil {
		fmt.Println("error creating graph:", err)
		os.Exit(1)
	}
	claimsMap, idsMap := generateClaimsMap(claims)

	part1(claimsMap)
	part2(claimsMap, idsMap)
}
