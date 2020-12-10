package main

import (
	"os"

	"./day09"
	"./day10"
	"./utils"
)

func main() {
	//fmt.Printf("%d number of bag colors contain at least one 'shiny gold'\n", day07.Solve1())
	//fmt.Printf("%d bags are required inside one 'shiny gold' bag\n", day07.Solve2())
	//acc,loop := day08.Solve1()
	//fmt.Printf("Accumulator is %d. Loop found: %t\n", acc, loop)
	//acc = day08.Solve2()
	//fmt.Printf("Accumulator is %d", acc)
	//fmt.Println(day09.Solve1())
	//fmt.Println(day09.Solve2())
	//fmt.Println(day10.Solve1())
	switch os.Args[1] {
	case "day09":
		utils.Perf("2020-12-09", day09.Solve1, day09.Solve2)
	case "day10":
		utils.Perf("2020-12-10", day10.Solve1, day10.Solve2)
	}
	//start := time.Now()
	//fmt.Println(day10.Solve2())
	//fmt.Printf("Time spent: %s\n", time.Since(start))
}
