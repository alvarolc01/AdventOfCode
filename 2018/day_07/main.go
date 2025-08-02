package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
)

const (
	NumWorkers       = 5
	BaseTaskDuration = 60
)

func part1(graph map[rune][]rune, precedingTasks map[rune]int) {
	output := []rune{}
	queue := []rune{}

	for key := range graph {
		if precedingTasks[key] == 0 {
			queue = append(queue, key)
		}
	}
	slices.Sort(queue)

	done := make(map[rune]bool)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if _, ok := done[top]; ok {
			continue
		}
		done[top] = true
		output = append(output, top)

		for _, dep := range graph[top] {
			precedingTasks[dep]--
			if precedingTasks[dep] == 0 {
				queue = append(queue, dep)
			}
		}

		slices.Sort(queue)
	}

	fmt.Printf("Part 1: %s\n", string(output))

}

func part2(graph map[rune][]rune, copyPrecedingTasks map[rune]int) {
	queue := []rune{}

	for key := range graph {
		if copyPrecedingTasks[key] == 0 {
			queue = append(queue, key)
		}
	}
	slices.Sort(queue)

	workers := make([]rune, NumWorkers)
	for i := range NumWorkers {
		workers[i] = '0'
	}
	times := make([]int, NumWorkers)
	done := make(map[rune]bool)
	time := 0
	idle := 0

	for ; idle < NumWorkers; time++ {
		idle = NumWorkers
		for i := range NumWorkers {
			if times[i] != 0 {
				idle--
				times[i]--
			} else if workers[i] != '0' {
				done[workers[i]] = true
				for _, dep := range graph[workers[i]] {
					copyPrecedingTasks[dep]--
					if copyPrecedingTasks[dep] == 0 {
						queue = append(queue, dep)
					}
				}
				workers[i] = '0'
			}
		}

		slices.Sort(queue)

		for len(queue) > 0 && idle > 0 {
			top := queue[0]
			queue = queue[1:]
			found := false
			for i := 0; i < NumWorkers && !found; i++ {
				if workers[i] == '0' {
					idle--
					found = true
					workers[i] = top
					times[i] = int(top) - 'A' + BaseTaskDuration
				}
			}

			if !found {
				queue = append(queue, top)
			}
		}
	}

	fmt.Printf("Part 2: %d\n", time-1)

}

func generatePrecedencesGraph(lines []string) (map[rune][]rune, map[rune]int, error) {
	output := make(map[rune][]rune)
	precedingTasks := make(map[rune]int)

	for _, line := range lines {
		var start, end rune
		_, err := fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &start, &end)
		if err != nil {
			return nil, nil, err
		}

		if _, ok := output[start]; !ok {
			output[start] = make([]rune, 0)
		}
		output[start] = append(output[start], end)
		precedingTasks[end]++
	}

	for _, val := range output {
		slices.Sort(val)
	}

	return output, precedingTasks, nil
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
	precedences, precedingTasks, err := generatePrecedencesGraph(input)
	if err != nil {
		fmt.Println("error creating graph:", err)
		os.Exit(1)
	}

	part1(precedences, maps.Clone(precedingTasks))
	part2(precedences, maps.Clone(precedingTasks))
}
