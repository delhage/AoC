package day06

import (
	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day06/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) ([]map[byte]int, []int) {
	var num []int
	var i int
	var qlist []map[byte]int
	qmap := make(map[byte]int)
	for _, line := range input {
		if line != "" {
			bs := []byte(line)
			for _, b := range bs {
				qmap[b]++
			}
			i++
		} else {
			qlist = append(qlist, qmap)
			num = append(num, i)
			qmap = make(map[byte]int)
			i = 0
		}
	}
	qlist = append(qlist, qmap)
	num = append(num, i)

	return qlist, num
}

// Solve1 returns answer to first problem
func Solve1() int {
	var i int
	qlist, _ := ParseLines(input)

	for _, q := range qlist {
		i += len(q)
	}
	return i
}

// Solve2 returns answer to second problem
func Solve2() int {
	var i int
	qlist, num := ParseLines(input)

	for l, q := range qlist {
		for k := range q {
			if q[k] == num[l] {
				i++
			}
		}
	}
	return i
}
