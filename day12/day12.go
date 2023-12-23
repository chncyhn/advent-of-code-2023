package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	arr   []int
	contg []int
}

var DIG_MAP = map[rune]int{
	'.': 0,
	'#': 1,
	'?': 2,
}

func mapToInt(A []string) (B []int) {
	for _, v := range A {
		dig, _ := strconv.Atoi(v)
		B = append(B, dig)
	}
	return
}

func readInput(fl string) (problems []Problem) {
	file, _ := os.Open(fl)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		arrS := s[0]
		var arr []int
		for i := 0; i < 5; i++ {
			for _, v := range arrS {
				arr = append(arr, DIG_MAP[v])
			}
			if i != 4 {
				arr = append(arr, 2)
			}
		}

		var contgS []int
		for i := 0; i < 5; i++ {
			contgS = append(contgS, mapToInt(strings.Split(s[1], ","))...)
		}

		problems = append(problems, Problem{arr, contgS})
	}
	return
}

func dfs(i, j int, p Problem) (cnt int) {
	if i >= len(p.arr) {
		if j == len(p.contg) {
			return 1
		}
		return 0
	}

	// if there's a tile here, it should cover the current contigous tile
	if p.arr[i] == 1 {
		if j >= len(p.contg) {
			return 0
		}
		for k := i; k < i+p.contg[j]; k++ {
			if k >= len(p.arr) || p.arr[k] == 0 {
				return 0
			}
		}
		ending := i + p.contg[j]
		if ending != len(p.arr) && p.arr[ending] == 1 {
			return 0
		}
		return dfs(i+p.contg[j]+1, j+1, p)
	}

	// check if we can place contigous tiles here
	canPlace := j < len(p.contg) && p.arr[i] == 2
	if canPlace {
		for k := i; k < i+p.contg[j]; k++ {
			if k >= len(p.arr) || p.arr[k] == 0 {
				canPlace = false
				break
			}
		}
		if canPlace {
			ending := i + p.contg[j]
			canPlace = canPlace && (ending == len(p.arr) || p.arr[ending] != 1)
		}
	}

	ret := 0
	if canPlace {
		ret += dfs(i+p.contg[j]+1, j+1, p)
	}

	// go without placing contigous tiles
	ret += dfs(i+1, j, p)
	return ret

}

func main() {
	problems := readInput("full.txt")

	ret := 0
	for _, p := range problems {
		ret += dfs(0, 0, p)
		println(ret)
	}

	println(ret)

}
