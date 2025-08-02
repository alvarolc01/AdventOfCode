package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	PartOneLowVal  = 17
	PartOneHighVal = 61
)

func initialValues(commands []string) map[int][]int {
	bots := make(map[int][]int)
	for _, command := range commands {
		if strings.HasPrefix(command, "value") {
			var val, bot int
			_, err := fmt.Sscanf(command, "value %d goes to bot %d", &val, &bot)
			if err != nil {
				continue
			}

			bots[bot] = append(bots[bot], val)
		}
	}
	return bots
}

func processCommands(commands []string, bots map[int][]int, outputs map[int]int, compareFunc func(botID, low, high int)) {
	for {
		change := false

		for _, cmd := range commands {
			if !strings.HasPrefix(cmd, "bot") {
				continue
			}

			fields := strings.Fields(cmd)
			if len(fields) < 12 {
				continue
			}

			botID, _ := strconv.Atoi(fields[1])
			lowID, _ := strconv.Atoi(fields[6])
			highID, _ := strconv.Atoi(fields[11])

			values, ok := bots[botID]
			if !ok || len(values) != 2 {
				continue
			}

			slices.Sort(values)
			low, high := values[0], values[1]

			if compareFunc != nil {
				compareFunc(botID, low, high)
			}

			if fields[5] == "bot" {
				bots[lowID] = append(bots[lowID], low)
			} else if outputs != nil {
				outputs[lowID] = low
			}

			if fields[10] == "bot" {
				bots[highID] = append(bots[highID], high)
			} else if outputs != nil {
				outputs[highID] = high
			}

			delete(bots, botID)
			change = true
		}

		if !change {
			break
		}
	}
}

func part1(commands []string) {
	bots := initialValues(commands)
	targetID := -1

	processCommands(commands, bots, nil, func(botID int, low int, high int) {
		if low == PartOneLowVal && high == PartOneHighVal {
			targetID = botID
		}
	})

	fmt.Printf("Part 1: %d\n", targetID)

}

func part2(commands []string) {
	bots := initialValues(commands)
	outputs := make(map[int]int)

	processCommands(commands, bots, outputs, nil)
	ans := outputs[0] * outputs[1] * outputs[2]

	fmt.Printf("Part 2: %d\n", ans)
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
	ipList := strings.Split(string(fileContent), "\n")

	part1(ipList)
	part2(ipList)
}
