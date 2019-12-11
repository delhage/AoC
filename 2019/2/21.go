package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}

func ReadArray(filePath string) ([]int, error) {
	fileHandle,_ := os.Open(filePath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Scan()
	//fmt.Println(fileScanner.Text())
	return sliceAtoi(strings.Split(fileScanner.Text(),","))
}

func calc(noun int, verb int) (int) {
	a,_ := ReadArray(os.Args[1])
	// fmt.Println(a)
	a[1] = noun
	a[2] = verb

	d:for i := 0; i < len(a); i = i+4 {
		//fmt.Println(a[i])
		switch a[i] {
		case 1:
			a[a[i+3]] = a[a[i+1]] + a[a[i+2]]
		case 2:
			a[a[i+3]] = a[a[i+1]] * a[a[i+2]]
		case 99:
			break d
		}
	}

	//fmt.Println(a[0])
	return a[0]
}

func main() {

	e:for i := 0; i < 100; i++ {
		for j :=0; j < 100; j++ {
			if calc(i,j) == 19690720 {
				fmt.Println(100*i+j)
				break e
			}
		}
	}
}
