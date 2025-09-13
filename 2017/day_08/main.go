package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	instructions []*Instruction
	registers    map[string]int
	maxValue     int
}

func (p *Program) EvaluateInstruction(idx int) {
	if idx < 0 || idx >= len(p.instructions) {
		return
	}

	inst := p.instructions[idx]
	if !inst.HoldsCondition(p.registers) {
		return
	}

	p.registers[inst.register] += inst.amountChange

	if val := p.registers[inst.register]; val > p.maxValue {
		p.maxValue = val
	}
}

func (p *Program) Run() {
	for idx := range p.instructions {
		p.EvaluateInstruction(idx)
	}
}

type Instruction struct {
	register           string
	amountChange       int
	conditionLeftSide  string
	conditionRightSide int
	conditionOperation string
}

func (i *Instruction) HoldsCondition(registers map[string]int) bool {
	leftVal := registers[i.conditionLeftSide]

	if i.conditionOperation == ">" {
		return leftVal > i.conditionRightSide
	} else if i.conditionOperation == "<" {
		return leftVal < i.conditionRightSide
	} else if i.conditionOperation == "<=" {
		return leftVal <= i.conditionRightSide
	} else if i.conditionOperation == ">=" {
		return leftVal >= i.conditionRightSide
	} else if i.conditionOperation == "==" {
		return leftVal == i.conditionRightSide
	} else if i.conditionOperation == "!=" {
		return leftVal != i.conditionRightSide
	}
	return true
}

func NewInstruction(line string) (*Instruction, error) {
	fields := strings.Fields(line)

	amountChange, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, err
	}

	if fields[1] == "dec" {
		amountChange = -amountChange
	}

	conditionRightSide, err := strconv.Atoi(fields[6])
	if err != nil {
		return nil, err
	}

	return &Instruction{
		register:           fields[0],
		amountChange:       amountChange,
		conditionLeftSide:  fields[4],
		conditionOperation: fields[5],
		conditionRightSide: conditionRightSide,
	}, nil
}

func part1(prog *Program) {

	prog.Run()

	maxVal := math.MinInt
	for _, val := range prog.registers {
		if val > maxVal {
			maxVal = val
		}
	}
	fmt.Printf("Part 1: %d\n", maxVal)
}

func part2(prog *Program) {
	fmt.Printf("Part 2: %d\n", prog.maxValue)
}

func parseInput(lines []string) (*Program, error) {
	output := &Program{
		registers: make(map[string]int),
		maxValue:  math.MinInt,
	}
	for _, line := range lines {
		ins, err := NewInstruction(line)
		if err != nil {
			return nil, err
		}

		output.instructions = append(output.instructions, ins)
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

	program, err := parseInput(input)
	if err != nil {
		fmt.Println("error parting input:", err)
		os.Exit(1)
	}

	part1(program)
	part2(program)
}
