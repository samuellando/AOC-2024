package main

import (
	"advent/common"
	"fmt"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type valueMeta struct {
	value    string
	visited  bool
	pos      point
	regionId int
}

type point struct {
	row int
	col int
}

func Part1() int {
	input := string(common.Input())
	snodeMatrix := common.StringToNodeMatrix(input)
	nodeMatrix := setupMetaData(snodeMatrix)
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[*valueMeta]) bool {
		return a.GetValue().value == b.GetValue().value
	})
	cost := 0
	for _, row := range nodeMatrix {
		for _, n := range row {
			if !n.GetValue().visited {
				area := 0
				perim := 0
				common.Bfs(n, func(a common.Node[*valueMeta]) bool {
					if !a.GetValue().visited {
						a.GetValue().visited = true
						area += 1
						perim += 4 - len(a.GetAdj())
						return true
					} else {
						return false
					}
				})
				cost += area * perim

			}
		}
	}
	return cost
}

func Part2() int {
	input := string(common.Input())
	snodeMatrix := common.StringToNodeMatrix(input)
	nodeMatrix := setupMetaData(snodeMatrix)
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[*valueMeta]) bool {
		return a.GetValue().value == b.GetValue().value
	})
	// Calc area and find all the corners
	areas := make([]int, 0)
	corners := make([]point, 0)
	for _, row := range nodeMatrix {
		for _, n := range row {
			if !n.GetValue().visited {
				area := 0
				common.Bfs(n, func(a common.Node[*valueMeta]) bool {
					if !a.GetValue().visited {
						a.GetValue().visited = true
						area += 1
						if len(a.GetAdj()) <= 2 {
							corners = append(corners, a.GetValue().pos)
						}
						a.GetValue().regionId = len(areas)
						return true
					} else {
						return false
					}
				})
				areas = append(areas, area)
			}
		}
	}
	// Calc sides based on corners
	sides := make([]int, len(areas))
	for _, p := range corners {
		node := nodeMatrix[p.row][p.col]
		adjn := len(node.GetAdj())
		if adjn == 0 {
			sides[node.GetValue().regionId] += 4
		} else if adjn == 1 {
			sides[node.GetValue().regionId] += 2
		}
		for _, r := range getInteriorCornerRegions(nodeMatrix, p) {
			sides[r] += 1
		}
	}
	// Calc total
	cost := 0
	for i, a := range areas {
		cost += a * sides[i]
	}
	return cost
}

const red = "\033[0;31m"
const none = "\033[0m"

func display(m [][]common.Node[*valueMeta], p point) {
	for i, row := range m {
		for j, n := range row {
			if i == p.row && j == p.col {
				fmt.Print(red, n.GetValue().value, none)
			} else {
				fmt.Print(n.GetValue().value)
			}
		}
        fmt.Println()
	}
}

func getInteriorCornerRegions(m [][]common.Node[*valueMeta], p point) []int {
    hostRegion := m[p.row][p.col].GetValue().regionId
	res := make([]int, 0)
	u := -1
	if p.row > 0 {
		u = m[p.row-1][p.col].GetValue().regionId
	}
	d := -1
	if p.row < len(m)-1 {
		d = m[p.row+1][p.col].GetValue().regionId
	}
	l := -1
	if p.col > 0 {
		l = m[p.row][p.col-1].GetValue().regionId
	}
	r := -1
	if p.col < len(m[p.row])-1 {
		r = m[p.row][p.col+1].GetValue().regionId
	}
	if u != -1 && u == r {
		diag := m[p.row-1][p.col+1].GetValue().regionId
		if u == diag || u == hostRegion {
			res = append(res, u)
		}
	}
	if u != -1 && u == l {
		diag := m[p.row-1][p.col-1].GetValue().regionId
		if u == diag || u == hostRegion {
			res = append(res, u)
		}
	}
	if d != -1 && d == r {
		diag := m[p.row+1][p.col+1].GetValue().regionId
		if d == diag || d == hostRegion {
			res = append(res, d)
		}
	}
	if d != -1 && d == l {
		diag := m[p.row+1][p.col-1].GetValue().regionId
		if d == diag || d == hostRegion {
			res = append(res, d)
		}
	}
	return res
}

func setupMetaData(i [][]common.Node[string]) [][]common.Node[*valueMeta] {
	o := make([][]common.Node[*valueMeta], len(i))
	for i, row := range i {
		or := make([]common.Node[*valueMeta], len(row))
		for j, n := range row {
			or[j] = common.CreateIndexedAdjNode(&valueMeta{n.GetValue(), false, point{i, j}, -1})
		}
		o[i] = or
	}
	return o
}
