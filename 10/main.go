package main

import (
	"advent/common"
	"advent/common/datastructures/queue"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	nodeMatrix := common.StringToNodeMatrix(string(common.Input()))
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[string]) bool {return true})
	starts := getStartNodes(nodeMatrix)
	// Use BFS to search from each start position
	total := 0
	for _, startNode := range starts {
		visited := make(map[common.Node[string]]bool)
		q := queue.New[common.Node[string]]()
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
	nodeMatrix := common.StringToNodeMatrix(string(common.Input()))
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[string]) bool {return true})
	starts := getStartNodes(nodeMatrix)
	// Use DFS to search from each start position
	total := 0
	for _, startNode := range starts {
		visited := make(map[common.Node[string]]bool)
		score := dfs(startNode, visited)
		total += score
	}
	return total
}

func dfs(n common.Node[string], visited map[common.Node[string]]bool) int {
	v := common.Net(strconv.Atoi(n.GetValue()))
    if v == 9 {
        return 1
    }
	count := 0
	thisVisited := make(map[common.Node[string]]bool)
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

func getStartNodes(nodeMatrix [][]common.Node[string]) []common.Node[string] {
	starts := make([]common.Node[string], 0)
	for _, row := range nodeMatrix {
		for _, node := range row {
			if node.GetValue() == "0" {
				starts = append(starts, node)
			}
		}
	}
	return starts
}
