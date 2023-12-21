package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instructions string
type Node string
type Graph map[Node][]Node

func solveForNode(node Node, graph Graph, instructions Instructions) int {
	cur := node
	steps := 0
	for {
		for _, i := range instructions {
			if i == 'L' {
				cur = graph[cur][0]
			} else {
				cur = graph[cur][1]
			}
			steps += 1
			if strings.HasSuffix(string(cur), "Z") {
				return steps
			}
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmOfAll(numbers []int) int {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = lcm(result, num)
	}
	return result
}

func readInput() (Instructions, Graph) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := Instructions(scanner.Text())
	scanner.Scan()
	graph := make(Graph)
	for scanner.Scan() {
		var from, left, right string
		s := strings.ReplaceAll(scanner.Text(), "(", "")
		s = strings.ReplaceAll(s, ")", "")
		s = strings.ReplaceAll(s, ",", "")
		fmt.Sscanf(s, "%s = %s %s", &from, &left, &right)
		graph[Node(from)] = []Node{Node(left), Node(right)}
	}
	return instructions, graph
}

func main() {
	instructions, graph := readInput()

	var solutions []int
	for k := range graph {
		if strings.HasSuffix(string(k), "A") {
			solutions = append(solutions, solveForNode(k, graph, instructions))
		}
	}
	println(lcmOfAll(solutions))

}
