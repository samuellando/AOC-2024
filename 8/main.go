package main

import (
	"advent/common"
	"fmt"
	"iter"
	"math"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	antenas := make(map[rune][][]int)
	input := string(common.Input())
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	nrows := len(rows)
	ncols := len(rows[0])
	for y, row := range rows {
		for x, c := range row {
			if c == '.' {
				continue
			}
			if _, ok := antenas[c]; !ok {
				antenas[c] = make([][]int, 0)
			}
			antenas[c] = append(antenas[c], []int{x, y})
		}
	}
	count := 0
	anodes := make(map[int]map[int]bool, 0)
	for c := range antenas {
		for _, pair := range pairs(antenas[c]) {
			points := calc(pair[0], pair[1])
			for _, point := range points {
				x := point[0]
				y := point[1]
				if x >= 0 && x < nrows && y >= 0 && y < ncols {
					if _, ok := anodes[x]; !ok {
						anodes[x] = make(map[int]bool)
					}
					if _, ok := anodes[x][y]; !ok {
						anodes[x][y] = true
						count += 1
					}
				}
			}
		}
	}
	return count
}

func calc(p1, p2 []int) [][]int {
	x1 := p1[0]
	y1 := p1[1]
	x2 := p2[0]
	y2 := p2[1]

	if x1 > x2 {
		return calc(p2, p1)
	}

	if x1 == x2 {
		dy := abs(y1 - y2)
		if y1 > y2 {
			return [][]int{{x1, y1 + dy}, {x1, y2 - dy}}
		} else {
			return [][]int{{x1, y2 + dy}, {x1, y2 - dy}}
		}
	}

	dx := x2 - x1
	dy := y2 - y1

	r1 := []int{x1 - dx, y1 - dy}
	r2 := []int{x2 + dx, y2 + dy}

	return [][]int{r1, r2}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pairs(l [][]int) [][][]int {
	res := make([][][]int, 0)
	if len(l) < 2 {
		return nil
	}
	if len(l) == 2 {
		res = append(res, l)
		return res
	}
	for _, pos := range l[1:] {
		res = append(res, [][]int{l[0], pos})
	}
	res = append(res, pairs(l[1:])...)
	return res
}

func Part2() int {
	antenas := make(map[rune][][]int)
	input := string(common.Input())

	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	nrows := len(rows)
	ncols := len(rows[0])
	for y, row := range rows {
		for x, c := range row {
			if c == '.' {
				continue
			}
			if _, ok := antenas[c]; !ok {
				antenas[c] = make([][]int, 0)
			}
			antenas[c] = append(antenas[c], []int{x, y})
		}
	}
	count := 0
	anodes := make(map[int]map[int]bool, 0)
	for c := range antenas {
		for _, pair := range pairs(antenas[c]) {
			points := calcLine(pair[0], pair[1], nrows)
			for point := range points {
				x := point[0]
				y := point[1]
				if x >= 0 && x < nrows && y >= 0 && y < ncols {
					if _, ok := anodes[x]; !ok {
						anodes[x] = make(map[int]bool)
					}
					if _, ok := anodes[x][y]; !ok {
						anodes[x][y] = true
						count += 1
					}
				}
			}
		}
	}
	return count
}

func calcLine(p1, p2 []int, s int) iter.Seq[[]int] {
	x1 := p1[0]
	y1 := p1[1]
	x2 := p2[0]
	y2 := p2[1]

	if x1 == x2 {
		return func(yeild func([]int) bool) {
			for y := range s {
				if !yeild([]int{x1, y}) {
					return
				}
			}
		}
	}

	dx := x2 - x1
	dy := y2 - y1
	a := float64(dy) / float64(dx)
	b := float64(y1) - a*float64(x1)
	return func(yeild func([]int) bool) {
		for x := range s {
			y := a*float64(x) + b
			var yint int
			isint := false
			if math.Abs(math.Ceil(y)-y) < 0.0001 {
				yint = int(math.Ceil(y))
				isint = true
			} else if math.Abs(y-math.Floor(y)) < 0.0001 {
				yint = int(math.Floor(y))
				isint = true
			}
			if isint {
				if !yeild([]int{x, yint}) {
					return
				}
			}
		}
	}
}
