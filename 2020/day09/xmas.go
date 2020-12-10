package day09

import (
	"strconv"

	"../utils"
)

var input, _ = utils.ReadFile("day09/input.txt")

var preamble = 25

func isValid(val int, list []int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == val {
				return true
			}
		}
	}
	return false
}

// ParseLines parses file input
func ParseLines(input []string) []int {
	var data []int
	for _, line := range input {
		v, _ := strconv.Atoi(line)
		data = append(data, v)
	}
	return data
}

// Solve1 returns answer to first problem
func Solve1() int {
	xmas := ParseLines(input)
	for i := preamble; i < len(xmas); i++ {
		val := xmas[i]
		var list []int
		for j := i - preamble; j < i; j++ {
			list = append(list, xmas[j])
		}
		if !isValid(val, list) {
			return xmas[i]
		}

	}
	return 0
}

// Solve2 returns answer to second problem
func Solve2() int {
	xmas := ParseLines(input)
	val := Solve1()
	for i := 0; i < len(xmas)-1; i++ {
		max := xmas[i]
		min := xmas[i]
		acc := xmas[i]
		for j := i + 1; j < len(xmas); j++ {
			if xmas[j] < min {
				min = xmas[j]
			}
			if xmas[j] > max {
				max = xmas[j]
			}
			acc += xmas[j]
			switch {
			case acc == val:
				return min + max
			case acc > val:
				break
			}
		}
	}
	return 0
}
