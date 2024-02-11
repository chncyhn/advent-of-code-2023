package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	label string
	op    string
	value int
}

type LabelAndValue struct {
	label string
	value int
}

func hash(input string) (h int) {
	for _, c := range input {
		h = int(c) + h
		h = h * 17
		h = h % 256
	}
	return
}

func parseCommands(inputs []string) (cmds []Command) {
	for _, input := range inputs {
		i := strings.Index(input, "=")
		if i == -1 {
			i = strings.Index(input, "-")
			cmds = append(cmds, Command{label: input[:i], op: "-"})
		} else {
			value, _ := strconv.Atoi(input[i+1:])
			cmds = append(cmds, Command{label: input[:i], op: "=", value: value})
		}
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
	cmds := parseCommands(inputs)

	boxes := make(map[int][]LabelAndValue)
	for _, cmd := range cmds {
		if cmd.op == "-" {
			loc := hash(cmd.label)
			box, ok := boxes[loc]
			if !ok {
				continue
			}
			for i, b := range box {
				if b.label == cmd.label {
					boxes[loc] = append(box[:i], box[(i+1):]...)
				}
			}
		}
		if cmd.op == "=" {
			loc := hash(cmd.label)
			box, ok := boxes[loc]
			if !ok {
				box = []LabelAndValue{}
			}
			updated := false
			for i, b := range box {
				if b.label == cmd.label {
					box[i] = LabelAndValue{label: cmd.label, value: cmd.value}
					updated = true
					break
				}
			}
			if !updated {
				box = append(box, LabelAndValue{label: cmd.label, value: cmd.value})
				boxes[loc] = box
			}
		}
	}

	ret := 0
	for k, box := range boxes {
		for i, labelAndVal := range box {
			ret += (k + 1) * (i + 1) * labelAndVal.value
		}
	}
	println(ret)
}
