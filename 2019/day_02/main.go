package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OP_PROGRAM_COMPLETION int = 99
	OP_ADDITION           int = 1
	OP_MULTIPLICATION     int = 2
	ANSWER_POSITION       int = 0
	REQUIRED_ANSWER       int = 19690720
	OPERATION_STEP        int = 4
)

type IntcodeProgram struct {
	programCounter int
	integers       []int
}

func (i *IntcodeProgram) Copy() *IntcodeProgram {
	newIntegers := make([]int, len(i.integers))
	copy(newIntegers, i.integers)

	return &IntcodeProgram{
		integers:       newIntegers,
		programCounter: 0,
	}
}

func NewIntcodeProgram(line string) (*IntcodeProgram, error) {
	nums := strings.Split(line, ",")
	integersProgram := make([]int, len(nums))

	for i, num := range nums {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		integersProgram[i] = convertedNum
	}

	return &IntcodeProgram{
		integers:       integersProgram,
		programCounter: 0,
	}, nil
}

func (i *IntcodeProgram) Completed() bool {
	return i.integers[i.programCounter] == OP_PROGRAM_COMPLETION
}

func (i *IntcodeProgram) addition() {
	firstSummandPosition := i.integers[i.programCounter+1]
	firstSummand := i.integers[firstSummandPosition]

	secondSummandPosition := i.integers[i.programCounter+2]
	secondSummand := i.integers[secondSummandPosition]

	ansPosition := i.integers[i.programCounter+3]
	i.integers[ansPosition] = firstSummand + secondSummand
}

func (i *IntcodeProgram) multiplication() {
	firstFactorPosition := i.integers[i.programCounter+1]
	firstFactor := i.integers[firstFactorPosition]

	secondFactorPosition := i.integers[i.programCounter+2]
	secondFactor := i.integers[secondFactorPosition]

	ansPosition := i.integers[i.programCounter+3]
	i.integers[ansPosition] = firstFactor * secondFactor
}

func (i *IntcodeProgram) NextOperation() {
	switch i.integers[i.programCounter] {
	case OP_ADDITION:
		i.addition()
	case OP_MULTIPLICATION:
		i.multiplication()
	}
	i.programCounter += OPERATION_STEP
}

func part1(prog *IntcodeProgram) {
	prog.integers[1] = 12
	prog.integers[2] = 2

	for !prog.Completed() {
		prog.NextOperation()
	}

	output := prog.integers[ANSWER_POSITION]
	fmt.Printf("Part 1: %d\n", output)
}

func part2(prog *IntcodeProgram) {
	found := false
	var output int
	for noun := 0; noun < 100 && !found; noun++ {
		for verb := 0; verb < 100 && !found; verb++ {
			currProg := prog.Copy()
			currProg.integers[1] = noun
			currProg.integers[2] = verb

			for !currProg.Completed() {
				currProg.NextOperation()
			}

			if currProg.integers[ANSWER_POSITION] == REQUIRED_ANSWER {
				found = true
				output = 100*noun + verb
			}
		}
	}

	fmt.Printf("Part 2: %d\n", output)
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

	program, err := NewIntcodeProgram(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	programCopy := program.Copy()
	part1(program)
	part2(programCopy)
}
