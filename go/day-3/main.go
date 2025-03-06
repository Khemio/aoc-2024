package main

import (
	// "bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	// input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// sc := bufio.NewScanner(f)

	result := 0

	// for sc.Scan() {
	// 	result += compute(sc.Text())
	// }
	result += compute(string(input))

	fmt.Println(result)
}

func compute(input string) int {
	rs := []rune(input)
	result := 0
	enabled := true

	for i := 0; i < len(rs); i++ {
		if enabled && rs[i] == 'm' {

			if slices.Equal(rs[i:i+4], []rune("mul(")) {
				i += 4
				num1 := make([]rune, 0)
				for rs[i] >= '0' && rs[i] <= '9' {
					num1 = append(num1, rs[i])
					i++
				}

				if rs[i] != ',' {
					continue
				}

				i++

				num2 := make([]rune, 0)
				for rs[i] >= '0' && rs[i] <= '9' {
					num2 = append(num2, rs[i])
					i++
				}

				if rs[i] != ')' {
					continue
				}

				a, _ := strconv.Atoi(string(num1))
				b, _ := strconv.Atoi(string(num2))

				result += a * b
			}
		} else if rs[i] == 'd' {
			if slices.Equal(rs[i:i+4], []rune("do()")) {
				enabled = true
				i += 4
			} else if slices.Equal(rs[i:i+7], []rune("don't()")) {
				enabled = false
				i += 7
			}
		}
	}

	return result
}
