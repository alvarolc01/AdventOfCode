package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RaceTime int = 2503

type Reindeer struct {
	name        string
	speed       int
	flyingTime  int
	restingTime int
}

func (r *Reindeer) DistanceAfterTime(seconds int) int {
	cycleTime := r.flyingTime + r.restingTime
	completeTurns := seconds / cycleTime
	remainingTime := min(seconds%cycleTime, r.flyingTime)
	timeFlying := completeTurns*r.flyingTime + remainingTime

	return timeFlying * r.speed
}

func NewReindeer(line string) *Reindeer {
	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 5 {
		return nil
	}

	speed, _ := strconv.Atoi(matches[2])
	flyingTime, _ := strconv.Atoi(matches[3])
	restingTime, _ := strconv.Atoi(matches[4])

	return &Reindeer{
		name:        matches[1],
		speed:       speed,
		flyingTime:  flyingTime,
		restingTime: restingTime,
	}

}

func part1(input []*Reindeer) {
	maxDistance := 0
	for _, reindeer := range input {
		distanceReindeer := reindeer.DistanceAfterTime(RaceTime)
		maxDistance = max(maxDistance, distanceReindeer)
	}

	fmt.Printf("Part 1: %d\n", maxDistance)
}

func part2(input []*Reindeer) {
	scores := make([]int, len(input))

	for time := 1; time <= RaceTime; time++ {
		distances := make([]int, len(input))
		maxDistanceAtTime := 0

		for idx, reindeer := range input {
			distances[idx] = reindeer.DistanceAfterTime(time)
			maxDistanceAtTime = max(maxDistanceAtTime, distances[idx])
		}

		for idx := range distances {
			if distances[idx] == maxDistanceAtTime {
				scores[idx]++
			}
		}
	}

	winningScore := 0
	for _, score := range scores {
		winningScore = max(winningScore, score)
	}

	fmt.Printf("Part 2: %d\n", winningScore)
}

func parseInput(input []string) []*Reindeer {
	var result []*Reindeer
	for _, line := range input {
		reindeer := NewReindeer(line)
		if reindeer != nil {
			result = append(result, reindeer)
		}
	}
	return result
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
	input := strings.Split(string(fileContent), "\n")
	listReindeer := parseInput(input)

	part1(listReindeer)
	part2(listReindeer)
}
