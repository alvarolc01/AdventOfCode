package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func (r *Race) TimesRecordSurpassed() int {
	timesSurpassed := 0

	for i := 0; i <= r.Time; i++ {
		distanceRace := i * (r.Time - i)
		if distanceRace > r.Distance {
			timesSurpassed++
		}
	}

	return timesSurpassed
}

func NewRace(raceTime, dist string) (*Race, error) {
	t, err := strconv.Atoi(raceTime)
	if err != nil {
		return nil, err
	}

	d, err := strconv.Atoi(dist)
	if err != nil {
		return nil, err
	}

	return &Race{
		Time:     t,
		Distance: d,
	}, nil
}

func part1(races []*Race) {
	mulAns := 1

	for _, r := range races {
		count := r.TimesRecordSurpassed()
		mulAns *= count
	}

	fmt.Printf("Part 1: %d\n", mulAns)
}

func appendNumber(original, num int) int {
	mult := 1
	for n := num; n > 0; n /= 10 {
		mult *= 10
	}

	return original*mult + num
}

func part2(races []*Race) {
	totalTime := 0
	totalDist := 0

	for _, race := range races {
		totalTime = appendNumber(totalTime, race.Time)
		totalDist = appendNumber(totalDist, race.Distance)
	}

	updatedRace := &Race{
		Time:     totalTime,
		Distance: totalDist,
	}
	ans := updatedRace.TimesRecordSurpassed()
	fmt.Printf("Part 2: %d\n", ans)
}

func parseInput(input []string) ([]*Race, error) {
	time := strings.Split(input[0], ":")[1]
	racesTimes := strings.Fields(strings.TrimSpace(time))

	distance := strings.Split(input[1], ":")[1]
	racesDistances := strings.Fields(strings.TrimSpace(distance))

	output := make([]*Race, len(racesDistances))
	for idx := range racesDistances {
		race, err := NewRace(racesTimes[idx], racesDistances[idx])
		if err != nil {
			return nil, err
		}
		output[idx] = race
	}

	return output, nil
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
	input := strings.Split(string(fileContent), "\n")

	races, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(races)
	part2(races)
}
