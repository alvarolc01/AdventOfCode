package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LeonardoMonorail struct {
	instructions []string
	pc           int
	registers    map[string]int
}

func NewLeonardoMonorail(instructions []string) *LeonardoMonorail {
	instructionsProgram := make([]string, len(instructions))
	copy(instructionsProgram, instructions)
	return &LeonardoMonorail{
		instructions: instructionsProgram,
		pc:           0,
		registers:    make(map[string]int),
	}
}

func (lm *LeonardoMonorail) GetRegister(n string) int {
	if convertedNum, err := strconv.Atoi(n); err == nil {
		return convertedNum
	}

	return lm.registers[n]
}

func (lm *LeonardoMonorail) Copy() {
	fields := strings.Fields(lm.instructions[lm.pc])

	val := lm.GetRegister(fields[1])
	targetReg := fields[2]

	lm.registers[targetReg] = val
	lm.pc++
}

func (lm *LeonardoMonorail) Increase() {
	fields := strings.Fields(lm.instructions[lm.pc])
	targetReg := fields[1]

	originalValue := lm.GetRegister(targetReg)
	lm.registers[targetReg] = originalValue + 1
	lm.pc++
}

func (lm *LeonardoMonorail) Decrease() {
	fields := strings.Fields(lm.instructions[lm.pc])
	targetReg := fields[1]

	originalValue := lm.GetRegister(targetReg)
	lm.registers[targetReg] = originalValue - 1
	lm.pc++
}

func (lm *LeonardoMonorail) JumpNotZero() {
	fields := strings.Fields(lm.instructions[lm.pc])
	register := fields[1]

	jmpStr := fields[2]
	jmp := lm.GetRegister(jmpStr)

	if lm.GetRegister(register) != 0 {
		lm.pc += jmp
	} else {
		lm.pc++
	}

}

func (lm *LeonardoMonorail) Toggle() {
	fields := strings.Fields(lm.instructions[lm.pc])
	register := fields[1]

	offset := lm.GetRegister(register)
	target := lm.pc + offset
	if target < 0 || target >= len(lm.instructions) {
		lm.pc++
		return
	}

	fieldsNewString := strings.Fields(lm.instructions[lm.pc+offset])
	if len(fieldsNewString) == 2 {
		if fieldsNewString[0] == "inc" {
			fieldsNewString[0] = "dec"
		} else {
			fieldsNewString[0] = "inc"
		}
	} else if len(fieldsNewString) == 3 {
		if fieldsNewString[0] == "jnz" {
			fieldsNewString[0] = "cpy"
		} else {
			fieldsNewString[0] = "jnz"
		}
	}
	lm.instructions[lm.pc+offset] = strings.Join(fieldsNewString, " ")
	lm.pc++
}

func (lm *LeonardoMonorail) ExecuteProgram() {
	for lm.pc >= 0 && lm.pc < len(lm.instructions) {
		ins := lm.instructions[lm.pc]
		if lm.pc+5 < len(lm.instructions) &&
			strings.HasPrefix(lm.instructions[lm.pc], "cpy") &&
			strings.HasPrefix(lm.instructions[lm.pc+1], "inc a") &&
			strings.HasPrefix(lm.instructions[lm.pc+2], "dec c") &&
			strings.HasPrefix(lm.instructions[lm.pc+3], "jnz c -2") &&
			strings.HasPrefix(lm.instructions[lm.pc+4], "dec d") &&
			strings.HasPrefix(lm.instructions[lm.pc+5], "jnz d -5") {

			lm.registers["a"] += lm.registers["b"] * lm.registers["d"]
			lm.registers["c"] = 0
			lm.registers["d"] = 0

			lm.pc += 6
			continue
		} else if strings.HasPrefix(ins, "cpy") {
			lm.Copy()
		} else if strings.HasPrefix(ins, "inc") {
			lm.Increase()
		} else if strings.HasPrefix(ins, "dec") {
			lm.Decrease()
		} else if strings.HasPrefix(ins, "jnz") {
			lm.JumpNotZero()
		} else if strings.HasPrefix(ins, "tgl") {
			lm.Toggle()
		}
	}
}

func part1(commands []string) {
	mono := NewLeonardoMonorail(commands)
	mono.registers["a"] = 7
	mono.ExecuteProgram()

	fmt.Printf("Part 1: %d\n", mono.GetRegister("a"))
}

func part2(commands []string) {
	mono := NewLeonardoMonorail(commands)
	mono.registers["a"] = 12
	mono.ExecuteProgram()

	fmt.Printf("Part 2: %d\n", mono.GetRegister("a"))

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
	commands := strings.Split(string(fileContent), "\n")

	part1(commands)
	part2(commands)
}
