package day09

import (
	"aoc/2023/utils"
	"strconv"
	"strings"
)

var input, _ = utils.ReadFile("day09/input.txt")

func stringToInts(s string) []int {
	var res []int
	t := strings.Fields(s)
	for _, c := range t {
		d, _ := strconv.Atoi(c)
		res = append(res, d)
	}
	return res
}

func intDiff(a []int) []int {
	var res []int
	for i := 0; i < len(a)-1; i++ {
		res = append(res, a[i+1]-a[i])
	}
	return res
}

func allZeros(a []int) bool {
	for _, n := range a {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextHistory(a []int) int {
	var res int
	var matr [][]int
	matr = append(matr, a)
	z := false
	t := a
	for !z {
		t = intDiff(t)
		matr = append(matr, t)
		z = allZeros(t)
	}

	for i := len(matr) - 2; i >= 0; i-- {
		res = res + matr[i][len(matr[i])-1]
	}

	return res
}

func NextHistory2(a []int) int {
	var res int
	var matr [][]int
	matr = append(matr, a)
	z := false
	t := a
	for !z {
		t = intDiff(t)
		matr = append(matr, t)
		z = allZeros(t)
	}

	for i := len(matr) - 2; i >= 0; i-- {
		res = matr[i][0] - res
	}

	return res
}

// Solve1 returns answer to first problem
func Solve1() int {
	i := 0
	for _, s := range input {
		i += nextHistory(stringToInts(s))
	}
	return i
}

// Solve2 returns answer to second problem
func Solve2() int {
	i := 0
	for _, s := range input {
		i += NextHistory2(stringToInts(s))
	}
	return i
}
