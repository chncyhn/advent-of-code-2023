package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand [5]byte

type HandAndMorph struct {
	hand  Hand
	morph Hand
}

func counts(hand Hand) map[byte]int {
	ret := make(map[byte]int)
	for _, c := range hand {
		ret[c]++
	}
	return ret
}

func morph(hand Hand) Hand {
	cnts := counts(hand)
	max := 0
	for k, v := range cnts {
		if k == 'J' {
			continue
		}
		if v > max {
			max = v
		}
	}
	highestWithMax := byte('-')
	for k, v := range cnts {
		if k == 'J' {
			continue
		}
		if v == max {
			if highestWithMax == byte('-') || value(k) > value(highestWithMax) {
				highestWithMax = k
			}
		}
	}
	if highestWithMax == byte('-') {
		return Hand{'A', 'A', 'A', 'A', 'A'}
	}
	morphedHand := Hand{}
	for i, c := range hand {
		if c == 'J' {
			morphedHand[i] = highestWithMax
		} else {
			morphedHand[i] = c
		}
	}
	return morphedHand
}

func toHand(str string) Hand {
	var ret Hand
	for i := 0; i < 5; i++ {
		ret[i] = str[i]
	}
	return ret
}

func label(hand Hand) int {
	c := counts(hand)
	if len(c) == 1 {
		return 7
	}
	if len(c) == 2 {
		for _, v := range c {
			if v == 4 {
				return 6
			}
		}
		return 5
	}
	if len(c) == 3 {
		for _, v := range c {
			if v == 3 {
				return 4
			}
		}
		return 3
	}
	if len(c) == 4 {
		return 2
	}
	return 1
}

func value(card byte) int {
	if card == 'T' {
		return 10
	}
	if card == 'J' {
		return 11
	}
	if card == 'Q' {
		return 12
	}
	if card == 'K' {
		return 13
	}
	if card == 'A' {
		return 14
	}
	return int(card - '0')
}

func value2(card byte) int {
	if card == 'J' {
		return 0
	}
	return value(card)
}

func part1(handToRank map[Hand]int, hands []HandAndMorph) {
	sort.Slice(hands, func(i, j int) bool {
		li, lj := label(hands[i].hand), label(hands[j].hand)
		if li != lj {
			return li < lj
		}
		for k := 0; k < 5; k++ {
			if value(hands[i].hand[k]) != value(hands[j].hand[k]) {
				return value(hands[i].hand[k]) < value(hands[j].hand[k])
			}
		}
		return false
	})
	ret := 0
	for i, hand := range hands {
		ret += handToRank[hand.hand] * (i + 1)
	}
	println(ret)
}

func part2(handToRank map[Hand]int, hands []HandAndMorph) {
	sort.Slice(hands, func(i, j int) bool {
		li, lj := label(hands[i].morph), label(hands[j].morph)
		if li != lj {
			return li < lj
		}
		for k := 0; k < 5; k++ {
			if value2(hands[i].hand[k]) != value2(hands[j].hand[k]) {
				return value2(hands[i].hand[k]) < value2(hands[j].hand[k])
			}
		}
		return false
	})
	ret := 0
	for i, hand := range hands {
		ret += handToRank[hand.hand] * (i + 1)
	}
	println(ret)
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	handToRank := make(map[Hand]int)
	var hands []HandAndMorph
	for scanner.Scan() {
		inp := strings.Split(scanner.Text(), " ")
		hand := toHand(inp[0])
		rank, _ := strconv.Atoi(inp[1])
		handToRank[hand] = rank
		hands = append(hands, HandAndMorph{hand, morph(hand)})
	}

	part1(handToRank, hands)
	part2(handToRank, hands)
}
