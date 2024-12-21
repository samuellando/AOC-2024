package main

import (
	"advent/common"
	"fmt"
	"iter"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type point struct {
	x int
	y int
}

type board [][]rune

func (b *board) get(p point) rune {
	return (*b)[p.y][p.x]
}

func (b *board) set(p point, r rune) {
	(*b)[p.y][p.x] = r
}

func Part1() int {
	b := loadBoard(common.InputLines())
	movements := loadMovements(common.InputLines())
	p := findPos(b)
	for _, m := range movements {
		move(m, &p, b)
	}
	gpsSum := 0
	for _, bp := range getBoxPositions(b) {
		gpsSum += bp.x + 100*bp.y
	}
	return gpsSum
}

func Part2() int {
	b := loadBoard(common.InputLines())
	scaleBoard(b)
	movements := loadMovements(common.InputLines())
	p := findPos(b)
	for _, m := range movements {
		move(m, &p, b)
	}
	gpsSum := 0
	for _, bp := range getBoxPositions(b) {
		gpsSum += bp.x + 100*bp.y
	}
	return gpsSum
}

func display(b board) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
	for _, row := range b {
		for _, c := range row {
			if c == '@' {
				fmt.Print(colorRed, string(c), colorNone)
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}

}

func scaleBoard(b board) {
	for i, row := range b {
		newRow := make([]rune, 0, 2*len(row))
		for _, c := range row {
			switch c {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '@':
				newRow = append(newRow, '@', '.')
			case '.':
				newRow = append(newRow, '.', '.')
			}
		}
		b[i] = newRow
	}
}

func getBoxPositions(b board) []point {
	bps := make([]point, 0)
	for y, row := range b {
		for x, c := range row {
			if c == 'O' || c == '[' {
				bps = append(bps, point{x, y})
			}
		}
	}
	return bps
}

func loadBoard(in iter.Seq[[]string]) board {
	board := make([][]rune, 0)
	for row := range in {
		if len(row) == 0 {
			break
		}
		s := row[0]
		board = append(board, []rune(s))
	}
	return board
}

func loadMovements(in iter.Seq[[]string]) []rune {
	movements := make([]rune, 0)
	passedBoard := false
	for row := range in {
		if len(row) == 0 {
			passedBoard = true
		} else if passedBoard {
			s := row[0]
			movements = append(movements, []rune(s)...)
		}
	}
	return movements
}

func findPos(b [][]rune) point {
	var pos point
	for y, row := range b {
		for x, c := range row {
			if c == '@' {
				pos = point{x, y}
			}
		}
	}
	return pos
}

func vec(d rune) point {
	switch d {
	case '<':
		return point{-1, 0}
	case '>':
		return point{1, 0}
	case 'v':
		return point{0, 1}
	case '^':
		return point{0, -1}
	default:
		panic("bad move inst")
	}
}

func (p point) add(v point) point {
	return point{p.x + v.x, p.y + v.y}
}

func move(d rune, p *point, b board) {
	if !checkBlocked(d, *p, b) {
		v := vec(d)
		buff := b.get(*p)
		b.set(*p, '.')
		c := point{p.x, p.y}
		p.x = p.add(v).x
		p.y = p.add(v).y
		for range 10000 {
			c = c.add(v)
			switch b.get(c) {
			case '.':
				b.set(c, buff)
				return
			case 'O':
				b.set(c, buff)
				buff = 'O'
			case '[':
				b.set(c, buff)
				buff = '['
				if v.y != 0 {
					adj := point{c.x + 1, c.y}
					move(d, &adj, b)
				}
				continue
			case ']':
				b.set(c, buff)
				buff = ']'
				if v.y != 0 {
					adj := point{c.x - 1, c.y}
					move(d, &adj, b)
				}
				continue
			case '#':
				panic("lies")
			}
		}
	}
}

func checkBlocked(d rune, p point, b board) bool {
	v := vec(d)
	c := p
	for range 10000 {
		c = c.add(v)
		switch b.get(c) {
		case '.':
			return false
		case 'O':
			continue
		case '[':
			if v.y != 0 {
				adj := point{c.x + 1, c.y}
				if checkBlocked(d, adj, b) {
					return true
				}
			}
			continue
		case ']':
			if v.y != 0 {
				adj := point{c.x - 1, c.y}
				if checkBlocked(d, adj, b) {
					return true
				}
			}
			continue
		case '#':
			return true
		default:
			panic("invalid char in view")
		}
	}
	return true
}
