package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cards struct {
	winning []int
	own     []int
}

func convert(a []string) (ret []int) {
	for _, ch := range a {
		d, _ := strconv.Atoi(ch)
		ret = append(ret, d)
	}
	return
}

func contains(arr []int, a int) bool {
	for _, x := range arr {
		if x == a {
			return true
		}
	}
	return false
}

func pow2(x int) int {
	return int(math.Pow(2.0, float64(x)))
}

func part1(allCards []Cards) {
	ans := 0
	for _, cards := range allCards {
		cnt := 0
		for _, c := range cards.own {
			if contains(cards.winning, c) {
				cnt += 1
			}
		}
		ans += pow2(cnt - 1)
	}
	fmt.Println(ans)
}

func part2(allCards []Cards) {
	counts := make(map[int]int)
	for i := range allCards {
		counts[i]++
	}
	for i, cards := range allCards {
		cnt := 0
		for _, c := range cards.own {
			if contains(cards.winning, c) {
				cnt += 1
			}
		}
		for j := 0; j < cnt; j++ {
			counts[i+j+1] += counts[i]
		}
	}
	ret := 0
	for _, v := range counts {
		ret += v
	}
	fmt.Println(ret)
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	var allCards []Cards
	for scanner.Scan() {
		inps := strings.Split(strings.Split(scanner.Text(), ": ")[1], " | ")

		winning := strings.Fields(inps[0])
		own := strings.Fields(inps[1])
		allCards = append(allCards, Cards{convert(winning), convert(own)})
	}
	part1(allCards)
	part2(allCards)
}
