package main

import (
	"fmt"
	// "strings"
	"bufio"
	"os"
)

// "fmt"

func main() {
	// input := `MMMSXXMASM
	// 		MSAMXMSMSA
	// 		AMXSXMAAMM
	// 		MSAMASMSMX
	// 		XMASAMXAMM
	// 		XXAMMXXAMA
	// 		SMSMSASXSS
	// 		SAXAMASAAA
	// 		MAMMMXMMMM
	// 		MXMXAXMASX`

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	sc := bufio.NewScanner(f)

	rows := make([]string, 0)

	for sc.Scan() {
		rows = append(rows, sc.Text())
	}

	// rows := strings.Split(input, "\n")
	// for i, row := range rows {
	// 	rows[i] = strings.TrimSpace(row)
	// }

	count := 0

	rowN := len(rows) - 1
	colN := len(rows[0]) - 1

	for i, row := range rows {
		if i < 1 || i >= rowN {
			continue
		}
		for j, ch := range row {
			if j < 1 || j >= colN {
				continue
			}
			if ch == 'A' {
				if checkX(rows, i, j) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func checkXMAS(rows []string) int {
	count := 0

	for i1, row := range rows {
		for j1, ch := range row {

			if ch == 'X' {
				i2s, j2s := checkAround(rows, i1, j1)

				if len(i2s) == 0 || len(j2s) == 0 {
					continue
				}

				for k := range i2s {
					i3, j3 := checkDir(rows, i1, j1, i2s[k], j2s[k], 'A')

					if i3 < 0 || j3 < 0 {
						continue
					}

					i4, j4 := checkDir(rows, i2s[k], j2s[k], i3, j3, 'S')

					if i4 < 0 || j4 < 0 {
						continue
					}
					count++
				}

			}
		}
	}

	return count
}

func checkAround(rows []string, row, col int) ([]int, []int) {
	rowIdx, colIdx := make([]int, 0), make([]int, 0)
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i <= len(rows)-1 && j >= 0 && j <= len(rows[0])-1 {
				if rows[i][j] == 'M' {
					rowIdx = append(rowIdx, i)
					colIdx = append(colIdx, j)
				}
			}
		}
	}

	return rowIdx, colIdx
}

func checkDir(rows []string, i1, j1, i2, j2 int, ch byte) (int, int) {
	i3, j3 := -1, -1

	di := i2 - i1
	dj := j2 - j1

	if i2+di >= 0 && i2+di <= len(rows)-1 && j2+dj >= 0 && j2+dj <= len(rows[0])-1 {

		if rows[i2+di][j2+dj] == ch {
			i3 = i2 + di
			j3 = j2 + dj
		}
	}

	return i3, j3
}

func checkCh(rows []string, i, j int) (bool, byte) {

	if rows[i][j] == 'M' || rows[i][j] == 'S' {
		return true, rows[i][j]
	}

	return false, 0
}

func checkX(rows []string, i, j int) bool {
	hit, ch1 := checkCh(rows, i-1, j-1)
	if !hit {
		return false
	}

	hit, ch2 := checkCh(rows, i+1, j+1)
	if !hit || ch1 == ch2 {
		return false
	}

	hit, ch3 := checkCh(rows, i-1, j+1)
	if !hit {
		return false
	}

	hit, ch4 := checkCh(rows, i+1, j-1)
	if !hit || ch3 == ch4 {
		return false
	}

	return true
}
