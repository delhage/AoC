package main

import (
	"os"

	"aoc/2023/day01"
	"aoc/2023/utils"
)

func main() {

	switch os.Args[1] {
	case "day01":
		utils.Perf("2023-12-01", day01.Solve1, day01.Solve2)
	}
}
