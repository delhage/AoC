package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadExpense reads a file into a slice of int:s
func ReadExpense(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func main() {
	var res int
	x, _ := ReadExpense(os.Args[1])
	for i, s := range x {
		for j := i + 1; j < len(x); j++ {
			res = s + x[j]
			// fmt.Println(res)
			if res == 2020 {
				fmt.Println(s * x[j])
				os.Exit(0)
			}
		}
	}
}
