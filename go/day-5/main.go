package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, lines := getData("input.txt")
	count := 0

	for _, line := range lines {
		fixed := false
		for i, n := range line {
			vals := rules[n]
			correct, idx := inOrder((line[:i]), vals)
			for !correct {
				if !fixed {
					fixed = true
				}
				temp := line[idx]
				line[idx] = line[i]
				line[i] = temp

				correct, idx = inOrder((line[:i]), vals)
				vals = rules[line[i]]
			}
			if i == len(line)-1 && fixed {
				mid := i / 2
				count += line[mid]
			}
		}

	}

	fmt.Println(count)
}

func inOrder(nums []int, vals []int) (bool, int) {
	for i, n := range nums {
		if slices.Contains(vals, n) {
			return false, i
		}
	}
	return true, -1
}

func getData(path string) (map[int][]int, [][]int) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	data := make([]string, 0)
	breakIdx := 0

	for i := 0; sc.Scan(); i++ {
		if len(sc.Text()) == 0 {
			breakIdx = i
		}
		data = append(data, sc.Text())
	}

	rules := data[:breakIdx]
	lines := data[breakIdx+1:]

	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		tokens := strings.Split(rule, "|")
		key, _ := strconv.Atoi(tokens[0])
		val, _ := strconv.Atoi(tokens[1])
		ruleMap[key] = append(ruleMap[key], val)
	}

	pages := make([][]int, 0)

	for i, line := range lines {
		tokens := strings.Split(line, ",")
		pages = append(pages, make([]int, 0))

		for _, t := range tokens {
			n, _ := strconv.Atoi(t)
			pages[i] = append(pages[i], n)
		}
	}

	return ruleMap, pages
}
