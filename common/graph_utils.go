package common

import (
    "errors"
)

func TopographicalOrdering(g Graph) ([][]Node, error) {
	loaded := make(map[string]bool)
	order := make([][]Node, 0)
	done := true
	ls := 0
	for {
		level := make([]Node, 0)
		for _, node := range g.GetNodes() {
			if loaded[node.GetValue()] {
				continue
			} else {
				done = false
			}
			met := true
			for _, adj := range node.GetAdj() {
				if !loaded[adj.GetValue()] {
					met = false
					break
				}
			}
			if met {
				level = append(level, node)
			}
		}
        if len(level) == 0  && !done {
            return nil, errors.New("Not DAG")
        }
		order = append(order, level)
		for _, n := range level {
			loaded[n.GetValue()] = true
			ls += 1
		}
		if done {
			break
		}
		done = true
	}
	return order, nil
}
