package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reOr     = regexp.MustCompile(`^(\w+) OR (\w+)$`)
	reAnd    = regexp.MustCompile(`^(\w+) AND (\w+)$`)
	reLShift = regexp.MustCompile(`^(\w+) LSHIFT (\d+)$`)
	reRShift = regexp.MustCompile(`^(\w+) RSHIFT (\d+)$`)
	reNot    = regexp.MustCompile(`^NOT (\w+)$`)
)

func getOrSignals(command string) (string, string) {
	match := reOr.FindStringSubmatch(command)

	return match[1], match[2]
}

func getAndSignals(command string) (string, string) {
	match := reAnd.FindStringSubmatch(command)

	return match[1], match[2]
}

func getLShiftSignals(command string) (string, int) {
	match := reLShift.FindStringSubmatch(command)

	firstOperand := match[1]
	numShiftsStr := match[2]
	numShifts, _ := strconv.Atoi(numShiftsStr)
	return firstOperand, numShifts
}

func getRShiftSignals(command string) (string, int) {
	match := reRShift.FindStringSubmatch(command)

	firstOperand := match[1]
	numShiftsStr := match[2]
	numShifts, _ := strconv.Atoi(numShiftsStr)
	return firstOperand, numShifts
}

func getNotSignal(command string) string {
	match := reNot.FindStringSubmatch(command)

	return match[1]
}

func createRequiredInstructions(input []string) map[string]string {
	instructions := make(map[string]string)
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			continue
		}

		instructions[parts[1]] = parts[0]
	}
	return instructions
}

func getResultingSignal(instructions map[string]string, targetSignal string, values map[string]uint16) uint16 {
	if val, ok := values[targetSignal]; ok {
		return val
	}

	getValue := func(id string) uint16 {
		numVal, err := strconv.ParseUint(id, 10, 16)
		if err != nil {
			return getResultingSignal(instructions, id, values)
		}
		return uint16(numVal)
	}

	requiredInstruction := instructions[targetSignal]

	var result uint16
	if strings.Contains(requiredInstruction, "AND") {
		firstWire, secondWire := getAndSignals(requiredInstruction)
		result = getValue(firstWire) & getValue(secondWire)
	} else if strings.Contains(requiredInstruction, "OR") {
		firstWire, secondWire := getOrSignals(requiredInstruction)
		result = getValue(firstWire) | getValue(secondWire)
	} else if strings.Contains(requiredInstruction, "LSHIFT") {
		wire, numShifts := getLShiftSignals(requiredInstruction)
		result = getValue(wire) << numShifts
	} else if strings.Contains(requiredInstruction, "RSHIFT") {
		wire, numShifts := getRShiftSignals(requiredInstruction)
		result = getValue(wire) >> numShifts
	} else if strings.Contains(requiredInstruction, "NOT") {
		wire := getNotSignal(requiredInstruction)
		result = ^getValue(wire)
	} else {
		result = getValue(requiredInstruction)
	}

	values[targetSignal] = result
	return result
}

func part1(input []string) {
	requiredInstructions := createRequiredInstructions(input)

	targetWire := "a"
	targetWireValue := getResultingSignal(requiredInstructions, targetWire, map[string]uint16{})

	fmt.Printf("Part 1: %d\n", targetWireValue)
}

func part2(input []string) {
	requiredInstructions := createRequiredInstructions(input)

	targetWire := "a"
	targetWireValue := getResultingSignal(requiredInstructions, targetWire, map[string]uint16{})
	targetWireValue = getResultingSignal(requiredInstructions, targetWire, map[string]uint16{"b": targetWireValue})

	fmt.Printf("Part 2: %d\n", targetWireValue)
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
