package main

import (
	"bufio"
	"os"
	"strings"
)

func hash(input string) (h int) {
	/*
		Determine the ASCII code for the current character of the string.
		Increase the current value by the ASCII code you just determined.
		Set the current value to itself multiplied by 17.
		Set the current value to the remainder of dividing itself by 256.
	*/
	for _, c := range input {
		h = int(c) + h
		h = h * 17
		h = h % 256
	}
	return
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	txt := ""
	for scanner.Scan() {
		txt = txt + scanner.Text()
	}

	inputs := strings.Split(txt, ",")

	ret := 0
	for _, input := range inputs {
		ret += hash(input)
	}
	println(ret)

}
