package day07

import (
	"strconv"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day07/input.txt")

// Bag is a structure of bags
type Bag struct {
	id      string
	content map[string]int
}

var bags []Bag

// ParseLines parses file input
func ParseLines(input []string) {
	for _, line := range input {
		parts := strings.Split(line, "contain")
		main := strings.Split(parts[0], " ")
		descr := main[0]
		color := main[1]
		rest := strings.Split(parts[1], ",")
		m := make(map[string]int)
		bag := Bag{id: descr + " " + color, content: m}
		for _, b := range rest {
			s := strings.Split(strings.TrimSpace(b), " ")
			num, _ := strconv.Atoi(s[0])
			id := s[1] + " " + s[2]
			bag.content[id] = num
		}
		bags = append(bags, bag)
	}
}

func contains(id string) bool {
	for _, s := range bags {
		if s.id == id {
			if s.content["shiny gold"] > 0 {
				return true
			}

			if len(s.content) > 0 {
				can := false
				for t := range s.content {
					if contains(t) {
						can = true
					}
				}
				return can
			}
		}
	}
	return false
}

func contains2(id string) int {
	var num int
	for _, s := range bags {
		if s.id == id {
			if len(s.content) == 0 {
				return 0
			}

			for c, sum := range s.content {
				num += sum + sum*contains2(c)
			}
			return num
		}
	}
	return 0
}

// Solve1 returns answer to first problem
func Solve1() int {
	var g int
	ParseLines(input)
	for _, s := range bags {
		if contains(s.id) {
			g++
		}
	}
	return g
}

// Solve2 returns answer to secon problem
func Solve2() int {
	ParseLines(input)
	return contains2("shiny gold")
}
