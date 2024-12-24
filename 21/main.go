package main

import (
	"advent/common"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type meta struct {
	val rune
	pos point
}

type node common.Node[meta]

type point struct {
	x int
	y int
}

type robot struct {
	nm         [][]common.Node[meta]
	keyPad     map[rune]point
	reversePad map[point]rune
	position   point
}

func newRobot() *robot {
	return &robot{make([][]common.Node[meta], 0), make(map[rune]point), make(map[point]rune), point{0, 0}}
}

func (r *robot) loadKeypad(m [][]rune) {
	r.nm = make([][]common.Node[meta], len(m))
	for y, row := range m {
		r.nm[y] = make([]common.Node[meta], len(row))
		for x, v := range row {
			p := point{x, y}
			r.keyPad[v] = p
			r.reversePad[p] = v
			r.nm[y][x] = common.CreateIndexedAdjNode(meta{v, p})
		}
	}
	common.ConnectAdjs(r.nm, func(a, b common.Node[meta]) bool {
		if a.GetValue().val != '-' && b.GetValue().val != '-' {
			return true
		} else {
			return false
		}
	})
}

func (r *robot) setPosition(v rune) {
	r.position = r.keyPad[v]
}

func (r *robot) getInputs(ss []string) []string {
	allResults := make([]string, 0)
	initialPos := r.position
	for _, s := range ss {
		results := make([]string, 0)
		for _, v := range s {
			result := r.pathsTo(v)
			newResults := make([]string, 0)
			if len(results) == 0 {
				for _, r := range result {
					newResults = append(newResults, r+"A")
				}
			} else {
				for _, rs := range results {
					for _, r := range result {
						newResults = append(newResults, rs+r+"A")
					}
				}
			}
			results = newResults
			r.setPosition(v)
		}
		allResults = append(allResults, results...)
	}
	r.position = initialPos
	minl := len(allResults[0])
	for _, r := range allResults {
		if len(r) < minl {
			minl = len(r)
		}
	}
	best := make([]string, 0)
	for _, r := range allResults {
		if len(r) == minl {
			best = append(best, r)
		}
	}
	return best
}

func (r *robot) pathsTo(v rune) []string {
	initialPos := r.position
	start := r.nm[initialPos.y][initialPos.x]
	finalPos := r.keyPad[v]
	end := r.nm[finalPos.y][finalPos.x]
	return dfs(start, end, "", make(map[node]bool))
}

type cachedResult struct {
	start  node
	end    node
	result []string
}

var cache = make(map[node]map[node][]string)

func dfs(start, end node, p string, visited map[node]bool) []string {
	if p == "" {
		if s, ok := cache[start]; ok {
			if r, ok := s[end]; ok {
				return r
			}
		}
	}
	if start == end {
		return []string{p}
	}
	paths := make([]string, 0)
	npos := start.GetValue().pos
	visited[start] = true
	for _, adj := range start.GetAdj() {
		if v, ok := visited[adj]; ok && v {
			continue
		}
		apos := adj.GetValue().pos
		var dir string
		if apos.x > npos.x {
			dir = ">"
		}
		if apos.x < npos.x {
			dir = "<"
		}
		if apos.y < npos.y {
			dir = "^"
		}
		if apos.y > npos.y {
			dir = "v"
		}
		r := dfs(adj, end, p+dir, visited)
		paths = append(paths, r...)
	}
	visited[start] = false
	if len(paths) == 0 {
		return paths
	}
	minl := len(paths[0])
	for _, p := range paths {
		if len(p) < minl {
			minl = len(p)
		}
	}
	best := make([]string, 0)
	for _, p := range paths {
		if len(p) == minl {
			best = append(best, p)
		}
	}
	if p == "" {
		if _, ok := cache[start]; !ok {
            cache[start] = make(map[node][]string)
        }
        cache[start][end] = best
	}
	return best
}

func (r *robot) run(s string) string {
	output := ""
	for _, c := range s {
		switch c {
		case '>':
			r.position.x += 1
		case '<':
			r.position.x -= 1
		case 'v':
			r.position.y += 1
		case '^':
			r.position.y -= 1
		case 'A':
			output += string(r.reversePad[r.position])
		}
	}
	return output
}

func Part1() int {
	keyPad := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'-', '0', 'A'},
	}
	directionalKeyPad := [][]rune{
		{'-', '^', 'A'},
		{'<', 'v', '>'},
	}
    robots := make([]*robot, 27)
    robots[0] = newRobot()
    robots[0].loadKeypad(keyPad)
    robots[0].setPosition('A')
    for i := 1; i < len(robots); i++ {
        robots[i] = newRobot()
        robots[i].loadKeypad(directionalKeyPad)
        robots[0].setPosition('A')
    }
	res := 0
	start := time.Now()
	for line := range common.InputLines() {
        paths := []string{line[0]}
        for i, r := range robots {
            fmt.Println(i)
            paths = r.getInputs(paths)
        }
		l := len(paths[0])
		v := common.Net(strconv.Atoi(line[0][:len(line[0])-1]))
		fmt.Println(v, l)
		res += l * v
	}
	fmt.Printf("%v\n", time.Since(start))
	return res
}
func Part2() int {
	// lines := common.GetInputs()
	return 0
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
