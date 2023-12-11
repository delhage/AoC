package main

import (
	"os"

	"aoc/2023/day01"
	"aoc/2023/day02"
	"aoc/2023/day03"
	"aoc/2023/day04"
	"aoc/2023/day06"
	"aoc/2023/day07"
	"aoc/2023/day08"
	"aoc/2023/day09"
	"aoc/2023/utils"
)

func main() {

	switch os.Args[1] {
	case "day01":
		utils.Perf("2023-12-01", day01.Solve1, day01.Solve2)
	case "day02":
		utils.Perf("2023-12-02", day02.Solve1, day02.Solve2)
	case "day03":
		utils.Perf("2023-12-03", day03.Solve1, day03.Solve2)
	case "day04":
		utils.Perf("2023-12-04", day04.Solve1, day04.Solve2)
	case "day06":
		utils.Perf("2023-12-06", day06.Solve1, day06.Solve2)
	case "day07":
		utils.Perf("2023-12-07", day07.Solve1, day07.Solve2)
	case "day08":
		utils.Perf("2023-12-08", day08.Solve1, day08.Solve2)
	case "day09":
		utils.Perf("2023-12-09", day09.Solve1, day09.Solve2)
	}
}
