package day08

import (
	"strconv"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day08/input.txt")
var program = parseLines(input)
var prg1 = parseLines(input)
var loop = false

// Instruction is an instruction in a program
type Instruction struct {
	op      string
	arg     int
	visited bool
}

func parseLines(input []string) []Instruction {
	var program []Instruction
	for _, line := range input {
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		instr := Instruction{
			op:      op,
			arg:     arg,
			visited: false,
		}
		program = append(program, instr)
	}
	return program
}

// Solve1 returns answer to first problem
func Solve1() int {
	var acc int
	i := 0
	for {
		if program[i].visited {
			loop = true
			break
		}
		program[i].visited = true
		switch program[i].op {
		case "acc":
			acc += program[i].arg
			i++
		case "nop":
			i++
		case "jmp":
			i += program[i].arg
		}
		if i >= len(program) {
			loop = false
			break
		}
	}
	return acc
}

// Solve2 returns answer to second problem
func Solve2() int {
	var acc int
	i := 0
	copy(program, prg1)
	for i < len(program) {
		acc = Solve1()
		if !loop {
			break
		}
		copy(program, prg1)
		switch program[i].op {
		case "jmp":
			program[i].op = "nop"
		case "nop":
			program[i].op = "jmp"
		}
		i++
	}
	return acc
}
