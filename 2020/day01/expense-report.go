package day01

import (
	"strconv"

	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day01/input.txt")

// ParseLines reads a file into a slice of int:s
func ParseLines(input []string) []int {
	var expense []int
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		expense = append(expense, i)
	}
	return expense
}

// Solve1 returns answer to first problem
func Solve1() int {
	var res int
	x := ParseLines(input)
	for i, s := range x {
		for j := i + 1; j < len(x); j++ {
			res = s + x[j]
			if res == 2020 {
				return s * x[j]
			}
		}
	}
	return -1
}

// Solve2 returns answer to second problem
func Solve2() int {
	var res int
	x := ParseLines(input)
	for i, s := range x {
		for j := i + 1; j < len(x); j++ {
			for k := j + i; k < len(x); k++ {
				res = s + x[j] + x[k]
				if res == 2020 {
					return s * x[j] * x[k]
				}
			}
		}
	}
	return -1
}
