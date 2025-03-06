package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Could not open the file")
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	left := make([]int, 1000)
	right := make([]int, 1000)

	count := 0

	for sc.Scan() {
		_, err := fmt.Sscanf(sc.Text(), "%d %d", &left[count], &right[count])
		if err != nil {
			fmt.Println(err)
		}

		count++
	}

	slices.Sort(left)
	slices.Sort(right)

	result := 0

	for i := 0; i < count; i++ {
		if x := left[i] - right[i]; x < 0 {
			result += -x
		} else {
			result += x
		}
	}

	fmt.Println(result)

	result = 0

	for i, j := 0, 0; i < count; i++ {
		c := 0
		for j < count && left[i] > right[j] {
			j++
		}

		for j < count && left[i] == right[j] {
			c++
			j++
		}

		result += left[i] * c
	}

	fmt.Println(result)

}
