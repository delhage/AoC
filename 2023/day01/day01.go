package day01

import (
	"aoc/2023/utils"
	"regexp"
	"strconv"
)

var input, _ = utils.ReadFile("day01/input.txt")

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func convertToNumber(s string) string {
	if i, err := strconv.Atoi(s); err == nil {
		return strconv.Itoa(i)
	} else {
		return strconv.Itoa(numberMap[s])
	}
}
func findFirst(sub *regexp.Regexp, line string) string {
	var res = sub.FindAllString(line, -1)
	if res != nil {
		return res[0]
	} else {
		return ""
	}
}

func findLast(sub *regexp.Regexp, line string) string {
	var res = sub.FindAllStringSubmatch(line, -1)
	if res != nil {
		var tmp = res[len(res)-1]
		return tmp[len(tmp)-1]
	} else {
		return ""
	}
}

// Solve1 returns answer to first problem
func Solve1() int {
	var numbersRegex = regexp.MustCompile(`[1-9]`)
	var number string
	var res int

	for _, s := range input {
		var first = findFirst(numbersRegex, s)
		var last = findLast(numbersRegex, s)

		number = first + last
		var i, _ = strconv.Atoi(number)
		res += i
	}

	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	var numbersRegex1 = regexp.MustCompile(`(?:.*)([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
	var numbersRegex2 = regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)
	var number string
	var res int
	for _, s := range input {
		var first = findFirst(numbersRegex2, s)
		var last = findLast(numbersRegex1, s)
		number = convertToNumber(first) + convertToNumber(last)
		var i, _ = strconv.Atoi(number)
		res += i
	}

	return res
}
