package day12

import (
	"strconv"

	"../utils"
)

var input, _ = utils.ReadFile("day12/input.txt")

// Ship describes the ship's position and direction
type Ship struct {
	x int
	y int
	a int
}

// Instr is moving instructions
type Instr struct {
	i byte
	n int
}

// ParseLines parses the file input
func ParseLines(input []string) []Instr {
	var in []Instr
	for _, line := range input {
		i := []byte(line[0:1])[0]
		n, _ := strconv.Atoi(line[1:])
		in = append(in, Instr{i: i, n: n})
	}
	return in
}

func travel(instr []Instr) int {
	ship := Ship{x: 0, y: 0, a: 0}
	for _, ins := range instr {
		switch ins.i {
		case 'N':
			ship.y -= ins.n
		case 'S':
			ship.y += ins.n
		case 'E':
			ship.x += ins.n
		case 'W':
			ship.x -= ins.n
		case 'R':
			ship.a = (ship.a + ins.n) % 360
		case 'L':
			ship.a = (ship.a - ins.n + 360) % 360
		case 'F':
			//fmt.Println(ship.a, ins.n)
			switch ship.a {
			case 0:
				ship.x += ins.n
			case 90:
				ship.y += ins.n
			case 180:
				ship.x -= ins.n
			case 270:
				ship.y -= ins.n
			}
		}
	}
	if ship.x < 0 {
		ship.x = -ship.x
	}
	if ship.y < 0 {
		ship.y = -ship.y
	}
	return ship.x + ship.y
}

func travel2(instr []Instr) int {
	ship := Ship{x: 0, y: 0, a: 0}
	waypoint := Ship{x: 10, y: -1, a: 0}
	for _, ins := range instr {
		switch ins.i {
		case 'N':
			waypoint.y -= ins.n
		case 'S':
			waypoint.y += ins.n
		case 'E':
			waypoint.x += ins.n
		case 'W':
			waypoint.x -= ins.n
		case 'R':
			switch ins.n {
			case 90:
				x := -waypoint.y
				y := waypoint.x
				waypoint.x = x
				waypoint.y = y
			case 180:
				waypoint.x = -waypoint.x
				waypoint.y = -waypoint.y
			case 270:
				x := waypoint.y
				y := -waypoint.x
				waypoint.x = x
				waypoint.y = y
			}
		case 'L':
			switch ins.n {
			case 90:
				x := waypoint.y
				y := -waypoint.x
				waypoint.x = x
				waypoint.y = y
			case 180:
				waypoint.x = -waypoint.x
				waypoint.y = -waypoint.y
			case 270:
				x := -waypoint.y
				y := waypoint.x
				waypoint.x = x
				waypoint.y = y
			}

		case 'F':
			ship.x += ins.n * waypoint.x
			ship.y += ins.n * waypoint.y
		}
	}
	if ship.x < 0 {
		ship.x = -ship.x
	}
	if ship.y < 0 {
		ship.y = -ship.y
	}
	return ship.x + ship.y
}

// Solve1 returns answer to first problem
func Solve1() int {
	in := ParseLines(input)
	return travel(in)
}

// Solve2 returns answer to second problem
func Solve2() int {
	in := ParseLines(input)
	return travel2(in)
}
