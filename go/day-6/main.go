package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Pos struct {
	row, col int
}

type Dir struct {
	dr, dc int
}

var dirs = []rune{'^', '>', 'v', '<'}

var guard = ' '

var guardPos = Pos{0, 0}

var dirMap = map[rune]Dir{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func main() {

	// data := getData("test.txt")
	data := getData("input.txt")

	guard, guardPos = findGuard(data)

	guardPath := getGuardPath(data, guard, guardPos)

	room := data[:][:]

	loopCount := findLoops(room, guardPath)

	// fmt.Println(len(data))
	// fmt.Println(len(data[0]))

	fmt.Println(len(guardPath))
	fmt.Println(loopCount)

}

func findLoops(room []string, visited map[Pos]bool) int {
	count := 0
	// loopPoints := make(map[Pos]bool)

	for p := range visited {
		if p != guardPos {
			modifyPos(p, room, '#')

			if isLoop(room) {
				count++
				// loopPoints[p] = true
			}
			modifyPos(p, room, '.')
		}
	}

	return count
	// return len(loopPoints)
}

func modifyPos(p Pos, room []string, sym rune) {
	row := room[p.row : p.row+1]
	temp := []rune(row[0])
	temp[p.col] = sym
	row[0] = string(temp)
}

func isLoop(room []string) bool {
	pos := guardPos
	sym := guard

	curDir := dirMap[sym]
	lastDir := curDir

	visited := make(map[Pos][]Dir)
	// blocks := make(map[Pos]bool)

	for pos.row >= 0 && pos.row < len(room) && pos.col >= 0 && pos.col < len(room[0]) {
		if room[pos.row][pos.col] == '#' {
			pos.row -= curDir.dr
			pos.col -= curDir.dc

			sym = nextDir(sym)
			lastDir = curDir
			curDir = dirMap[sym]

		}

		if visited[pos] == nil {
			dirs := make([]Dir, 1)
			visited[pos] = dirs
		} else {
			dirs := visited[pos]

			for _, dir := range dirs {
				if dir == lastDir {
					return true
				}
			}

			dirs = append(dirs, lastDir)

			visited[pos] = dirs
		}
		pos.row += curDir.dr
		pos.col += curDir.dc
		// lastDir = curDir

	}

	return false
}

func nextDir(sym rune) rune {
	for i, s := range dirs {
		if s == sym && i < len(dirs)-1 {
			// new := dirs[i+1]
			return dirs[i+1]
		}
	}

	return dirs[0]
}

func getGuardPath(data []string, sym rune, pos Pos) map[Pos]bool {
	dir := dirMap[sym]

	visited := make(map[Pos]bool)

	for pos.row >= 0 && pos.row < len(data) && pos.col >= 0 && pos.col < len(data[0]) {

		if data[pos.row][pos.col] == '#' {
			pos.row -= dir.dr
			pos.col -= dir.dc

			sym = nextDir(sym)
			dir = dirMap[sym]

		}

		visited[pos] = true
		pos.row += dir.dr
		pos.col += dir.dc

	}

	return visited
}

func findGuard(data []string) (rune, Pos) {
	for i, row := range data {
		for j, col := range row {
			// if col == '^' {
			if slices.Contains(dirs, col) {
				return col, Pos{i, j}
			}
		}
	}

	return '0', Pos{-1, -1}
}

func getData(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	sc := bufio.NewScanner(f)

	data := make([]string, 0)

	for sc.Scan() {
		data = append(data, sc.Text())
	}

	return data
}
