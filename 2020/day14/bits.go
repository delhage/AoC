package day14

import (
	"regexp"
	"strconv"
	"strings"

	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day14/input.txt")

//var mem map[int]int

type mw struct {
	a int
	v int
}

// Instr struct of instruction
type Instr struct {
	mask string
	in   []mw
}

var mem map[int]int

// ParseLines parses the file input
func ParseLines(input []string) []Instr {
	var in []Instr
	var mask string
	var i int
	var m []mw
	re := regexp.MustCompile("^mask.*$")
	for j, line := range input {
		if re.MatchString(line) {
			if j > 0 {
				in = append(in, Instr{mask: mask, in: m})
			}
			mask = strings.Split(line, " = ")[1]
			m = []mw{}
			i = 0
		} else {
			parts := strings.Split(line, " = ")
			rgx := `\[(.*?)\]`
			a := regexp.MustCompile(rgx)
			addr, _ := strconv.Atoi(a.FindStringSubmatch(parts[0])[1])
			value, _ := strconv.Atoi(parts[1])
			m = append(m, mw{a: addr, v: value})
			i++
		}
	}
	in = append(in, Instr{mask: mask, in: m})
	return in

}

func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos int) int {
	n &^= (1 << pos)
	return n
}

func maskIt(n int, mask string) int {
	for i, s := range mask {
		b := string(s)
		if b == "X" {
			continue
		}
		i = len(mask) - i - 1
		if b == "0" {
			n = clearBit(n, i)
		}
		if b == "1" {
			n = setBit(n, i)
		}
	}
	return n
}

func maskIt2(n int, mask string) []int {
	var ma []int
	for i, s := range mask {
		i = len(mask) - i - 1
		b := string(s)
		if b == "0" {
			continue
		}
		if b == "1" {
			n = setBit(n, i)
		}
		if b == "X" {
			n = setBit(n, i)
			ma = append(ma, n)
		}
	}
	return ma
}
func setmem(in Instr) {
	for _, i := range in.in {
		mem[i.a] = maskIt(i.v, in.mask)
	}
}

// Solve1 returns answer to first problem
func Solve1() int {
	res := 0
	in := ParseLines(input)
	mem = make(map[int]int)
	for _, i := range in {
		setmem(i)
	}
	for _, m := range mem {
		res += m
	}
	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	return 0
}
