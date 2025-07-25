package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Guard struct {
	nappingTimes []int
	napStart     int
	id           int
}

func (g *Guard) timeSleeping() int {
	totalSleepingTime := 0

	for _, daysAsleepAtMin := range g.nappingTimes {
		totalSleepingTime += daysAsleepAtMin
	}

	return totalSleepingTime
}

func (g *Guard) fallsAsleep(minute int) {
	g.napStart = minute
}

func (g *Guard) wakesUp(minute int) {
	for i := g.napStart; i < minute; i++ {
		g.nappingTimes[i]++
	}
}

func (g *Guard) minMostFrequestlyAsleep() int {
	minMostAsleep := 0

	for idx, val := range g.nappingTimes {
		if val > g.nappingTimes[minMostAsleep] {
			minMostAsleep = idx
		}
	}

	return minMostAsleep
}

type GuardTeam struct {
	guards        map[int]*Guard
	activeGuardID int
}

func (g *GuardTeam) ChangeShift(log string) {
	var guardId int
	commandParts := strings.Split(log, "]")
	_, err := fmt.Sscanf(commandParts[1], " Guard #%d begins shift", &guardId)
	if err != nil {
		return
	}

	g.activeGuardID = guardId

	if _, ok := g.guards[guardId]; !ok {
		g.guards[guardId] = &Guard{id: guardId, nappingTimes: make([]int, 60), napStart: 0}
	}

}

func (g *GuardTeam) GuardFallsAsleep(minute int) {
	guard, ok := g.guards[g.activeGuardID]
	if !ok {
		return
	}

	guard.fallsAsleep(minute)
}

func (g *GuardTeam) GuardWakesUp(minute int) {
	guard, ok := g.guards[g.activeGuardID]
	if !ok {
		return
	}

	guard.wakesUp(minute)
}

func part1(patrolTeam *GuardTeam) {
	var mostAsleepGuard *Guard
	maxSleepTime := -1

	for _, currentGuard := range patrolTeam.guards {
		currentSleep := currentGuard.timeSleeping()
		if currentSleep > maxSleepTime {
			mostAsleepGuard = currentGuard
			maxSleepTime = currentSleep
		}
	}

	mostAsleepMinute := mostAsleepGuard.minMostFrequestlyAsleep()
	ans := mostAsleepGuard.id * mostAsleepMinute

	fmt.Printf("Part 1: %d \n", ans)
}

func part2(patrolTeam *GuardTeam) {
	var mostAsleepGuard *Guard
	maxSleepTime := -1
	mostAsleepMin := -1

	for _, currentGuard := range patrolTeam.guards {
		currMinMostAsleep := currentGuard.minMostFrequestlyAsleep()
		if maxSleepTime < currentGuard.nappingTimes[currMinMostAsleep] {
			mostAsleepGuard = currentGuard
			mostAsleepMin = currMinMostAsleep
			maxSleepTime = currentGuard.nappingTimes[currMinMostAsleep]
		}
	}

	ans := mostAsleepMin * mostAsleepGuard.id
	fmt.Printf("Part 2: %d \n", ans)
}

func getMinute(log string) (int, error) {
	var year, month, day, hour, minute int
	var action string

	_, err := fmt.Sscanf(log, "[%d-%d-%d %d:%d] %s", &year, &month, &day, &hour, &minute, &action)
	return minute, err

}

func parseInput(guardLogs []string) (*GuardTeam, error) {
	slices.Sort(guardLogs)
	guardsPatrol := GuardTeam{
		guards:        make(map[int]*Guard),
		activeGuardID: 0,
	}

	for _, activity := range guardLogs {
		minuteCommand, err := getMinute(activity)
		if err != nil {
			return nil, err
		}

		if strings.HasSuffix(activity, "begins shift") {
			guardsPatrol.ChangeShift(activity)
		} else if strings.HasSuffix(activity, "falls asleep") {
			guardsPatrol.GuardFallsAsleep(minuteCommand)
		} else if strings.HasSuffix(activity, "wakes up") {
			guardsPatrol.GuardWakesUp(minuteCommand)
		}

	}

	return &guardsPatrol, nil
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
	guardTeam, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(guardTeam)
	part2(guardTeam)
}
