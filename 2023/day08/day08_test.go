package day08

import (
	"aoc/2023/utils"
	"testing"
)

func init() {
	input, _ = utils.ReadFile("testinput1.txt")
}

func TestSolve1(t *testing.T) {
	want := 2
	res := Solve1()
	if res != want {
		t.Fatalf(`Solve1() = %d, want match for %d`, res, want)
	}
}

func TestSolve2(t *testing.T) {
	want := 2
	res := Solve2()
	if res != want {
		t.Fatalf(`Solve2() = %d, want match for %d`, res, want)
	}
}
