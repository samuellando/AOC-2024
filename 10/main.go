package main

import (
	"advent/common"
	"advent/common/datastructures/queue"
	"fmt"
	"iter"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	nodeMatrix := toNodeMatrix(common.InputLines())
	connectAdjs(nodeMatrix)
	starts := getStartNodes(nodeMatrix)
	// Use BFS to search from each start position
	total := 0
	for _, startNode := range starts {
		visited := make(map[common.Node]bool)
		q := queue.New[common.Node]()
		q.Enqueue(startNode)
		score := 0
		for q.Length() > 0 {
			node := q.Dequeue()
			v := common.Net(strconv.Atoi(node.GetValue()))
			adj := node.GetAdj()
			for _, a := range adj {
				if _, ok := visited[a]; !ok {
					av := common.Net(strconv.Atoi(a.GetValue()))
					if av-v == 1 {
						if av == 9 {
							score++
						}
						q.Enqueue(a)
						visited[a] = true
					}
				}
			}
		}
		total += score
	}
	return total
}

func Part2() int {
	nodeMatrix := toNodeMatrix(common.InputLines())
	connectAdjs(nodeMatrix)
	starts := getStartNodes(nodeMatrix)
	// Use DFS to search from each start position
	total := 0
	for _, startNode := range starts {
		visited := make(map[common.Node]bool)
		score := dfs(startNode, visited)
		total += score
	}
	return total
}

func dfs(n common.Node, visited map[common.Node]bool) int {
	v := common.Net(strconv.Atoi(n.GetValue()))
    if v == 9 {
        return 1
    }
	count := 0
	thisVisited := make(map[common.Node]bool)
	for k, v := range visited {
		thisVisited[k] = v
	}
	thisVisited[n] = true
	for _, a := range n.GetAdj() {
        av := common.Net(strconv.Atoi(a.GetValue()))
        if av - v == 1 {
            count += dfs(a, thisVisited)
        }
	}
	return count
}

func toNodeMatrix(lines iter.Seq[[]string]) [][]common.Node {
	nodeMatrix := make([][]common.Node, 0)
	for line := range lines {
		row := make([]common.Node, 0)
		for _, c := range line[0] {
			row = append(row, common.CreateIndexedAdjNode(string(c)))
		}
		nodeMatrix = append(nodeMatrix, row)
	}
	return nodeMatrix
}

func connectAdjs(nodeMatrix [][]common.Node) {
	for i, row := range nodeMatrix {
		for j, node := range row {
			if i > 0 {
				node.Connect(nodeMatrix[i-1][j])
			}
			if i < len(nodeMatrix)-1 {
				node.Connect(nodeMatrix[i+1][j])
			}
			if j > 0 {
				node.Connect(nodeMatrix[i][j-1])
			}
			if j < len(row)-1 {
				node.Connect(nodeMatrix[i][j+1])
			}
		}
	}
}

func getStartNodes(nodeMatrix [][]common.Node) []common.Node {
	starts := make([]common.Node, 0)
	for _, row := range nodeMatrix {
		for _, node := range row {
			if node.GetValue() == "0" {
				starts = append(starts, node)
			}
		}
	}
	return starts
}
