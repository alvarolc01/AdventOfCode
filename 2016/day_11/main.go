package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Item struct {
	chip bool
	name string
}

type Floor []Item

type Building struct {
	lift   int
	floors []Floor
}

type State struct {
	floors []Floor
	lift   int
	steps  int
}

func cloneFloors(floors []Floor) []Floor {
	cloned := make([]Floor, len(floors))
	for i := range floors {
		cloned[i] = make(Floor, len(floors[i]))
		copy(cloned[i], floors[i])
	}
	return cloned
}

func toString(floors []Floor, lift int) string {
	type listItem struct {
		floor  int
		isChip bool
	}

	positions := []listItem{}
	for i, floor := range floors {
		for _, item := range floor {
			li := listItem{}
			li.isChip = item.chip
			li.floor = i
			positions = append(positions, li)
		}
	}

	sort.Slice(positions, func(i, j int) bool {
		if positions[i].isChip == positions[j].isChip {
			return positions[i].floor < positions[j].floor
		}
		return positions[i].isChip
	})

	result := fmt.Sprintf("Lift%d_", lift)
	for _, p := range positions {
		result += fmt.Sprintf("%d,%t_", p.floor, p.isChip)
	}
	return result
}

func (f Floor) isSafe() bool {
	gens := map[string]bool{}
	for _, i := range f {
		if !i.chip {
			gens[i.name] = true
		}
	}

	if len(gens) == 0 {
		return true
	}

	for _, i := range f {
		if i.chip && !gens[i.name] {
			return false
		}
	}

	return true
}

func (f Floor) getPermutations() [][]Item {
	var output [][]Item
	for i := range f {
		output = append(output, []Item{f[i]})
		for j := i + 1; j < len(f); j++ {
			output = append(output, []Item{f[i], f[j]})
		}
	}
	return output
}

func (f Floor) removeItems(toRemove []Item) Floor {
	rem := make(map[Item]bool)
	for _, i := range toRemove {
		rem[i] = true
	}
	var result Floor
	for _, i := range f {
		if !rem[i] {
			result = append(result, i)
		}
	}
	return result
}

func isFinished(floors []Floor) bool {
	for i := range len(floors) - 1 {
		if len(floors[i]) > 0 {
			return false
		}
	}
	return true
}

func stepsToComplete(building *Building) int {
	queue := []State{}
	processed := map[string]bool{}

	initState := State{
		floors: cloneFloors(building.floors),
		lift:   0,
		steps:  0,
	}
	queue = append(queue, initState)
	processed[toString(initState.floors, initState.lift)] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if isFinished(curr.floors) {
			return curr.steps
		}

		permutations := curr.floors[curr.lift].getPermutations()
		for _, dir := range []int{1, -1} {
			newLift := curr.lift + dir
			if newLift < 0 || newLift >= len(curr.floors) {
				continue
			}
			for _, perm := range permutations {
				nextFloors := cloneFloors(curr.floors)
				nextFloors[curr.lift] = nextFloors[curr.lift].removeItems(perm)
				nextFloors[newLift] = append(nextFloors[newLift], perm...)

				if nextFloors[curr.lift].isSafe() && nextFloors[newLift].isSafe() {
					key := toString(nextFloors, newLift)
					if !processed[key] {
						processed[key] = true
						queue = append(queue, State{
							floors: nextFloors,
							lift:   newLift,
							steps:  curr.steps + 1,
						})
					}
				}
			}
		}
	}

	return -1
}

func part1(b *Building) {
	ans := stepsToComplete(b)
	fmt.Printf("Part 1: %d\n", ans)
}

func part2(b *Building) {
	b.floors[0] = append(b.floors[0], Item{name: "elerium", chip: false},
		Item{name: "elerium", chip: true}, Item{name: "dilithium", chip: false}, Item{name: "dilithium", chip: true})
	ans := stepsToComplete(b)
	fmt.Printf("Part 2: %d\n", ans)

}

func NewFloor(line string) (Floor, error) {
	output := []Item{}

	fields := strings.Fields(line)
	for i, f := range fields {
		if f == "a" {
			it := Item{
				name: strings.Split(fields[i+1], "-")[0],
				chip: fields[i+2][0] != 'g',
			}
			output = append(output, it)
		}
	}

	return output, nil
}

func parseInput(lines []string) (*Building, error) {
	building := &Building{}
	for _, line := range lines {
		floor, err := NewFloor(line)
		if err != nil {
			return nil, err
		}
		building.floors = append(building.floors, floor)
	}
	return building, nil
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
	building, err := parseInput(ipList)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(building)
	part2(building)
}
