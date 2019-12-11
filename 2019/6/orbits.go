package main

import (
	"strings"
)

type object struct {
	name string
	parent *object
}

func ReadLines(filePath string) (f [2][]string) {
	
	fileHandle,_ := os.Open(filePath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	orbits make([]object,0)
	
	for fileScanner.Scan() {
		f, err := fileScanner.Text()
		a,b := strings.Split(f,")")
		
	}

	return f
}
