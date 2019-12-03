package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func fuel(mass int) (int) {
	
	if mass <= 0 {
		return 0
	} else {
		return 	mass + fuel(mass/3-2)
	}
}

func ReadFreq(filePath string) (int, error) {
	fileHandle,_ := os.Open(filePath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	
	var mass int = 0
	
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		i, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			return mass, err
		}
		fmt.Println(fuel(i) - i)
		mass += fuel(i) - i
	}
	//fmt.Println(result)
	return mass, fileScanner.Err()
}

func main() {
	freq,err := ReadFreq(os.Args[1])
	fmt.Println(freq, err)
}
