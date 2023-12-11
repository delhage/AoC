package day06

import (
	"aoc/2023/utils"
	"math"
	"strconv"
	"strings"
)

type Race struct {
	t int
	d int
}

var input, _ = utils.ReadFile("day06/input.txt")

func getRaces() []Race {
	var rs []Race
	var r Race
	x := strings.Fields(input[0])
	y := strings.Fields(input[1])
	for i := 0; i < len(x)-1; i++ {
		r.t, _ = strconv.Atoi(x[i+1])
		r.d, _ = strconv.Atoi(y[i+1])
		rs = append(rs, r)
	}
	//fmt.Println("rs:", rs)
	return rs
}

func getRace() Race {
	var t string
	var d string
	var r Race
	x := strings.Fields(input[0])
	y := strings.Fields(input[1])
	for i := 0; i < len(x)-1; i++ {
		t = t + x[i+1]
		d = d + y[i+1]
	}
	r.t, _ = strconv.Atoi(t)
	r.d, _ = strconv.Atoi(d)
	return r
}

func getWins(r Race) int {
	var res int
	for i := 0; i <= r.t; i++ {
		if (r.t-i)*i > r.d {
			res += 1
		}
	}
	return res
}

func solveQuad(a, b, c int) (float64, float64) {
	var d = float64(b*b - 4*a*c)
	var x1, x2 float64
	x1 = (float64(-b) + math.Sqrt(d)) / float64(2*a)
	x2 = (float64(-b) - math.Sqrt(d)) / float64(2*a)
	return x1, x2
}

// Solve1 returns answer to first problem
func Solve1() int {
	res := 1
	rs := getRaces()
	for _, r := range rs {
		x, y := solveQuad(-1, r.t, -r.d)
		res *= (int(math.Ceil(y) - (math.Floor(x)) - 1))
		//res *= getWins(r)
	}
	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	r := getRace()
	x, y := solveQuad(-1, r.t, -r.d)
	return int(math.Ceil(y)-(math.Floor(x))) - 1
	//return getWins(r)
}
