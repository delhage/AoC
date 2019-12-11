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

func main() {
	fmt.Print("Enter input value: ")
	var text string
	fmt.Scanln(&text)
	inp, err := strconv.Atoi(text)

	fmt.Println(err)
	
	fmt.Printf("%T, %d, %T, %s\n",inp,inp,text,text)
	a,_ := ReadArray(os.Args[1])
	//fmt.Println(a)

	var op,pc,c,d,val1,val2 int
	
	d:for pc < len(a) {
		//fmt.Println(a[i])
		op = a[pc] % 100
		c = a[pc] / 100 % 10
		d = a[pc] / 1000 % 10
		//e = a[pc] / 10000 % 10

		switch op {

		case 1:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}


			a[a[pc+3]] = val1 + val2
			pc += 4
		case 2:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}


			a[a[pc+3]] = val1 * val2
			pc += 4
		case 3:
			a[a[pc+1]] = inp
			pc +=2
		case 4:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			fmt.Println(val1)
			pc +=2
		case 5:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}

			if val1 != 0 {
				pc = val2
			} else {
				pc +=3
			}
		case 6:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}

			if val1 == 0 {
				pc = val2
			} else {
				pc +=3
			}
		case 7:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}
			if val1 < val2 {
				a[a[pc+3]] = 1
			} else {
				a[a[pc+3]] = 0
			}
			pc +=4
		case 8:
			switch c {
			case 0:
				val1 = a[a[pc+1]]
			case 1:
				val1 = a[pc+1]
			}
			switch d {
			case 0:
				val2 = a[a[pc+2]]
			case 1:
				val2 = a[pc+2]
			}

			if val1 == val2 {
				a[a[pc+3]] = 1
			} else {
				a[a[pc+3]] = 0
			}
			pc +=4
		case 99:
			break d
		}
	}

	//fmt.Println(a)
}
