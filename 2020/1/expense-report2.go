package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readfile(path string) ([]int, error) {
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
	x, _ := readfile(os.Args[1])
	for i, s := range x {
		for j := i + 1; j < len(x); j++ {
			for k := j + i; k < len(x); k++ {
				res = s + x[j] + x[k]
				// fmt.Println(res)
				if res == 2020 {
					fmt.Println(s * x[j] * x[k])
					os.Exit(0)
				}
			}
		}
	}
}
