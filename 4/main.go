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

func Part1() int {
	input := string(common.Input())
	rows := strings.Split(input, "\n")

	r := "XMAS"
	count := 0
	for i, row := range rows {
		for j, c := range row {
			if string(c) == "X" {
				for _, dx := range []int{0, 1, -1} {
					for _, dy := range []int{0, 1, -1} {
						count += search(rows, i, j, dx, dy, r)
					}
				}
			}
		}
	}
	return count
}

func search(in []string, sx, sy, dx, dy int, w string) int {
	rows := len(in)
	for k, m := range w {
		x := sx + dx*k
		if x >= rows || x < 0 {
			return 0
		}
		cols := len(in[x])
		y := sy + dy*k
		if y >= cols || y < 0 {
			return 0
		}
		if rune(in[x][y]) != m {
			return 0
		}
	}
	return 1
}

func Part2() int {
	input := string(common.Input())
	rows := strings.Split(input, "\n")

	count := 0
	s1 := "AM"
	s2 := "AS"
	for x, row := range rows {
		for y, c := range row {
			if string(c) == "A" {
				r0 := 0
				for _, dx := range []int{1, -1} {
					dy := 1
					for _, dir := range []int{1, -1} {
						r1 := search(rows, x, y, dir*dx, dir*dy, s1)
						r2 := search(rows, x, y, dir*-dx, dir*-dy, s2)
						if r1 == 1 && r2 == 1 {
							r0 += 1
							break
						}
					}
				}
                if r0 == 2 {
                    count += 1
                }
			}
		}
	}
	return count
}
