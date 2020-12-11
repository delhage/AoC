package utils

import (
	"log"
	"time"
)

// Perf executes the problems and prints running time
func Perf(date string, solve1 func() int, solve2 func() int) {
	// Store time to print out execution time
	start := time.Now()

	log.Printf("%s: Starting program to solve %s\n", date, date)

	//parse(os.Args[2])

	//log.Printf("%s: Spent %s on parsing input\n", date, time.Since(start))
	//start = time.Now()

	s1 := solve1()

	log.Printf("Answer part 1: %d, time spent: %s\n", s1, time.Since(start))
	start = time.Now()

	s2 := solve2()

	log.Printf("Answer part 2: %d, time spent: %s\n", s2, time.Since(start))
}
