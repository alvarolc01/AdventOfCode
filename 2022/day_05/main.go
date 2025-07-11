package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Move struct {
	Count, From, To int
}

func NewMove(line string) (*Move, error) {
	var count, from, to int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
	if err != nil {
		return nil, err
	}

	return &Move{
		Count: count,
		From:  from - 1,
		To:    to - 1,
	}, nil
}

type Crane struct {
	stacks    [][]rune
	movements []Move
}

func parseStacks(block string) ([][]rune, error) {
	lines := strings.Split(block, "\n")
	lineNums := lines[len(lines)-1]
	lineNums = strings.TrimSpace(lineNums)
	nums := strings.Fields(lineNums)
	stacks := make([][]rune, len(nums))

	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < len(line); j += 4 {
			if j+1 >= len(line) {
				continue
			}
			ch := line[j+1]

			if ch != ' ' {
				stackIndex := j / 4
				stacks[stackIndex] = append(stacks[stackIndex], rune(ch))
			}
		}
	}

	return stacks, nil
}

func NewCrane(input string) (*Crane, error) {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) != 2 {
		return nil, fmt.Errorf("invalid format, expected 2 blocks")
	}

	stakcs, _ := parseStacks(blocks[0])

	moves := strings.Split(blocks[1], "\n")
	movesCranes := make([]Move, len(moves))
	for i, m := range moves {
		mov, err := NewMove(m)
		if err != nil {
			return nil, err
		}

		movesCranes[i] = *mov
	}

	return &Crane{
		stacks:    stakcs,
		movements: movesCranes,
	}, nil
}

func (c *Crane) ExecuteMovements() {
	for _, mov := range c.movements {
		for range mov.Count {
			crate := c.stacks[mov.From][len(c.stacks[mov.From])-1]
			c.stacks[mov.From] = c.stacks[mov.From][:len(c.stacks[mov.From])-1]
			c.stacks[mov.To] = append(c.stacks[mov.To], crate)
		}
	}
}

func (c *Crane) ExecuteMovementsModernModel() {
	for _, mov := range c.movements {
		fromStack := c.stacks[mov.From]
		crates := fromStack[len(fromStack)-mov.Count:]
		c.stacks[mov.From] = fromStack[:len(fromStack)-mov.Count]
		c.stacks[mov.To] = append(c.stacks[mov.To], crates...)
	}
}

func (c *Crane) TopCrates() string {
	var b strings.Builder
	for _, stack := range c.stacks {
		if len(stack) > 0 {
			b.WriteRune(stack[len(stack)-1])
		}
	}
	return b.String()
}

func part1(crane *Crane) {

	crane.ExecuteMovements()
	output := crane.TopCrates()

	fmt.Printf("Part 1: %s\n", output)
}

func part2(crane *Crane) {

	crane.ExecuteMovementsModernModel()
	output := crane.TopCrates()

	fmt.Printf("Part 2: %s\n", output)
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
	crane, err := NewCrane(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	part1(crane)

	crane, err = NewCrane(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	part2(crane)
}
