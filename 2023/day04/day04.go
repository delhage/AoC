package day04

import (
	"aoc/2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input, _ = utils.ReadFile("day04/input.txt")

type card struct {
	id      int
	winning []int
	play    []int
	result  int
}

func initCards() []card {
	var res []card
	var numbersRegex = regexp.MustCompile(`([0-9]+)`)

	for i, s := range input {
		k := strings.Split(s, ":")
		l := strings.Split(k[1], "|")
		m := numbersRegex.FindAllString(l[0], -1)
		n := numbersRegex.FindAllString(l[1], -1)
		var c card

		c.winning = make([]int, len(m))
		for i, s := range m {
			c.winning[i], _ = strconv.Atoi(s)
		}
		c.play = make([]int, len(n))
		for i, s := range n {
			c.play[i], _ = strconv.Atoi(s)
		}
		c.id = i + 1
		c.result = winningNumbers(c)
		fmt.Println(c.result)
		res = append(res, c)
	}
	return res
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func winningNumbers(c card) int {
	var res = 0
	for _, n := range c.play {
		if Contains(c.winning, n) {
			if res == 0 {
				res++
			} else {
				res = 2 * res
			}
		}
	}
	return res
}

// Solve1 returns answer to first problem
func Solve1() int {
	var res = initCards()
	var total int
	for _, c := range res {
		total += c.result
	}
	return total
}

// Solve2 returns answer to second problem
func Solve2() int {
	return 1
}
