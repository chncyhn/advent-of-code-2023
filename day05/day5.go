package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Mapping struct {
	sourceStart int
	destStart   int
	count       int
}

type SourceDest struct {
	source, dest string
}

type Mappings struct {
	sd       SourceDest
	mappings []Mapping
}

func mapToInt(seeds []string) (ret []int) {
	for _, s := range seeds {
		dig, _ := strconv.Atoi(s)
		ret = append(ret, dig)
	}
	return
}

func startsWithLetter(str string) bool {
	for _, c := range str {
		return unicode.IsLetter(c)
	}
	panic("empty string!")
}

func apply(mapping Mapping, input int) int {
	if mapping.sourceStart > input || input > mapping.sourceStart+mapping.count {
		return -1
	}
	return mapping.destStart + (input - mapping.sourceStart)
}

func min(vals []int) int {
	m := vals[0]
	for _, v := range vals {
		if v < m {
			m = v
		}
	}
	return m
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := mapToInt(strings.Split(strings.Replace(scanner.Text(), "seeds: ", "", 1), " "))
	scanner.Scan()

	var mappings []Mappings

	sd := SourceDest{"", ""}
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}
		if startsWithLetter(str) {
			sdi := strings.Split(strings.Replace(str, " map:", "", 1), "-to-")
			sd = SourceDest{sdi[0], sdi[1]}
			mappings = append(mappings, Mappings{sd, []Mapping{}})

		} else {
			dsc := mapToInt(strings.Split(str, " "))
			mp := Mapping{sourceStart: dsc[1], destStart: dsc[0], count: dsc[2]}
			mappings[len(mappings)-1].mappings = append(mappings[len(mappings)-1].mappings, mp)
		}
	}

	frontier := seeds
	for _, mapping := range mappings {
		var newFrontier []int
		for _, input := range frontier {
			mapped := false
			for _, mp := range mapping.mappings {
				output := apply(mp, input)
				if output != -1 {
					newFrontier = append(newFrontier, output)
					mapped = true
					break
				}
			}
			if !mapped {
				newFrontier = append(newFrontier, input)
			}
		}
		frontier = newFrontier
	}

	fmt.Println(min(frontier))
}
