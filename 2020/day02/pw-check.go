package day02

import (
	"bytes"
	"strconv"
	"strings"

	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day02/input.txt")

// Pw is a structure of password policy and password
type Pw struct {
	min int
	max int
	c   []byte
	pw  string
}

// ParseLines parses the file input
func ParseLines(input []string) []Pw {
	var pwlist []Pw
	for _, line := range input {
		parts := strings.Fields(line)
		k := strings.Split(parts[0], "-")
		c := []byte(strings.Split(parts[1], ":")[0])
		min, _ := strconv.Atoi(k[0])
		max, _ := strconv.Atoi(k[1])
		pw := parts[2]
		pwlist = append(pwlist, Pw{min: min, max: max, c: c, pw: pw})

	}
	return pwlist
}

func validpw1(pass Pw) int {
	b := []byte(pass.pw)
	num := bytes.Count(b, pass.c)
	if num > pass.max {
		return 0
	}
	if num < pass.min {
		return 0
	}
	return 1
}

func validpw2(pass Pw) int {
	b := []byte(pass.pw)
	var res int = 0
	if b[pass.min-1] == pass.c[0] {
		res++
	}
	if b[pass.max-1] == pass.c[0] {
		res++
	}
	if res == 1 {
		return 1
	}
	return 0

}

// Solve1 returns answer to first problem
func Solve1() int {
	var res int = 0
	pwlist := ParseLines(input)
	for _, pw := range pwlist {
		res += validpw1(pw)
	}
	return res

}

// Solve2 returns answer to second problem
func Solve2() int {
	var res int = 0
	pwlist := ParseLines(input)
	for _, pw := range pwlist {
		res += validpw2(pw)
	}
	return res
}
