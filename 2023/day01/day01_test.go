package day01

import (
	"aoc/2023/utils"
	"testing"
)

func TestSolve1(t *testing.T) {
	input, _ = utils.ReadFile("testinput.txt")
	want := 142
	res := Solve1()
	if res != want {
		t.Fatalf(`Solve1() = %d, want match for %d`, res, want)
	}
}

func TestSolve2(t *testing.T) {
	input, _ = utils.ReadFile("testinput2.txt")
	want := 281
	res := Solve2()
	if res != want {
		t.Fatalf(`Solve2() = %d, want match for %d`, res, want)
	}
}
