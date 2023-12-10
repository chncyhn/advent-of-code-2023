package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func applyMappings(input int, mappings Mappings) int {
	for _, mp := range mappings.mappings {
		output := apply(mp, input)
		if output != -1 {
			return output
		}
	}
	return input
}

func readInput() (seeds []int, mappings []Mappings) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds = mapToInt(strings.Split(strings.Replace(scanner.Text(), "seeds: ", "", 1), " "))
	scanner.Scan()
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
	return
}

func part1(seeds []int, mappings []Mappings) {
	frontier := seeds
	for _, mapping := range mappings {
		var newFrontier []int
		for _, input := range frontier {
			output := applyMappings(input, mapping)
			newFrontier = append(newFrontier, output)
		}
		frontier = newFrontier
	}
	fmt.Println(min(frontier))
}

func part2(seeds []int, mappings []Mappings) {
	var wg sync.WaitGroup
	resultChan := make(chan int, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			seed := seeds[ix]
			cnt := seeds[ix+1]
			minVal := 1 << 32
			for j := 0; j < cnt; j++ {
				cur := seed + j
				for _, mp := range mappings {
					cur = applyMappings(cur, mp)
				}
				if cur < minVal {
					minVal = cur
				}
			}
			resultChan <- minVal
		}(i)
	}

	wg.Wait()
	close(resultChan)

	minVal := 1 << 32
	for res := range resultChan {
		if res < minVal {
			minVal = res
		}
	}
	fmt.Println(minVal)
}

func main() {
	seeds, mappings := readInput()
	part1(seeds, mappings)
	part2(seeds, mappings)
}
