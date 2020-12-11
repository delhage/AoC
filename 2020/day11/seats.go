package day11

import (
	"fmt"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day11/input.txt")

// ParseLines parses the file input
func ParseLines(input []string) [][]byte {
	var seatmap [][]byte
	for _, line := range input {
		seatmap = append(seatmap, []byte(strings.TrimSpace(line)))
	}
	return seatmap
}

func changeEmpty(i int, j int, seatmap [][]byte) bool {
	left := false
	right := false
	top := false
	bottom := false
	if i == 0 {
		top = true
	}
	if i == len(seatmap)-1 {
		bottom = true
	}
	if j == 0 {
		left = true
	}
	if j == len(seatmap[0])-1 {
		right = true
	}

	if !top && !left && seatmap[i-1][j-1] == '#' {
		return false
	}
	if !left && seatmap[i][j-1] == '#' {
		return false
	}
	if !bottom && !left && seatmap[i+1][j-1] == '#' {
		return false
	}
	if !top && seatmap[i-1][j] == '#' {
		return false
	}
	if !bottom && seatmap[i+1][j] == '#' {
		return false
	}
	if !top && !right && seatmap[i-1][j+1] == '#' {
		//fmt.Println("Indices", i, j)
		return false
	}
	if !right && seatmap[i][j+1] == '#' {
		return false
	}
	if !bottom && !right && seatmap[i+1][j+1] == '#' {
		return false
	}
	return true
}

func changeOccupied(i int, j int, seatmap [][]byte) bool {
	occ := 0
	left := false
	right := false
	top := false
	bottom := false
	if i == 0 {
		top = true
	}
	if i == len(seatmap)-1 {
		bottom = true
	}
	if j == 0 {
		left = true
	}
	if j == len(seatmap[0])-1 {
		right = true
	}

	if !top && !left && seatmap[i-1][j-1] == '#' {
		occ++
	}
	if !left && seatmap[i][j-1] == '#' {
		occ++
	}
	if !bottom && !left && seatmap[i+1][j-1] == '#' {
		occ++
	}
	if !top && seatmap[i-1][j] == '#' {
		occ++
	}
	if !bottom && seatmap[i+1][j] == '#' {
		occ++
	}
	if !top && !right && seatmap[i-1][j+1] == '#' {
		occ++
	}
	if !right && seatmap[i][j+1] == '#' {
		occ++
	}
	if !bottom && !right && seatmap[i+1][j+1] == '#' {
		occ++
	}
	return occ > 3
}

func nextMap(seatmap [][]byte) ([][]byte, bool) {
	changed := false
	nextmap := make([][]byte, len(seatmap))
	for i := range seatmap {
		nextmap[i] = make([]byte, len(seatmap[i]))
		copy(nextmap[i], seatmap[i])
	}
	for i, row := range seatmap {
		for j, seat := range row {
			switch seat {
			case '.':
				continue
			case 'L':
				if changeEmpty(i, j, seatmap) {
					nextmap[i][j] = '#'
					changed = true
				}
			case '#':
				if changeOccupied(i, j, seatmap) {
					nextmap[i][j] = 'L'
					changed = true
				}
			}
		}
	}
	return nextmap, changed
}

func changeEmpty2(i int, j int, seatmap [][]byte) bool {
	maxcol := len(seatmap[0])
	maxrow := len(seatmap)
	// Look left
	for l := j - 1; l >= 0; l-- {
		s := seatmap[i][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
	}
	// Look right
	for l := j + 1; l < maxcol; l++ {
		s := seatmap[i][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
	}
	// Look up
	for k := i - 1; k >= 0; k-- {
		s := seatmap[k][j]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
	}
	// Look down
	for k := i + 1; k < maxrow; k++ {
		s := seatmap[k][j]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
	}
	// Look up left
	for k, l := i-1, j-1; k >= 0 && l >= 0; k-- {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
		l--
	}
	// Look down left
	for k, l := i+1, j-1; k < maxrow && l >= 0; k++ {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
		l--
	}
	// Look up right
	for k, l := i-1, j+1; k >= 0 && l < maxcol; k-- {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
		l++
	}
	// Look down right
	for k, l := i+1, j+1; k < maxrow && l < maxcol; k++ {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			return false
		}
		l++
	}
	return true
}

func changeOccupied2(i int, j int, seatmap [][]byte) bool {
	maxcol := len(seatmap[0])
	maxrow := len(seatmap)
	occ := 0
	// Look left
	for l := j - 1; l >= 0; l-- {
		s := seatmap[i][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
	}
	// Look right
	for l := j + 1; l < maxcol; l++ {
		s := seatmap[i][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
	}
	// Look up
	for k := i - 1; k >= 0; k-- {
		s := seatmap[k][j]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
	}
	// Look down
	for k := i + 1; k < maxrow; k++ {
		s := seatmap[k][j]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
	}
	// Look up left
	for k, l := i-1, j-1; k >= 0 && l >= 0; k-- {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
		l--
	}
	// Look down left
	for k, l := i+1, j-1; k < maxrow && l >= 0; k++ {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
		l--
	}
	// Look up right
	for k, l := i-1, j+1; k >= 0 && l < maxcol; k-- {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
		l++
	}
	// Look down right
	for k, l := i+1, j+1; k < maxrow && l < maxcol; k++ {
		s := seatmap[k][l]
		if s == 'L' {
			break
		}
		if s == '#' {
			occ++
			break
		}
		l++
	}

	return occ > 4
}

func nextMap2(seatmap [][]byte) ([][]byte, bool) {
	changed := false
	nextmap := make([][]byte, len(seatmap))
	for i := range seatmap {
		nextmap[i] = make([]byte, len(seatmap[i]))
		copy(nextmap[i], seatmap[i])
	}
	for i, row := range seatmap {
		for j, seat := range row {
			switch seat {
			case 'L':
				if changeEmpty2(i, j, seatmap) {
					nextmap[i][j] = '#'
					changed = true
				}
			case '#':
				if changeOccupied2(i, j, seatmap) {
					nextmap[i][j] = 'L'
					changed = true
				}
			}
		}
	}
	return nextmap, changed
}

// Solve1 returns answer to first problem
func Solve1() int {
	occ := 0
	seatmap := ParseLines(input)
	//nextmap, changed := nextMap(seatmap)
	changed := true

	for changed {
		nextmap, ch := nextMap(seatmap)
		for i := range nextmap {
			copy(seatmap[i], nextmap[i])
		}
		changed = ch
	}
	for _, row := range seatmap {
		for _, seat := range row {
			if seat == '#' {
				occ++
			}
		}
	}
	return occ
}

// Solve2 returns answer to second problem
func Solve2() int {
	occ := 0
	seatmap := ParseLines(input)
	changed := true
	for changed {
		nextmap, ch := nextMap2(seatmap)

		for i := range nextmap {
			copy(seatmap[i], nextmap[i])
		}
		//printmap(nextmap)
		changed = ch
	}
	for _, row := range seatmap {
		//fmt.Println(string(row))
		for _, seat := range row {
			if seat == '#' {
				occ++
			}
		}
	}
	return occ
}

func printmap(pmap [][]byte) {
	for _, row := range pmap {
		fmt.Println(string(row))
	}
	fmt.Println("")
}
