package main

import (
	"advent/common"
	"advent/common/datastructures/priorityqueue"
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

var G = common.CreateGrammar(
	common.CreateRule("root", common.OneOrMoreExpression(
		common.SeqExpression(
			common.ReferenceExpression("cord"),
			common.TerminalExpression("\n")))),
	common.CreateRule("cord",
		common.SeqExpression(
			common.ReferenceExpression("num"),
			common.TerminalExpression(","),
			common.ReferenceExpression("num"))),
	common.CreateRule("num", common.RegexExpression(`\d+`)))

func loadCords(st common.SyntaxTree) []point {
	cords := st.Find("cord")
	points := make([]point, len(cords))
	for i, c := range cords {
		nums := c.Find("num")
		x := common.Net(strconv.Atoi(nums[0].Value()))
		y := common.Net(strconv.Atoi(nums[1].Value()))
		points[i] = point{x, y}
	}
	return points
}

type point struct {
	x int
	y int
}

type meta struct {
	val string
}

type board [][]common.Node[*meta]

func initializeBoard(rows int, cols int) board {
	board := make(board, rows)
	for i := range rows {
		board[i] = make([]common.Node[*meta], cols)
		for j := range cols {
			board[i][j] = common.CreateIndexedAdjNode(&meta{"."})
		}
	}
	return board
}

func (b board) place(s string, p point) {
	b[p.y][p.x].GetValue().val = s
}

func (b board) display() {
	for _, row := range b {
		for _, n := range row {
			fmt.Print(n.GetValue().val)
		}
		fmt.Println()
	}
}

func Part1() int {
	input := string(common.Input())
	board := initializeBoard(71, 71)
	st, _ := G.Parse(input)
	cords := loadCords(st)
	for _, c := range cords[:1024] {
		board.place("#", c)
	}
	common.ConnectAdjs(board, func(a, b common.Node[*meta]) bool {
		if a.GetValue().val == "." && b.GetValue().val == "." {
			return true
		} else {
			return false
		}
	})
	start := board[0][0]
	end := board[70][70]
	d := dijk(start, end)
	return d
}

func Part2() string {
	input := string(common.Input())
	st, _ := G.Parse(input)
	cords := loadCords(st)
	var sol point
	for drops := 1024; drops <= len(cords); drops++ {
		board := initializeBoard(71, 71)
		for _, c := range cords[:drops] {
			board.place("#", c)
		}
		common.ConnectAdjs(board, func(a, b common.Node[*meta]) bool {
			if a.GetValue().val == "." && b.GetValue().val == "." {
				return true
			} else {
				return false
			}
		})
		start := board[0][0]
		end := board[70][70]
		d := dijk(start, end)
		if d == -1 {
			sol = cords[drops-1]
			break
		}
	}
	return fmt.Sprintf("%d,%d", sol.x, sol.y)
}

func wrap[T any](n T, p int) *priorityqueue.Item[T] {
	return &priorityqueue.Item[T]{Value: n, Priority: p, Index: 0}
}

func dijk(start, end common.Node[*meta]) int {
	distances := make(map[common.Node[*meta]]int)
	visited := make(map[common.Node[*meta]]bool)
	pqItems := make(map[common.Node[*meta]]*priorityqueue.Item[common.Node[*meta]])
	pq := priorityqueue.New[common.Node[*meta]]()
	pqItems[start] = wrap(start, 0)
	heap.Push(&pq, pqItems[start])
	distances[start] = 0
	for pq.Len() > 0 {
		topItem := heap.Pop(&pq).(*priorityqueue.Item[common.Node[*meta]])
		node := topItem.Value
		for _, adj := range node.GetAdj() {
			if _, ok := visited[adj]; !ok {
				newD := distances[node] + 1
				if d, ok := distances[adj]; !ok || newD < d {
					distances[adj] = newD
					if item, ok := pqItems[adj]; !ok {
						pqItems[adj] = wrap(adj, newD)
						heap.Push(&pq, pqItems[adj])
					} else {
						pq.Update(item, adj, newD)
					}
				}
			}
		}
		visited[node] = true
	}
	if d, ok := distances[end]; ok {
		return d
	} else {
		return -1
	}
}
