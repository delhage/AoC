package main

import (
	"testing"
	"strings"
)

func TestDist(t *testing.T) {
	var f [2][]string
	tables := []struct {
		a string
		b string
		n int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72","U62,R66,U55,R34,D71,R55,D58,R83",159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51","U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",135},
	}

	for _, table := range tables {
		f[0] = strings.Split(table.a,",")
		f[1] = strings.Split(table.b,",")
		dist := getCross(f)

		if dist != table.n {
			t.Errorf("Distance incorrect, got: %d, want: %d.", dist, table.n)
		}
	}
}

func TestLine(t *testing.T) {
	tables := []struct {
		a point
		b string
		c []point
	}{
		{point{0,0},"RL",[]point{{1,0},{2,0},{3,0}}},
		{point{0,0},"L3",[]point{{-1,0},{-2,0},{-3,0}}},
		{point{0,0},"U3",[]point{{0,1},{0,2},{0,3}}},
		{point{0,0},"D3",[]point{{0,-1},{0,-2},{0,-3}}},
	}

	for _, table := range tables {
		line := makeLine(table.a,table.b)

		for i := 0; i < len(line); i++ {
			if line[i] != table.c[i] {
			t.Errorf("Line incorrect, got: %d, want: %d.", line[i], table.c[i])
			}
		}
	}
}

func TestSteps(t *testing.T) {
	tables := []struct {
		a []string
		b point
		n int
	}{
		{[]string{"R8","U5","L5","D3"},point{3,3},20},
		{[]string{"R8","U5","L5","D3"},point{6,5},15},
	}

	for _, table := range tables {
		steps := getSteps(table.a,table.b)

		if steps != table.n {
			t.Errorf("Steps incorrect, got: %d, want: %d.", steps, table.n)
		}
	}
}
