package day13

import (
	"sort"
	"strconv"
	"strings"

	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day13/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) (int, []int) {
	e, _ := strconv.Atoi(input[0])
	var i int
	var buses []int
	for _, entry := range strings.Split(input[1], ",") {
		if entry == "x" {
			i = -1
		} else {
			i, _ = strconv.Atoi(entry)
		}
		buses = append(buses, i)
	}
	return e, buses
}

func getBus(e int, ids []int) int {
	var t []int
	var id = make(map[int]int)
	for _, i := range ids {
		l := 0
		if i >= 0 {
			for l < e {
				l += i
			}
			t = append(t, l-e)
			id[l-e] = i
		}

	}
	sort.Ints(t)
	return t[0] * id[t[0]]
}

func getTime(ids []int) int {
	var t int
	var inc = ids[0]
	for i := 1; i < len(ids); i++ {
		id := ids[i]
		if id < 0 {
			continue
		}
		for {
			t += inc
			if (t+i)%id == 0 {
				inc *= id
				break
			}
		}
	}
	return t
}

// Solve1 returns answer to first problem
func Solve1() int {
	e, ids := ParseLines(input)
	return getBus(e, ids)
}

// Solve2 returns answer to second problem
func Solve2() int {
	_, ids := ParseLines(input)
	return getTime(ids)
}
