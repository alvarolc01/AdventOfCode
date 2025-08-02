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
	return &LeonardoMonorail{
		instructions: instructions,
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
	jmp, err := strconv.Atoi(jmpStr)
	if err != nil {
		return
	}

	if lm.GetRegister(register) != 0 {
		lm.pc += jmp
	} else {
		lm.pc++
	}

}

func (lm *LeonardoMonorail) ExecuteProgram() {
	for lm.pc >= 0 && lm.pc < len(lm.instructions) {
		ins := lm.instructions[lm.pc]

		if strings.HasPrefix(ins, "cpy") {
			lm.Copy()
		} else if strings.HasPrefix(ins, "inc") {
			lm.Increase()
		} else if strings.HasPrefix(ins, "dec") {
			lm.Decrease()
		} else if strings.HasPrefix(ins, "jnz") {
			lm.JumpNotZero()
		}
	}
}

func part1(commands []string) {
	mono := NewLeonardoMonorail(commands)
	mono.ExecuteProgram()

	fmt.Printf("Part 1: %d\n", mono.GetRegister("a"))

}

func part2(commands []string) {
	mono := NewLeonardoMonorail(commands)
	mono.registers["c"] = 1
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
