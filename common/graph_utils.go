package common

import (
	"errors"
	"fmt"
	"strings"
    "advent/common/datastructures/queue"
)

func StringToNodeMatrix(input string) [][]Node[string] {
	lines := strings.Split(input, "\n")
	nodeMatrix := make([][]Node[string], 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		row := make([]Node[string], 0)
		for _, c := range line {
			row = append(row, CreateIndexedAdjNode(string(c)))
		}
		nodeMatrix = append(nodeMatrix, row)
	}
	return nodeMatrix
}

func ConnectAdjs[T any](nodeMatrix [][]Node[T], condition func(a, b Node[T]) bool) {
	for i, row := range nodeMatrix {
		for j, node := range row {
			var a Node[T]
			if i > 0 {
				a = nodeMatrix[i-1][j]
				if condition(node, a) {
					node.Connect(a)
				}
			}
			if i < len(nodeMatrix)-1 {
				a = nodeMatrix[i+1][j]
				if condition(node, a) {
					node.Connect(a)
				}
			}
			if j > 0 {
				a = nodeMatrix[i][j-1]
				if condition(node, a) {
					node.Connect(a)
				}
			}
			if j < len(row)-1 {
				a = nodeMatrix[i][j+1]
				if condition(node, a) {
					node.Connect(a)
				}
			}
		}
	}
}

func Bfs[T any](n Node[T], callback func(n Node[T]) bool) {
    visited := make(map[Node[T]]bool) 
	q := queue.New[Node[T]]()
	q.Enqueue(n)
    if !callback(n) {
        return
    }
	for q.Length() > 0 {
		node := q.Dequeue()
		adj := node.GetAdj()
		for _, a := range adj {
            if _, ok := visited[a]; !ok {
                visited[a] = true
                if callback(a) {
                    q.Enqueue(a)
                }
            }
		}
	}
}

func TopographicalOrdering[T any](g Graph[T]) ([][]Node[T], error) {
	loaded := make(map[string]bool)
	order := make([][]Node[T], 0)
	done := true
	ls := 0
	for {
		level := make([]Node[T], 0)
		for _, node := range g.GetNodes() {
			str := fmt.Sprint(node.GetValue())
			if loaded[str] {
				continue
			} else {
				done = false
			}
			met := true
			for _, adj := range node.GetAdj() {
				str := fmt.Sprint(adj.GetValue())
				if !loaded[str] {
					met = false
					break
				}
			}
			if met {
				level = append(level, node)
			}
		}
		if len(level) == 0 && !done {
			return nil, errors.New("Not DAG")
		}
		order = append(order, level)
		for _, n := range level {
			str := fmt.Sprint(n.GetValue())
			loaded[str] = true
			ls += 1
		}
		if done {
			break
		}
		done = true
	}
	return order, nil
}
