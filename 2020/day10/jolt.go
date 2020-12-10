package day10

import (
	"sort"
	"strconv"

	"../utils"
)

var input, _ = utils.ReadFile("day10/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) []int {
	var data []int
	for _, line := range input {
		v, _ := strconv.Atoi(line)
		data = append(data, v)
	}
	sort.Ints(data)
	return data
}

// Solve1 returns answer to first problem
func Solve1() int {
	data := ParseLines(input)
	var ones int
	var threes int
	switch data[0] {
	case 1:
		ones++
	case 3:
		threes++
	}
	data = append(data, data[len(data)-1]+3)
	for i := 0; i < len(data)-1; i++ {
		switch data[i+1] - data[i] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}
	//fmt.Println(data, ones, threes)
	return ones * threes
}

func numsets(data []int) int {
	switch len(data) {
	case 0, 1, 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	case 5:
		return 7
	}
	i := 0
	for i < len(data)-1 {
		if data[i+1] == data[i]+1 {
			i++
		} else {
			break
		}
	}
	if len(data[:i]) == 0 {
		return numsets(data[i+1:])
	} else {
		return numsets(data[:i+1]) * numsets(data[i+1:])
	}
}

// Solve2 returns answer to second problem
func Solve2() int {
	data := ParseLines(input)
	data = append(data, data[len(data)-1]+3)
	data = append([]int{0}, data...)
	return numsets(data)
}
