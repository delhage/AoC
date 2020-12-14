package main

import (
	"os"

	"./day01"
	"./day02"
	"./day03"
	"./day04"
	"./day05"
	"./day07"
	"./day09"
	"./day10"
	"./day11"
	"./day12"
	"./day13"
	"./utils"
)

func main() {

	switch os.Args[1] {
	case "day01":
		utils.Perf("2020-12-01", day01.Solve1, day01.Solve2)
	case "day02":
		utils.Perf("2020-12-02", day02.Solve1, day02.Solve2)
	case "day03":
		utils.Perf("2020-12-03", day03.Solve1, day03.Solve2)
	case "day04":
		utils.Perf("2020-12-04", day04.Solve1, day04.Solve2)
	case "day05":
		utils.Perf("2020-12-05", day05.Solve1, day05.Solve2)
	case "day07":
		utils.Perf("2020-12-07", day07.Solve1, day07.Solve2)
	case "day09":
		utils.Perf("2020-12-09", day09.Solve1, day09.Solve2)
	case "day10":
		utils.Perf("2020-12-10", day10.Solve1, day10.Solve2)
	case "day11":
		utils.Perf("2020-12-11", day11.Solve1, day11.Solve2)
	case "day12":
		utils.Perf("2020-12-12", day12.Solve1, day12.Solve2)
	case "day13":
		utils.Perf("2020-12-13", day13.Solve1, day13.Solve2)
	}
}
