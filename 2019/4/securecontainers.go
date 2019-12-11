package main

import (
	"fmt"
	//	"strings"
	"strconv"
)

func twoSame(a string) (bool) {
	for i := 1; i < len(a); i++ {
		k,_ := strconv.Atoi(a[i-1:i])
		l,_ := strconv.Atoi(a[i:i+1])
		
		if k == l {
			return true
		}
	}
	return false
}

func twoSame2(a string) (bool) {

	fmt.Println(a)
	var ia [6]int
	for i := 0; i < len(ia); i++ {
		ia[i],_ = strconv.Atoi(a[i:i+1])
	}
	
	if ia[0] == ia[1] && ia[1] != ia[2] {
		return true
	}

	if ia[5] == ia[4] && ia[4] != ia[3] {
		return true
	}
	
	for i := 1; i < len(ia) - 2; i++ {
		k := ia[i-1]
		l := ia[i]
		m := ia[i+1]
		n := ia[i+2]
		
		if l == m {
			if k !=l && n !=m {
				return true
			}
		}
	}
	return false
}

func increasing(a string) (bool) {
	for i := 1; i < len(a); i++ {
		k,_ := strconv.Atoi(a[i-1:i])
		l,_ := strconv.Atoi(a[i:i+1])
		if l < k {
			return false
		}
	}
	return true
	
}

func passOk(a string) (bool) {
	return twoSame(a) && increasing(a)
}

func passOk2(a string) (bool) {
	return twoSame2(a) && increasing(a)
}

func main() {
	num := 0
	num2 := 0
	for i := 136818; i < 685980; i++ {
		if passOk(strconv.Itoa(i)) {
			num++
		}
		if passOk2(strconv.Itoa(i)) {
			num2++
		}
	}
	fmt.Println(num, num2)
	// fmt.Println(passOk2("112233"))
	// fmt.Println(passOk2("123444"))
	// fmt.Println(passOk2("111122"))
}
