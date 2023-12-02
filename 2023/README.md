# Advent of Code 2023

This directory contains one subdirectory for each day plus a `utils` directory. The `main.go` in 
the top (this) direcory runs functions `Solve1()` and `Solve2()` in the respective day packages. It does this
using `utils/perf.go` which will also measure and print the time it takes for each function.

Run the program for one day (day01 in this example):

```bash
$ go run main day01
```
