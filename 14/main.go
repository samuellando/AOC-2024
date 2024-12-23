package main

import (
	"advent/common"
	"fmt"
	"strconv"
    "strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

var G = common.CreateGrammar(
	common.CreateRule("root", common.OneOrMoreExpression(
		common.ReferenceExpression("robot"))),
	common.CreateRule("robot",
		common.SeqExpression(
			common.TerminalExpression("p="),
			common.ReferenceExpression("point"),
			common.TerminalExpression(" v="),
			common.ReferenceExpression("point"),
			common.TerminalExpression("\n"))),
	common.CreateRule("point", common.SeqExpression(
		common.ReferenceExpression("num"),
		common.TerminalExpression(","),
		common.ReferenceExpression("num"))),
	common.CreateRule("num",
		common.RegexExpression(`(-*\d+)`)))

type robot struct {
	p vec
	v vec
}

type vec struct {
	x int
	y int
}

func loadRobot(config common.SyntaxTree) *robot {
	points := config.Find("point")
	pos := points[0].Find("num")
	vel := points[1].Find("num")
	toInt := func(x common.SyntaxTree) int {
		return common.Net(strconv.Atoi(x.Value()))
	}
	x := toInt(pos[0])
	y := toInt(pos[1])
	vx := toInt(vel[0])
	vy := toInt(vel[1])
	return &robot{vec{x, y}, vec{vx, vy}}
}

func Part1() int {
	input := string(common.Input())
	st,_ := G.Parse(input)
	robotConfigs := st.Find("robot")
	w := 101
	h := 103
	counts := []int{0, 0, 0, 0}

	for _, config := range robotConfigs {
		robot := loadRobot(config)
		for range 100 {
			robot.move(w, h)
		}
		q := getQuadrant(robot, w, h)
		if q >= 0 {
			counts[q]++
		}
	}
	res := 1
	for _, c := range counts {
		res *= c
	}
	return res
}

func Part2() int {
	input := string(common.Input())
	st,_ := G.Parse(input)
	robotConfigs := st.Find("robot")
	w := 101
	h := 103

	board := make([][]int, h)
	for i := range h {
		board[i] = make([]int, w)
	}

	robots := make([]*robot, len(robotConfigs))
	for i, config := range robotConfigs {
		robots[i] = loadRobot(config)
	}

	time := 0
	for {
		for _, robot := range robots {
			if board[robot.p.y][robot.p.x] != 0 {

				board[robot.p.y][robot.p.x]--
			}
			robot.move(w, h)
			board[robot.p.y][robot.p.x]++
		}
		time += 1
        found := false
		for _, row := range board {
            s := ""
			for _, i := range row {
                s += strconv.Itoa(i)
			}
            if strings.Contains(s, "11111111111111111111111111111") {
                found = true
            }
		}
		if found {
            break
        }
	}
	return time
}

func (r *robot) move(w, h int) {
	r.p.x += r.v.x
	r.p.y += r.v.y

	if r.p.x < 0 {
		r.p.x = w + r.p.x
	}
	if r.p.y < 0 {
		r.p.y = h + r.p.y
	}
	if r.p.x >= w {
		r.p.x = r.p.x - w
	}
	if r.p.y >= h {
		r.p.y = r.p.y - h
	}
}

func getQuadrant(r *robot, w, h int) int {
	qw := w / 2 // 5
	qh := h / 2 // 3
	xq := -1
	yq := -1
	if r.p.x < qw {
		xq = 0
	}
	if r.p.y < qh {
		yq = 0
	}
	if r.p.x >= w-qw {
		xq = 1
	}
	if r.p.y >= h-qh {
		yq = 1
	}
	var res int
	if xq == -1 || yq == -1 {
		res = -1
	} else {
		res = xq + 2*yq
	}
	return res
}
