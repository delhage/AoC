package day07

import (
	"aoc/2023/utils"
	"sort"
	"strconv"
	"strings"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

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

var cardMap2 = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type Hand struct {
	cards    []string
	handType int
	bid      int
}

type HandSlice []Hand

func (h HandSlice) Len() int { return len(h) }
func (h HandSlice) Less(i, j int) bool {
	if h[i].handType < h[j].handType {
		return true
	} else if h[i].handType == h[j].handType {
		for k := 0; k < 5; k++ {
			if cardMap[h[i].cards[k]] < cardMap[h[j].cards[k]] {
				return true
			} else if cardMap[h[i].cards[k]] > cardMap[h[j].cards[k]] {
				return false
			}
		}
	}
	return false
}
func (h HandSlice) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func stringToHand(s string) Hand {
	var res Hand
	t := strings.Fields(s)
	for _, c := range t[0] {
		res.cards = append(res.cards, string(c))
	}
	res.handType = 0
	res.bid, _ = strconv.Atoi(t[1])
	return res
}

func getType(h []string) int {
	var cardArray [15]int
	//fmt.Println("h", h)
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
	var hands HandSlice

	for _, s := range input {
		h1 := stringToHand(s)
		h1.handType = getType(h1.cards)
		//fmt.Println("t1", h1.handType)
		hands = append(hands, h1)
	}
	//fmt.Println(hands)
	sort.Stable(HandSlice(hands))
	//fmt.Println(hands)
	for i := 0; i < len(hands); i++ {
		res += hands[i].bid * (i + 1)
	}
	return res
}

// Solve2 returns answer to second problem
func Solve2() int {
	return 1
}
