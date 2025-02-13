package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TuringLock struct {
	instructions   []string
	programCounter int
	registers      map[string]int
}

func NewTuringLock(instructions []string) *TuringLock {
	return &TuringLock{
		instructions:   instructions,
		registers:      make(map[string]int),
		programCounter: 0,
	}
}

func (t *TuringLock) GetRegister(registerName string) int {
	return t.registers[registerName]
}

func (t *TuringLock) halfInstruction() {
	re := regexp.MustCompile(`^hlf (\w+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	registerName := match[1]
	originalValue := t.registers[registerName]

	t.registers[registerName] = originalValue / 2
	t.programCounter++
}

func (t *TuringLock) tripleInstruction() {
	re := regexp.MustCompile(`^tpl (\w+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	registerName := match[1]
	originalValue := t.registers[registerName]

	t.registers[registerName] = originalValue * 3
	t.programCounter++
}

func (t *TuringLock) incrementInstruction() {
	re := regexp.MustCompile(`^inc (\w+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	registerName := match[1]
	originalValue := t.registers[registerName]

	t.registers[registerName] = originalValue + 1
	t.programCounter++
}

func (t *TuringLock) jumpInstruction() {
	re := regexp.MustCompile(`^jmp ([+-]?\d+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	offsetStr := match[1]
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return
	}

	t.programCounter += offset
}

func (t *TuringLock) jumpIfEventInstruction() {
	re := regexp.MustCompile(`^jie (\w+), ([+-]?\d+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	registerName := match[1]
	offsetStr := match[2]
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return
	}
	if t.registers[registerName]%2 == 0 {
		t.programCounter += offset
	} else {
		t.programCounter++
	}

}

func (t *TuringLock) jumpIfOneInstruction() {
	re := regexp.MustCompile(`^jio (\w+), ([+-]?\d+)$`)
	match := re.FindStringSubmatch(t.instructions[t.programCounter])

	registerName := match[1]
	offsetStr := match[2]
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return
	}
	if t.registers[registerName] == 1 {
		t.programCounter += offset
	} else {
		t.programCounter++
	}
}

func (t *TuringLock) ExecuteProgram() {
	for t.programCounter >= 0 && t.programCounter < len(t.instructions) {
		currentInstruction := t.instructions[t.programCounter]
		if strings.HasPrefix(currentInstruction, "hlf") {
			t.halfInstruction()
		} else if strings.HasPrefix(currentInstruction, "tpl") {
			t.tripleInstruction()
		} else if strings.HasPrefix(currentInstruction, "inc") {
			t.incrementInstruction()
		} else if strings.HasPrefix(currentInstruction, "jmp") {
			t.jumpInstruction()
		} else if strings.HasPrefix(currentInstruction, "jie") {
			t.jumpIfEventInstruction()
		} else if strings.HasPrefix(currentInstruction, "jio") {
			t.jumpIfOneInstruction()
		}
	}
}

func part1(instructions []string) {
	turingLock := NewTuringLock(instructions)
	turingLock.ExecuteProgram()

	fmt.Printf("Part 1: %d\n", turingLock.GetRegister("b"))
}

func part2(instructions []string) {
	turingLock := NewTuringLock(instructions)
	turingLock.registers["a"] = 1
	turingLock.ExecuteProgram()

	fmt.Printf("Part 2: %d\n", turingLock.GetRegister("b"))
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

	part1(input)
	part2(input)
}
