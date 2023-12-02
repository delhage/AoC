package day02

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/oleiade/reflections"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var input, _ = utils.ReadFile("day02/input.txt")

// ParseLines reads a file into a slice of int:s

type bag struct {
	Red   int
	Blue  int
	Green int
}

type game struct {
	id     int
	sample []bag
}

var bag1 = bag{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func ParseLines(input []string) []game {
	var res []game
	for _, line := range input {
		var g game

		k := strings.Split(line, ":")
		id := strings.Split(k[0], " ")[1]

		var handful []bag

		var tBag = bag{
			Red:   0,
			Green: 0,
			Blue:  0,
		}

		c := strings.Split(k[1], ";")

		for i, s := range c {
			b := strings.Split(s, ",")
			for _, b2 := range b {
				if len(handful) == i {
					handful = append(handful, tBag)
				}
				d := strings.Split(strings.TrimSpace((b2)), " ")

				caser := cases.Title(language.English)

				num, _ := strconv.Atoi(d[0])
				_ = reflections.SetField(&handful[i], caser.String(d[1]), num)
			}
		}
		fmt.Println("Handful: ", handful)
		fmt.Println("id: ", id)
		fmt.Println("")
		g.id, _ = strconv.Atoi(id)
		g.sample = handful
		res = append(res, g)
	}
	return res
}

// Solve1 returns answer to first problem
func Solve1() int {
	res := 0
	games := ParseLines(input)
	fmt.Println(games)
	for _, g := range games {
		for _, h := range g.sample {
			if h.Blue > bag1.Blue || h.Green > bag1.Green || h.Red > bag1.Red {
				g.id = 0
			}
		}
		res += g.id
	}
	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	res := 0
	games := ParseLines(input)

	for _, g := range games {
		blue := 1
		red := 1
		green := 1
		for _, h := range g.sample {
			if h.Blue > blue {
				blue = h.Blue
			}
			if h.Red > red {
				red = h.Red
			}
			if h.Green > green {
				green = h.Green
			}
		}
		power := blue * red * green
		res += power
	}
	return res
}
