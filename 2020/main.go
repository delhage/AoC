package main

import (
	"os"

	"./day07"
	"./day09"
	"./day10"
	"./day11"
	"./day12"
	"./utils"
)

func main() {

	switch os.Args[1] {
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

	}
}
