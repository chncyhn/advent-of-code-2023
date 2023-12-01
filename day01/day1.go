package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var NUMBERS = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type ValIx struct {
	val int
	ix  int
}

func isNum(c rune) bool {
	return c-'0' >= 0 && c-'0' <= 9
}

func part1(lines []string) {
	var total int
	for _, line := range lines {
		var lo = -1
		var hi = -1
		for _, c := range line {
			if isNum(c) {
				if lo == -1 {
					lo = int(c) - '0'
				}
				hi = int(c) - '0'
			}
		}
		total += 10*lo + hi
	}
	fmt.Println(total)
}

func part2(lines []string) {
	var total int
	for _, line := range lines {
		var lo = ValIx{-1, -1}
		var hi = ValIx{-1, -1}
		for dig, val := range NUMBERS {
			if ix := strings.Index(line, dig); ix != -1 {
				if lo.ix == -1 || ix < lo.ix {
					lo = ValIx{val, ix}
				}
			}
			if ix := strings.LastIndex(line, dig); ix != -1 {
				if hi.ix == -1 || ix > hi.ix {
					hi = ValIx{val, ix}
				}
			}
		}
		for i, c := range line {
			if isNum(c) {
				if lo.ix == -1 || i < lo.ix {
					lo = ValIx{int(c) - '0', i}
				}
				if hi.ix == -1 || i > hi.ix {
					hi = ValIx{int(c) - '0', i}
				}
			}
		}
		total += 10*lo.val + hi.val
	}
	fmt.Println(total)
}

func main() {
	file, err := os.Open("full.txt")
	if err != nil {
		panic(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1(lines)
	part2(lines)
}
