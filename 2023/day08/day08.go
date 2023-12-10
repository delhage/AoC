package day08

import (
	"aoc/2023/utils"
	"regexp"
)

var input, _ = utils.ReadFile("day08/input.txt")

type nodePair struct {
	left  string
	right string
}

var mapOfNodes = make(map[string][]nodePair)
var lrinstr string

func initNodes() {
	var nodeRegex = regexp.MustCompile(`([A-Z]+)`)
	for i, s := range input {
		if i == 0 {
			lrinstr = s
			continue
		}
		if i == 1 {
			continue
		}
		k := nodeRegex.FindAllString(s, -1)
		mapOfNodes[k[0]] = append(mapOfNodes[k[0]], nodePair{k[1], k[2]})
	}

}

// Solve1 returns answer to first problem
func Solve1() int {
	initNodes()
	steps := 0
	curNode := "AAA"
	for curNode != "ZZZ" {
		for i := 0; i < len(lrinstr); i++ {
			if string(lrinstr[i]) == "L" {
				curNode = mapOfNodes[curNode][0].left
			} else {
				curNode = mapOfNodes[curNode][0].right
			}
			steps++
			if curNode == "ZZZ" {
				break
			}
		}
	}

	return steps
}

// Solve2 returns answer to second problem
func Solve2() int {
	return 1
}
