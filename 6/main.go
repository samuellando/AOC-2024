package main

import (
	"advent/common"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func getGrid() [][]rune {
	input := string(common.Input())
	rows := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(rows))
	for _, row := range rows {
		if len(row) != 0 {
			grid = append(grid, []rune(row))
		}
	}
	return grid
}

func display(grid [][]rune) {
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func Part1() int {
	grid := getGrid()
	pos := []int{0, 0}
	for i, row := range grid {
		for j, c := range row {
			if c == '^' {
				pos[1] = i
				pos[0] = j
			}
		}
	}
	d := []int{0, -1}
	count := 1
	for pos[1] < len(grid) && pos[0] >= 0 && pos[0] < len(grid[0]) && pos[0] >= 0 {
		next := []int{pos[0] + d[0], pos[1] + d[1]}
		if next[1] < len(grid) && next[0] >= 0 && next[0] < len(grid[0]) && next[0] >= 0 {
			nextC := grid[next[1]][next[0]]
			if nextC == '#' {
				rotate(d)
				continue
			}
			if nextC != 'X' {
				count += 1
			}
			grid[next[1]][next[0]] = 'X'
		}
		pos[0] = next[0]
		pos[1] = next[1]
	}
	return count
}

func rotate(d []int) {
	if d[0] == 0 && d[1] == -1 {
		d[0] = 1
		d[1] = 0
	} else if d[0] == -1 && d[1] == 0 {
		d[0] = 0
		d[1] = -1
	} else if d[0] == 0 && d[1] == 1 {
		d[0] = -1
		d[1] = 0
	} else if d[0] == 1 && d[1] == 0 {
		d[0] = 0
		d[1] = 1
	}
}

func getDc(d []int) rune {
	if d[0] == 0 && d[1] == -1 {
		return 'A'
	} else if d[0] == -1 && d[1] == 0 {
		return 'B'
	} else if d[0] == 0 && d[1] == 1 {
		return 'C'
	} else {
		return 'D'
	}
}

func Part2() int {
	grid := getGrid()
	pos := []int{0, 0}
	for i, row := range grid {
		for j, c := range row {
			if c == '^' {
				pos[1] = i
				pos[0] = j
			}
		}
	}
	start := []int{pos[0], pos[1]}
	count := 0
	for i, row := range grid {
		for j := range row {
			d := []int{0, -1}
			dc := getDc(d)
			pos[0] = start[0]
			pos[1] = start[1]
			if grid[i][j] != '.' {
				continue
			}
			grid[i][j] = '#'
			for pos[1] < len(grid) && pos[1] >= 0 && pos[0] < len(grid[0]) && pos[0] >= 0 {
				next := []int{pos[0] + d[0], pos[1] + d[1]}
				if next[1] < len(grid) && next[1] >= 0 && next[0] < len(grid[0]) && next[0] >= 0 {
					nextC := grid[next[1]][next[0]]
					if nextC == '#' {
						rotate(d)
						dc = getDc(d)
						continue
					}
					if nextC == dc {
						count++
						break
					}
					grid[next[1]][next[0]] = dc
				}
				pos[0] = next[0]
				pos[1] = next[1]
			}
			grid[i][j] = '.'
			grid[start[1]][start[0]] = '^'
			for a, row := range grid {
				for b, c := range row {
					if c != '#' && c != '.' && c != '^' {
						grid[a][b] = '.'
					}
				}
			}
		}
	}
	return count
}
