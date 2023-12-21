package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func mapToInt(s []string) (ret []int) {
	for _, v := range s {
		dig, _ := strconv.Atoi(v)
		ret = append(ret, dig)
	}
	return
}

func isAllZeroes(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

func solve(input []int) (left, right int) {
	var diffArrs [][]int
	diffArrs = append(diffArrs, input)
	for {
		var diffArr []int
		baseArr := diffArrs[len(diffArrs)-1]
		for i := 0; i < len(baseArr)-1; i++ {
			diffArr = append(diffArr, baseArr[i+1]-baseArr[i])
		}
		diffArrs = append(diffArrs, diffArr)
		if isAllZeroes(diffArr) {
			break
		}
	}
	for i := len(diffArrs) - 1; i >= 0; i-- {
		right += diffArrs[i][len(diffArrs[i])-1]
		left = diffArrs[i][0] - left
	}
	return
}

func readInput() (ret [][]int) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, mapToInt(strings.Split(scanner.Text(), " ")))
	}
	return
}

func main() {
	inputs := readInput()
	var part1, part2 int
	for _, input := range inputs {
		l, r := solve(input)
		part1 += r
		part2 += l
	}
	println(part1, part2)
}
