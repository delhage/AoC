package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

func (c *point) dist() (int) {
	return abs(c.x) + abs(c.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadLines(filePath string) (f [2][]string) {
	
	fileHandle,_ := os.Open(filePath)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	
	for i := 0; i < 2; i++ {
		fileScanner.Scan()
		f[i] = strings.Split(fileScanner.Text(),",")
	}

	return f
}


func makeLine(a point, b string) ([]point) {
	op := b[0]
	len,_ := strconv.Atoi(b[1:])
	c := make([]point,0)

	switch op {
	case 'R':
		for i := 1; i < len + 1; i++ {
			c = append(c, point{a.x+i,a.y})
		}
	case 'L':
		for i := 1; i < len + 1; i++ {
			c = append(c, point{a.x-i,a.y})
		}
	case 'U':
		for i := 1; i < len + 1; i++ {
			c = append(c, point{a.x,a.y+i})
		}
	case 'D':
		for i := 1; i < len + 1; i++ {
			c = append(c, point{a.x,a.y-i})
		}
	}
	return c
}

func makeCoordinates(f []string) ([]point) {
	c := make([]point,0)
	c = append(c,point{0,0})

	for i := 0; i < len(f); i++ {
		line := makeLine(c[len(c)-1],f[i])
		c = append(c,line...)
		}
	return c
}

func Intersection(a, b []point) (c []point) {
	m := make(map[point]bool)

	for _, item := range a {
		m[item] = true
	}

	delete(m, point{0,0})

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}


func getCross(f[2][]string) (int){
	c0 := makeCoordinates(f[0])
	c1 := makeCoordinates(f[1])

	intersection := Intersection(c0,c1)
	//fmt.Println(intersection)
	distance := 100000
	for i :=0; i< len(intersection);i++ {
		if intersection[i].dist() < distance {
			distance = intersection[i].dist()
		}
	}
	return distance
}

func getCross2(f[2][]string) ([]point){
	c0 := makeCoordinates(f[0])
	c1 := makeCoordinates(f[1])

	return Intersection(c0,c1)
}

func getSteps(a []string, b point) (int){
	c := makeCoordinates(a)
	//fmt.Println(c)
	//fmt.Println(b)
	
	for i := 0; i < len(c); i++ {
		if c[i] == b {
			return i
		}
	}
	return -1
}

func main() {
	lines := ReadLines(os.Args[1])
	
	fmt.Println(getCross(lines))

	c0 := makeCoordinates(lines[0])
	c1 := makeCoordinates(lines[1])

	inter := Intersection(c0,c1)

	totalsteps := 100000
	for i := 0; i < len(inter); i++ {
		steps := getSteps(lines[0],inter[i]) + getSteps(lines[1],inter[i])
		if steps < totalsteps {
			totalsteps = steps
		}
	}

	fmt.Println(totalsteps)

}
