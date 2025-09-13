package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	instructions     []*Instruction
	idx, acc         int
	seenInstructions map[int]bool
}

func (p *Program) Copy() *Program {
	nextIns := make([]*Instruction, len(p.instructions))
	for i, ins := range p.instructions {
		nextIns[i] = &Instruction{
			op:    ins.op,
			value: ins.value,
		}
	}
	return &Program{
		instructions:     nextIns,
		seenInstructions: make(map[int]bool),
	}
}

func (p *Program) Run() bool {
	for p.idx >= 0 && p.idx < len(p.instructions) {
		if _, ok := p.seenInstructions[p.idx]; ok {
			return false
		}
		p.seenInstructions[p.idx] = true

		ins := p.instructions[p.idx]
		if ins.op == "nop" {
			p.idx++
		} else if ins.op == "jmp" {
			p.idx += ins.value
		} else if ins.op == "acc" {
			p.acc += ins.value
			p.idx++
		}
	}

	return true
}

type Instruction struct {
	op    string
	value int
}

func NewInstruction(line string) (*Instruction, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid instruction: %s", line)
	}

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &Instruction{op: parts[0], value: value}, nil
}

func part1(prog *Program) {
	prog.Run()
	fmt.Printf("Part 1: %d\n", prog.acc)
}

func part2(prog *Program) {
	for i := range prog.instructions {
		if prog.instructions[i].op == "acc" {
			continue
		}

		copyProg := prog.Copy()
		if copyProg.instructions[i].op == "nop" {
			copyProg.instructions[i].op = "jmp"
		} else {
			copyProg.instructions[i].op = "nop"
		}

		if completedExecution := copyProg.Run(); completedExecution {
			fmt.Println("Part 2:", copyProg.acc)
			break
		}
	}
}

func parseInput(lines []string) (*Program, error) {
	output := &Program{
		seenInstructions: make(map[int]bool),
		idx:              0,
		acc:              0,
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
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(program)
	part2(program)
}
