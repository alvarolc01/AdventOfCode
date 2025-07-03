package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

type Interval struct {
	Start int
	End   int
}

type Almanac struct {
	seeds       []int
	conversions [][]Range
}

func getRange(line string) (*Range, error) {
	nums := strings.Fields(line)

	DestinationStart, err := strconv.Atoi(nums[0])
	if err != nil {
		return nil, err
	}
	SourceStart, err := strconv.Atoi(nums[1])
	if err != nil {
		return nil, err
	}
	Length, err := strconv.Atoi(nums[2])
	if err != nil {
		return nil, err
	}

	return &Range{
		DestinationStart: DestinationStart,
		SourceStart:      SourceStart,
		Length:           Length,
	}, nil
}

func createConversionStep(conversion string) ([]Range, error) {
	parts := strings.Split(strings.TrimSpace(conversion), "\n")[1:]
	ran := make([]Range, len(parts))

	for idx, part := range parts {
		r, err := getRange(part)
		if err != nil {
			return nil, err
		}
		ran[idx] = *r
	}

	return ran, nil
}

func applyStepToValue(value int, step []Range) int {
	for _, r := range step {
		if r.SourceStart <= value && value < r.SourceStart+r.Length {
			return value - r.SourceStart + r.DestinationStart
		}
	}
	return value
}

func part1(alm *Almanac) {
	minLocation := math.MaxInt

	for _, seed := range alm.seeds {
		loc := seed
		for _, step := range alm.conversions {
			loc = applyStepToValue(loc, step)
		}
		minLocation = min(minLocation, loc)
	}

	fmt.Printf("Part 1: %d\n", minLocation)
}

func applyStepToIntervals(intervals []Interval, step []Range) []Interval {
	var result []Interval

	for _, interval := range intervals {
		queue := []Interval{interval}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			matched := false

			for _, currRange := range step {
				rangeStart := currRange.SourceStart
				rangeEnd := currRange.SourceStart + currRange.Length

				if current.End <= rangeStart || current.Start >= rangeEnd {
					continue
				}
				matched = true

				if current.Start < rangeStart {
					queue = append(queue, Interval{current.Start, rangeStart})
				}

				overlapStart := max(current.Start, rangeStart)
				overlapEnd := min(current.End, rangeEnd)
				offset := currRange.DestinationStart - currRange.SourceStart
				result = append(result, Interval{overlapStart + offset, overlapEnd + offset})

				if rangeEnd < current.End {
					queue = append(queue, Interval{rangeEnd, current.End})
				}

				break
			}

			if !matched {
				result = append(result, current)
			}
		}
	}

	return result
}

func part2(alm *Almanac) {
	var intervals []Interval

	for i := 0; i < len(alm.seeds)-1; i += 2 {
		start := alm.seeds[i]
		length := alm.seeds[i+1]
		intervals = append(intervals, Interval{start, start + length})
	}

	for _, step := range alm.conversions {
		intervals = applyStepToIntervals(intervals, step)
	}

	minLocation := math.MaxInt
	for _, interval := range intervals {
		minLocation = min(minLocation, interval.Start)
	}

	fmt.Printf("Part 2: %d\n", minLocation)
}

func parseInput(input string) (*Almanac, error) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	seedsLine := strings.TrimSpace(parts[0])
	seedsNums := strings.Fields(seedsLine)[1:]
	seeds := make([]int, len(seedsNums))
	for idx, seedStr := range seedsNums {
		seedNum, err := strconv.Atoi(seedStr)
		if err != nil {
			return nil, err
		}
		seeds[idx] = seedNum
	}

	conversions := make([][]Range, len(parts[1:]))
	for idx, block := range parts[1:] {
		step, err := createConversionStep(block)
		if err != nil {
			return nil, err
		}
		conversions[idx] = step
	}

	return &Almanac{
		seeds:       seeds,
		conversions: conversions,
	}, nil
}

func main() {
	fileName := flag.String("file", "", "Path to input file")
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

	almanac, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(almanac)
	part2(almanac)
}
