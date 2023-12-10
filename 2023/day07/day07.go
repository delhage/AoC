package day07

import (
	"aoc/2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var input, _ = utils.ReadFile("day07/input.txt")

var cardMap = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	cards    []string
	handType int
	bid      int
	rank     int
}

func stringToHand(s string) Hand {
	var res Hand
	t := strings.Fields(s)
	for _, c := range t[0] {
		res.cards = append(res.cards, string(c))
	}
	res.handType = 0
	res.bid, _ = strconv.Atoi(t[1])
	res.rank = 0
	return res
}

func getType(h []string) int {
	var cardArray [15]int
	fmt.Println("h", h)
	for i := 0; i < len(h); i++ {
		//fmt.Println("cardMap[h[i]]", cardMap[h[i]])
		cardArray[cardMap[h[i]]]++
	}

	var pair = 0
	var three = 0

	for i := 0; i < len(cardArray); i++ {
		if cardArray[i] == 5 {
			return 6
		} else if cardArray[i] == 4 {
			return 5
		} else if cardArray[i] == 3 {
			three++
		} else if cardArray[i] == 2 {
			pair++
		} else if cardArray[i] == 1 {
			continue
		}
	}

	if pair == 1 && three == 1 {
		return 4
	}
	if three == 1 {
		return 3
	}
	if pair == 2 {
		return 2
	}
	if pair == 1 {
		return 1
	}

	//fmt.Println("cardArray", cardArray)
	return 0
}

// Solve1 returns answer to first problem
func Solve1() int {
	var res = 0
	var hands []Hand

	for _, s := range input {
		h1 := stringToHand(s)
		h1.handType = getType(h1.cards)
		fmt.Println("t1", h1.handType)
		hands = append(hands, h1)
	}
	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	return 1
}
