package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day16/input.txt")

type rules struct {
	//field                  string
	min1, min2, max1, max2 int
}

func s2i(s []string, d string) [][]int {
	var res [][]int
	for _, s := range s {
		var r []int
		s1 := strings.Split(s, d)
		for _, t := range s1 {
			num, _ := strconv.Atoi(t)
			r = append(r, num)
		}
		res = append(res, r)
	}
	return res
}

func s2rules(s []string) map[string]rules {
	r := make(map[string]rules)
	for _, t := range s {
		main := strings.Split(t, ":")
		//fmt.Println(main)
		field := main[0]
		rest := main[1]
		interv := strings.Split(rest, "or")
		range1 := strings.TrimSpace(interv[0])
		range2 := strings.TrimSpace(interv[1])
		s1 := strings.Split(range1, "-")
		min1, _ := strconv.Atoi(s1[0])
		max1, _ := strconv.Atoi(s1[1])
		s2 := strings.Split(range2, "-")
		min2, _ := strconv.Atoi(s2[0])
		max2, _ := strconv.Atoi(s2[1])
		r[field] = rules{min1: min1, min2: min2, max1: max1, max2: max2}
		//r = append(r, rules{field: field, min1: min1, min2: min2, max1: max1, max2: max2})

	}
	return r
}

// ParseLines parses the file input
func ParseLines(input []string) (map[string]rules, []int, [][]int) {
	//w := make(map[int]int)
	var r []string
	var t []int
	var nt []string
	var ind int
	for i, s := range input {
		if s == "" {
			ind = i + 1
			break
		}
		r = append(r, s)
	}
	t1 := strings.Split(input[ind+1], ",")
	for _, s := range t1 {
		i, _ := strconv.Atoi(s)
		t = append(t, i)
	}

	for i := ind + 4; i < len(input); i++ {
		nt = append(nt, input[i])
	}
	//fmt.Println(r)
	//fmt.Println(s2rules(r))
	//fmt.Println(t)
	//fmt.Println(s2i(nt, ","))
	return s2rules(r), t, s2i(nt, ",")
}

func validateticket(r map[string]rules, ticket []int) int {
	err := 0
	for _, i := range ticket {
		//fmt.Println("I:", i)
		valid := false
		for _, s := range r {
			if (i >= s.min1 && i <= s.max1) || (i >= s.min2 && i <= s.max2) {
				valid = true
				break
			}
		}
		if !valid {
			err += i
		}
	}
	return err
}

func getfield(a []int, r map[string]rules) []string {
	var f []string
	for field, s := range r {
		fmt.Println(field)
		valid := true
		for _, i := range a {
			if !((i >= s.min1 && i <= s.max1) || (i >= s.min2 && i <= s.max2)) {
				valid = false
				break
			}
		}
		if valid {
			f = append(f, field)
		}
	}
	return f
}

// Solve1 returns answer to first problem
func Solve1() int {
	r, _, nt := ParseLines(input)
	fmt.Println(nt)
	err := 0
	for _, p := range nt {
		err += validateticket(r, p)
	}
	return err
}

// Solve2 returns answer to second problem
func Solve2() int {
	r, t, nt := ParseLines(input)
	var valid [][]int
	res := 1
	for _, p := range nt {
		if validateticket(r, p) == 0 {
			valid = append(valid, p)
		}
	}
	//fmt.Println(valid)
	fi := make(map[int][]string)
	for i := 0; i < len(valid[0]); i++ {
		var f []int
		for j := 0; j < len(valid); j++ {
			//fmt.Println(valid[j][i])
			f = append(f, valid[j][i])
		}
		//fmt.Println(f)
		field := getfield(f, r)
		fi[i] = field
		//delete(r, field)
	}
	for i := range valid[0] {
		var f []int
		for j := 0; j < len(valid); j++ {
			f = append(f, valid[j][i])
		}
		field := getfield(f, r)
		fi[i] = field
	}
	fmt.Println(fi)
	return 0
	//fmt.Println(fi)
	rulemap := make(map[int]string)
	itermap := make(map[int]bool)
	//fi := make(map[int][]string)
	for i := range valid[0] {
		itermap[i] = true
	}
	fmt.Println(itermap)
	for len(fi) > 0 {
		fi = make(map[int][]string)
		for i := range itermap {
			//fmt.Println(i, v)
			var f []int
			for j := 0; j < len(valid); j++ {
				//fmt.Println(valid[j][i])
				f = append(f, valid[j][i])
			}
			//fmt.Println(f)
			field := getfield(f, r)
			if len(field) == 0 {
				fmt.Println("Field ", i, " is zero")
			}
			//fmt.Println(field)
			fi[i] = field
			//delete(r, field)
		}
		//fmt.Println(fi)
		for k, s := range fi {
			if len(s) == 1 {
				rulemap[k] = s[0]
				delete(r, s[0])
				delete(itermap, k)
			}
			//fmt.Println(fi)
			//fmt.Println(rulemap)
			//fmt.Println(itermap)
		}
		//return 0
		//fmt.Println(rulemap)
	}
	fmt.Println(rulemap)
	reg := "^departure.*"
	re := regexp.MustCompile(reg)
	for i, s := range rulemap {
		if re.MatchString(s) {
			res *= t[i]
		}
	}
	//fmt.Println(fi)

	//fmt.Println(nt)
	//fmt.Println(valid)
	return res
}
