package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// f, err := os.Open("test2.txt")
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	safe := 0
	faulty := 0

	for sc.Scan() {
		line := sc.Text()
		tokens := strings.Split(line, " ")
		s, idx := isSafe(tokens)

		if s {
			safe++
		} else if faultySafe(tokens, idx) {
			// if lazySafe(tokens) {
			faulty++
			// }

		}

	}

	fmt.Println("safe: ", safe)
	fmt.Println("faulty: ", faulty)
	fmt.Println("total: ", safe+faulty)
}

func faultySafe(tokens []string, i int) bool {
	// hasFault := false

	// temp := make([]string, len(tokens))

	if i == len(tokens)-2 {
		x, err := strconv.Atoi(tokens[i-2])
		if err != nil {
			fmt.Println(err)
		}
		a, err := strconv.Atoi(tokens[i-1])
		if err != nil {
			fmt.Println(err)
		}
		// b, err := strconv.Atoi(tokens[i])
		// if err != nil {
		// 	fmt.Println(err)
		// }
		c, err := strconv.Atoi(tokens[i+1])
		if err != nil {
			fmt.Println(err)
		}

		sign := getSign(x - a)
		diff1 := a - c
		// diff2 := b - c

		// if abs(diff2) >= 1 || abs(diff2) <= 3 {
		// 	f := i
		// 	safe, _ := isSafe(append(tokens[:f], tokens[f+1:]...))
		// 	return safe

		// }

		if sign == getSign(diff1) && (abs(diff1) >= 1 || abs(diff1) <= 3) {
			f := i
			safe, _ := isSafe(append(tokens[:f], tokens[f+1:]...))
			return safe
		}

		f := i + 1
		// safe, _ := isSafe(append(tokens[:f], tokens[f+1:]...))
		safe, _ := isSafe(tokens[:f])
		return safe

	}

	if i == 0 {
		a, err := strconv.Atoi(tokens[i])
		if err != nil {
			fmt.Println(err)
		}
		b, err := strconv.Atoi(tokens[i+1])
		if err != nil {
			fmt.Println(err)
		}
		c, err := strconv.Atoi(tokens[i+2])
		if err != nil {
			fmt.Println(err)
		}

		// sign := getSign(a - b)
		diff1 := a - c
		diff2 := b - c

		if abs(diff2) >= 1 || abs(diff2) <= 3 {
			f := i
			safe, _ := isSafe(append(tokens[:f], tokens[f+1:]...))
			return safe

		}

		if abs(diff1) >= 1 || abs(diff1) <= 3 {
			f := i + 1
			safe, _ := isSafe(append(tokens[:f], tokens[f+1:]...))
			return safe
		}

		return false
	}

	return false
}

func isSafe(tokens []string) (bool, int) {
	// safe := true

	sign := 0

	// hasFault := false

	for i := 0; i < len(tokens)-1; i++ {
		a, err := strconv.Atoi(tokens[i])
		if err != nil {
			fmt.Println(err)
		}

		b, err := strconv.Atoi(tokens[i+1])
		if err != nil {
			fmt.Println(err)
		}

		diff := a - b

		if sign == 0 {
			sign = getSign(diff)
		}

		if sign < 0 {
			if diff < -3 || diff > -1 {
				return false, i
			}
		} else {
			if diff < 1 || diff > 3 {
				return false, i
			}
		}
	}

	return true, -1
}

func lazySafe(tokens []string) bool {
	temp := make([]string, len(tokens))
	for i := 0; i < len(tokens); i++ {

		copy(temp, tokens)
		new := append(temp[:i], temp[i+1:]...)

		s, _ := isSafe(new)

		// fmt.Println(new)
		if s {
			// faulty++
			// break
			return true
			// errCount++
		}
	}

	return false
}

func getSign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
