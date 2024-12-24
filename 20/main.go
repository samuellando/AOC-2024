package main

import (
	"advent/common"
	"advent/common/datastructures/priorityqueue"
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type node common.Node[string]

type point struct {
	x int
	y int
}

func Part1() int {
	input := string(common.Input())
	nodeMatrix := common.StringToNodeMatrix(input)
	connect(nodeMatrix)
	start := findNode(nodeMatrix, "S")
	distances := dijk(start)
	tracks := findPositions(nodeMatrix, ".", "E", "S")
	count := 0
	for _, p1 := range tracks {
		for _, p2 := range tracks {
			/* Physical distance vs on track distance */
			dist := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			p1d, ok1 := distances[nodeMatrix[p1.y][p1.x]]
			p2d, ok2 := distances[nodeMatrix[p2.y][p2.x]]
			trackDist := p2d - p1d
			save := trackDist - dist
			if ok1 && ok2 && save >= 100 && dist <= 2 {
				count++
			}
		}
	}
	return count
}

func Part2() int {
	input := string(common.Input())
	nodeMatrix := common.StringToNodeMatrix(input)
	connect(nodeMatrix)
	start := findNode(nodeMatrix, "S")
	distances := dijk(start)
	tracks := findPositions(nodeMatrix, ".", "E", "S")
	count := 0
	for _, p1 := range tracks {
		for _, p2 := range tracks {
			/* Physical distance vs on track distance */
			dist := abs(p1.x-p2.x) + abs(p1.y-p2.y)
			p1d, ok1 := distances[nodeMatrix[p1.y][p1.x]]
			p2d, ok2 := distances[nodeMatrix[p2.y][p2.x]]
			trackDist := p2d - p1d
			save := trackDist - dist
			if ok1 && ok2 && save >= 100 && dist <= 20 {
				count++
			}
		}
	}
	return count
}

func connect(m [][]common.Node[string]) {
	common.ConnectAdjs(m, func(a, b common.Node[string]) bool {
		av := a.GetValue()
		bv := b.GetValue()
		return (av == "." && bv == ".") || (av == "." && bv == "E") || (av == "S" && bv == ".")
	})
}

func wrap(n node) *priorityqueue.Item[node] {
	return &priorityqueue.Item[node]{Value: n, Priority: 0, Index: 0}
}

func dijk(start node) map[node]int {
	visited := make(map[node]bool)
	distances := make(map[node]int)
	pq := priorityqueue.New[node]()
	heap.Push(&pq, wrap(start))
    distances[start] = 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*priorityqueue.Item[node])
		node := item.Value
		dist := distances[node]
		for _, adj := range node.GetAdj() {
			if _, ok := visited[adj]; !ok {
				newD := dist + 1
				adist, ok := distances[adj]
				if newD < adist || !ok {
					distances[adj] = newD
					heap.Push(&pq, wrap(adj))
				}
			}
		}
		visited[node] = true
	}
	return distances
}

func display(m [][]common.Node[string]) {
	for _, row := range m {
		for _, c := range row {
			fmt.Print(c.GetValue())
		}
		fmt.Println()
	}
}

func findNode(m [][]common.Node[string], v string) node {
	for _, row := range m {
		for _, n := range row {
			if n.GetValue() == v {
				return n
			}
		}
	}
	return nil
}

func findPositions(m [][]common.Node[string], vs ...string) []point {
	l := make([]point, 0)
	for y, row := range m {
		for x, n := range row {
			if slices.Contains(vs, n.GetValue()) {
				l = append(l, point{x, y})
			}
		}
	}
	return l
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
