package day05

import (
	"sort"
	"strconv"
	"strings"

	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day05/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) []string {
	var seats []string
	for _, line := range input {
		seats = append(seats, line)
	}
	return seats
}

func seatids(seats []string) (sid []int) {
	for _, seat := range seats {
		row := seat[:7]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")
		r, _ := strconv.ParseInt(row, 2, 0)
		col := seat[7:len(seat)]
		col = strings.ReplaceAll(col, "R", "1")
		col = strings.ReplaceAll(col, "L", "0")
		c, _ := strconv.ParseInt(col, 2, 0)
		sid = append(sid, int(r*8+c))
	}
	sort.Ints(sid)
	return sid
}

// Solve1 returns answer to first problem
func Solve1() int {
	seats := ParseLines(input)
	seatid := seatids(seats)
	return seatid[len(seatid)-1]
}

// Solve2 returns answer to second problem
func Solve2() int {
	seats := ParseLines(input)
	seatid := seatids(seats)
	s := seatid[1 : len(seatid)-2]
	for i := range s {
		if seatid[i+1] != seatid[i]+1 {
			return seatid[i] + 1
		}
	}
	return -1
}
