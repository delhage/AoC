package day03

import (
	"aoc/2020/utils"
)

var input, _ = utils.ReadFile("day03/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) [][]byte {
	var roadmap [][]byte
	for _, line := range input {
		roadmap = append(roadmap, []byte(line))
	}
	return roadmap
}

func numtrees(dx, dy int, roadmap [][]byte) int {
	x := dx
	y := dy
	l := len(roadmap[0])
	h := len(roadmap)

	var trees int
	for {
		if y >= h {
			break
		}
		if x >= l {
			x = x - l
		}
		if roadmap[y][x] == []byte("#")[0] {
			trees++
		}
		x += dx
		y += dy
	}
	return trees
}

// Solve1 returns answer to first problem
func Solve1() int {
	roadmap := ParseLines(input)
	return numtrees(3, 1, roadmap)
}

// Solve2 returns answer to second problem
func Solve2() int {
	roadmap := ParseLines(input)
	res1 := numtrees(1, 1, roadmap)
	res2 := numtrees(3, 1, roadmap)
	res3 := numtrees(5, 1, roadmap)
	res4 := numtrees(7, 1, roadmap)
	res5 := numtrees(1, 2, roadmap)
	return res1 * res2 * res3 * res4 * res5
}
