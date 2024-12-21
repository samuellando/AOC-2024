package main

import (
	"advent/common"
	"fmt"
    "slices"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type meta struct {
	value string
	state state
}

func Part1() int {
	input := string(common.Input())
	snodeMatrix := common.StringToNodeMatrix(input)
	nodeMatrix := setupMeta(snodeMatrix)
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[*meta]) bool {
		av := a.GetValue()
		bv := b.GetValue()
		if (av.value == "S" || av.value == "." || av.value == "E") && (bv.value == "." || bv.value == "E") {
			return true
		} else {
			return false
		}
	})
	start := getLabeledtNode(nodeMatrix, "S")
	start.GetValue().state.facing = complex(1, 0)
	end := getLabeledtNode(nodeMatrix, "E")
	cost, _ := dijk(start, end)
	return cost
}

func Part2() int {
	input := string(common.Input())
	snodeMatrix := common.StringToNodeMatrix(input)
	nodeMatrix := setupMeta(snodeMatrix)
	common.ConnectAdjs(nodeMatrix, func(a, b common.Node[*meta]) bool {
		av := a.GetValue()
		bv := b.GetValue()
		if (av.value == "S" || av.value == "." || av.value == "E") && (bv.value == "." || bv.value == "E") {
			return true
		} else {
			return false
		}
	})
	start := getLabeledtNode(nodeMatrix, "S")
	start.GetValue().state.facing = complex(1, 0)
	end := getLabeledtNode(nodeMatrix, "E")
	_, paths := dijk(start, end)

    counted := make(map[complex128]bool)
    count := 0
    for _, path := range paths {
        for _, pos := range path {
            if !counted[pos] {
                count++
                counted[pos] = true
            } 
        }
    }
	return count
}

type state struct {
	pos    complex128
	facing complex128
}

type rotation struct {
	r complex128
	c int
}

func dijk(start, end common.Node[*meta]) (int, [][]complex128) {
	distances := make(map[state]int)

	pq := make(map[common.Node[*meta]][][]complex128)

	sv := start.GetValue()
	rotations := []rotation{{complex(1, 0), 1}, {complex(-1, 0), 2001}, {complex(0, 1), 1001}, {complex(0, -1), 1001}}
	for _, r := range rotations {
		n := common.CreateNode(&meta{sv.value, state{sv.state.pos, sv.state.facing * r.r}})
		for _, a := range start.GetAdj() {
			n.Connect(a)
		}
		distances[n.GetValue().state] = r.c - 1
		pq[n] = [][]complex128{{n.GetValue().state.pos}}
	}

	best := -1
    var paths [][]complex128
	for len(pq) > 0 {
		var node common.Node[*meta]
		minDistance := -1
		// TODO use a priority queue
		for n := range pq {
			d, ok := distances[n.GetValue().state]
			if !ok || d < 0 {
				panic("Invalid distance, internal error")
			}
			if d < minDistance || minDistance == -1 {
				node = n
				minDistance = d
			}
		}
		adj := node.GetAdj()
		for _, a := range adj {
			nv := node.GetValue()
			av := a.GetValue()
			// It's only possilbe to connect to a node if we are facing it
			if nv.state.pos+nv.state.facing != av.state.pos {
				continue
			}
			if av.state.pos == end.GetValue().state.pos {
				newD := distances[nv.state] + 1
                es := end.GetValue().state
				if newD < best || best == -1 {
					paths = make([][]complex128, 0)
					for _, p := range pq[node] {
						paths = append(paths, append(slices.Clone(p), es.pos))
					}
					best = newD
				} else if newD == best {
					for _, p := range pq[node] {
						paths = append(paths, append(slices.Clone(p), es.pos))
					}
				}
				continue
			}
			for _, r := range rotations {
				perm := common.CreateNode(&meta{av.value, state{av.state.pos, nv.state.facing * r.r}})
				pv := perm.GetValue()
				d, ok := distances[pv.state]
				newD := distances[nv.state] + r.c
				if !ok || newD <= d {
					for _, b := range a.GetAdj() {
						perm.Connect(b)
					}
					distances[pv.state] = newD
                    pq[perm] = make([][]complex128, 0)
					for _, p := range pq[node] {
						pq[perm] = append(pq[perm], append(slices.Clone(p), pv.state.pos))
					}
				}
			}
		}
		delete(pq, node)
	}
	return best, paths
}

func setupMeta(m [][]common.Node[string]) [][]common.Node[*meta] {
	mm := make([][]common.Node[*meta], len(m))
	for y, row := range m {
		mrow := make([]common.Node[*meta], len(row))
		mm[y] = mrow
		for x, n := range row {
			mm[y][x] = common.CreateNode(&meta{n.GetValue(), state{complex(float64(x), float64(y)), complex(0, 0)}})
		}
	}
	return mm
}

func getLabeledtNode(m [][]common.Node[*meta], c string) common.Node[*meta] {
	for _, row := range m {
		for _, n := range row {
			if n.GetValue().value == c {
				return n
			}
		}
	}
	return nil
}

func display(m [][]common.Node[*meta], d map[state]bool) {
	for y, row := range m {
		for x, n := range row {
			found := false
			for k := range d {
				if k.pos == complex(float64(x), float64(y)) {
					found = true
				}
			}
			if found {
				fmt.Print("*")
			} else {
				fmt.Print(n.GetValue().value)
			}
		}
		fmt.Println()
	}
}
