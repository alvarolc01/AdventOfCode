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
	OP_INPUT              int = 3
	OP_OUTPUT             int = 4
	OP_JUMP_IF_TRUE       int = 5
	OP_JUMP_IF_FALSE      int = 6
	OP_LESS_THAN          int = 7
	OP_EQUALS             int = 8
)

var (
	inputValue  int
	outputValue int
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

func (i *IntcodeProgram) addition(indexes []int) {
	i.integers[indexes[2]] = i.integers[indexes[0]] + i.integers[indexes[1]]
	i.programCounter += 4
}

func (i *IntcodeProgram) multiplication(indexes []int) {
	i.integers[indexes[2]] = i.integers[indexes[0]] * i.integers[indexes[1]]
	i.programCounter += 4
}

func (i *IntcodeProgram) input(indexes []int) {
	i.integers[indexes[0]] = inputValue
	i.programCounter += 2
}

func (i *IntcodeProgram) output(indexes []int) {
	outputValue = i.integers[indexes[0]]
	i.programCounter += 2
}

func (i *IntcodeProgram) jumpIfTrue(indexes []int) {
	if i.integers[indexes[0]] != 0 {
		i.programCounter = i.integers[indexes[1]]
	} else {
		i.programCounter += 3
	}
}

func (i *IntcodeProgram) jumpIfFalse(indexes []int) {
	if i.integers[indexes[0]] == 0 {
		i.programCounter = i.integers[indexes[1]]
	} else {
		i.programCounter += 3
	}
}

func (i *IntcodeProgram) lessThan(indexes []int) {
	if i.integers[indexes[0]] < i.integers[indexes[1]] {
		i.integers[indexes[2]] = 1
	} else {
		i.integers[indexes[2]] = 0
	}
	i.programCounter += 4
}

func (i *IntcodeProgram) equals(indexes []int) {
	if i.integers[indexes[0]] == i.integers[indexes[1]] {
		i.integers[indexes[2]] = 1
	} else {
		i.integers[indexes[2]] = 0
	}
	i.programCounter += 4
}

func (i *IntcodeProgram) parseOperation() (int, []int) {
	memVal := i.integers[i.programCounter]
	opCode := memVal % 100

	memVal /= 100
	indexesParameters := make([]int, 0)
	for idx := 1; idx <= 3 && +i.programCounter+idx < len(i.integers); idx++ {
		mode := memVal % 10
		memVal /= 10

		if mode == 1 {
			indexesParameters = append(indexesParameters, i.programCounter+idx)
		} else if mode == 0 {
			indexesParameters = append(indexesParameters, i.integers[i.programCounter+idx])
		}
	}

	return opCode, indexesParameters
}

func (i *IntcodeProgram) NextOperation() {
	opCode, indexes := i.parseOperation()
	switch opCode {
	case OP_ADDITION:
		i.addition(indexes)
	case OP_MULTIPLICATION:
		i.multiplication(indexes)
	case OP_INPUT:
		i.input(indexes)
	case OP_OUTPUT:
		i.output(indexes)
	case OP_JUMP_IF_TRUE:
		i.jumpIfTrue(indexes)
	case OP_JUMP_IF_FALSE:
		i.jumpIfFalse(indexes)
	case OP_LESS_THAN:
		i.lessThan(indexes)
	case OP_EQUALS:
		i.equals(indexes)
	}
}

func part1(prog *IntcodeProgram) {
	inputValue = 1
	for !prog.Completed() {
		prog.NextOperation()
	}

	fmt.Printf("Part 1: %d\n", outputValue)
}

func part2(prog *IntcodeProgram) {
	inputValue = 5
	for !prog.Completed() {
		prog.NextOperation()
	}

	fmt.Printf("Part 2: %d\n", outputValue)
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
