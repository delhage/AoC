package day15

import (
	"aoc/2023/utils"
	"strings"
)

var input, _ = utils.ReadFile("day15/testinput.txt")

type lens struct {
	name   string
	length int
}

func getHash(s string) int {
	h := 0
	for _, c := range s {
		h += int(c)
		h *= 17
		h = h % 256
	}
	return h
}

// func getBox(s string) (string, int, int, bool) {
// 	var c string
// 	var length int
// 	var remove bool
// 	if s[len(s)-1] == '-' {
// 		c = s[0 : len(s)-1]
// 		remove = true
// 		length = 0
// 	} else {
// 		c = s[0 : len(s)-2]
// 		remove = false
// 		length, _ = strconv.Atoi(string(c[len(s)-1]))
// 	}
// 	fmt.Println("c:", c)
// 	return c, getHash(c), length, remove
// }

// Solve1 returns answer to first problem
func Solve1() int {
	i := 0
	s := strings.Split(input[0], ",")
	for _, c := range s {
		i += getHash(c)
	}
	return i
}

// Solve2 returns answer to second problem
func Solve2() int {
	// box := make([]lens, 256)
	// s := strings.Split(input[0], ",")
	// for _, c := range s {
	// 	//fmt.Println("c:", getBox(c))
	// 	_, i, l, remove := getBox(c)
	// 	if !remove {
	// 		box[i].name = c
	// 		box[i].length = l
	// 	} else {

	// 	}
	// }
	return 1
}
