package day04

import (
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

var input, _ = utils.ReadFile("day04/input.txt")

type passport map[string]string

// ParseLines parses the file input
func ParseLines(input []string) []passport {
	passlist := make([]passport, 0, 0)
	pmap := make(passport)
	var pass string
	for _, line := range input {
		if line != "" {
			pass = pass + " " + line
			pass = strings.TrimSpace(pass)
			for _, p := range strings.Split(pass, " ") {
				k := strings.Split(p, ":")
				pmap[k[0]] = k[1]
			}
		} else {
			passlist = append(passlist, pmap)
			pass = ""
			pmap = make(passport)
		}
	}
	passlist = append(passlist, pmap)
	return passlist
}

func validateyear(p passport, k string, reg string, min, max int) bool {
	re := regexp.MustCompile(reg)
	if re.MatchString(p[k]) {
		y, _ := strconv.Atoi(p[k])
		if y < min || y > max {
			return false
		}

	} else {
		return false
	}
	return true
}

func validateheight(p passport) bool {
	h := p["hgt"]
	reg := "^[0-9]{1,3}[incm]{2}$"
	re := regexp.MustCompile(reg)
	if !re.MatchString(h) {
		return false
	}
	unit := h[len(h)-2 : len(h)]
	height, _ := strconv.Atoi(h[0 : len(h)-2])
	//fmt.Println(h, unit, height)
	if unit == "cm" {
		if height < 150 || height > 193 {
			return false
		}
	}
	if unit == "in" {
		if height < 59 || height > 76 {
			return false
		}
	}
	return true
}

func validateeyecolor(p passport) bool {
	var colors = []string{
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}
	for _, c := range colors {
		if c == p["ecl"] {
			return true
		}
	}
	return false

}

func validatepass1(p passport) int {
	var keys = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, k := range keys {
		if _, ok := p[k]; !ok {
			return 0
		}
	}
	return 1
}

func validatepass2(p passport) int {
	var keys = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, k := range keys {
		if _, ok := p[k]; !ok {
			return 0
		}
	}
	reg := "^[0-9]{4}$"
	if !validateyear(p, "byr", reg, 1920, 2002) {
		return 0
	}
	if !validateyear(p, "iyr", reg, 2010, 2020) {
		return 0
	}
	if !validateyear(p, "eyr", reg, 2020, 2030) {
		return 0
	}
	reg = "^[0-9]{9}$"
	re := regexp.MustCompile(reg)
	if !re.MatchString(p["pid"]) {
		return 0
	}
	reg = "^#[0-9a-f]{6}"
	re = regexp.MustCompile(reg)
	if !re.MatchString(p["hcl"]) {
		return 0
	}
	if !validateeyecolor(p) {
		return 0
	}
	if !validateheight(p) {
		return 0
	}
	return 1
}

// Solve1 returns answer to first problem
func Solve1() int {
	var i int
	passlist := ParseLines(input)
	for _, p := range passlist {
		i += validatepass1(p)
	}
	return i
}

// Solve2 returns answer to second problem
func Solve2() int {
	var i int
	passlist := ParseLines(input)
	for _, p := range passlist {
		i += validatepass2(p)
	}
	return i
}
