package day15

import (
	"strconv"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day15/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) []int {
	//w := make(map[int]int)
	var w []int
	for _, s := range strings.Split(input[0], ",") {
		i, _ := strconv.Atoi(string(s))
		w = append(w, i)
	}
	return w
}

func speak(w []int, n int) int {
	var last int
	seen := make(map[int]int)
	i := 0
	for _, l := range w {
		i++
		seen[l] = i
		last = l
	}
	delete(seen, last)
	//fmt.Println(seen)
	for i < n {
		now := 0
		if val, ok := seen[last]; !ok {
			now = 0
		} else {
			now = i - val
		}
		seen[last] = i
		last = now
		i++
	}
	return last
}

// Solve1 returns answer to first problem
func Solve1() int {
	w := ParseLines(input)
	return speak(w, 2020)
}

// Solve2 returns answer to second problem
func Solve2() int {
	w := ParseLines(input)
	return speak(w, 30000000)
}
