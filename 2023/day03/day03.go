package day03

import (
	"aoc/2023/utils"
	"fmt"
	"regexp"
	"strconv"
)

var input, _ = utils.ReadFile("day03/testinput1.txt")
var symbols [10][10]bool
var numbers [10][10]bool

func isAdjacent(row int, col []int) bool {
	// Check to the left
	if col[0] > 0 {
		if symbols[row][col[0]-1] {
			return true
		}
		if row > 0 {
			if symbols[row-1][col[0]-1] {
				return true
			}
		}
		if row < len(symbols)-1 {
			if symbols[row+1][col[0]-1] {
				return true
			}
		}
	}
	// Check to the right
	//fmt.Println(len(symbols), col[1])
	if col[1] < len(symbols) {
		if symbols[row][col[1]] {
			return true
		}
		if row > 0 {
			if symbols[row-1][col[1]] {
				return true
			}
		}
		if row < len(symbols)-1 {
			if symbols[row+1][col[1]] {
				return true
			}
		}
	}
	for i := col[0]; i < col[1]; i++ {
		// Check row above
		if row > 0 {
			if symbols[row-1][i] {
				return true
			}
		}
		// Check row below
		if row < len(symbols)-1 {
			if symbols[row+1][i] {
				return true
			}
		}
	}
	return false
}

// Solve1 returns answer to first problem
func Solve1() int {
	var numbersRegex = regexp.MustCompile(`([0-9]+)`)
	var symbolsRegex = regexp.MustCompile(`[^0-9\.]`)
	//var number string
	//var res int
	var sum int
	var total int

	// Set symbols matrix
	for i, s := range input {
		othersindex := symbolsRegex.FindAllStringIndex(s, -1)
		for _, c := range othersindex {
			symbols[i][c[0]] = true
		}
	}

	// Get valid numbers
	for i, s := range input {
		res := numbersRegex.FindAllString(s, -1)
		if i == 0 {
			fmt.Println(res)
		}
		resindex := numbersRegex.FindAllStringIndex(s, -1)
		for _, h := range res {
			l, _ := strconv.Atoi(h)
			total += l
		}
		for k, j := range resindex {
			//fmt.Println("J, K, res: ", j, k, res[k])
			if isAdjacent(i, j) {
				l, _ := strconv.Atoi(res[k])
				sum += l
				//fmt.Println("SUM, L: ", sum, l)
			}
		}
	}
	fmt.Println("Total: ", total)
	return sum
}

// Solve2 returns answer to second problem
func Solve2() int {
	var numbersRegex = regexp.MustCompile(`([0-9]+)`)
	var gearsRegex = regexp.MustCompile(`*`)

	// Set numbers matrix
	for i, s := range input {
		numbersIndex := numbersRegex.FindAllStringIndex(s, -1)
		for _, c := range numbersIndex {
			fmt.Println("C: ", c)
			for k := c[0]; k < c[1]; k++ {
				numbers[i][k] = true
			}
		}
	}

	// Get valid numbers
	for i, s := range input {
		res := numbersRegex.FindAllString(s, -1)
		if i == 0 {
			fmt.Println(res)
		}
		resindex := numbersRegex.FindAllStringIndex(s, -1)
		for _, h := range res {
			l, _ := strconv.Atoi(h)
			total += l
		}
		for k, j := range resindex {
			//fmt.Println("J, K, res: ", j, k, res[k])
			if isAdjacent(i, j) {
				l, _ := strconv.Atoi(res[k])
				sum += l
				//fmt.Println("SUM, L: ", sum, l)
			}
		}
	}
	fmt.Println(numbers)
	return 1
}
