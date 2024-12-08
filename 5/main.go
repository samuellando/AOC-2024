package main

import (
	"advent/common"
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
	common.CreateRule("root", common.OneOrMoreExpression(common.SeqExpression(
		common.OptionalExpression(
			common.OrExpression(
				common.ReferenceExpression("order"),
				common.ReferenceExpression("list"))),
		common.TerminalExpression("\n")))),
	common.CreateRule("order", common.SeqExpression(
		common.ReferenceExpression("num"),
		common.TerminalExpression("|"),
		common.ReferenceExpression("num"))),
	common.CreateRule("list", common.OrExpression(
		common.SeqExpression(
			common.ReferenceExpression("num"),
			common.OptionalExpression(
				common.SeqExpression(
					common.TerminalExpression(","),
					common.ReferenceExpression("list")))))),
	common.CreateRule("num", common.RegexExpression(`\d+`)))

func Part1() int {
	in := string(common.Input())
	st := G.Parse(in)
	orders := st.Find("order")
	lists := st.Find("list")
	count := 0
	for _, list := range lists {
		numbers := list.Find("num")
        filter := make(map[string]bool)
        for _, n := range(numbers) {
            filter[n.Value()] = true
        }
		levels, err := getLevels(orders, filter)
		if err != nil {
			panic(err)
		}
		level := 0
		good := true
		for _, number := range numbers {
			l, ok := levels[number.Value()]
			if ok {
				if l < level {
					good = false
					break
				}
				level = l
			}
		}
		if good {
			middle := (len(numbers) / 2)
			middleV := common.Net(strconv.Atoi(numbers[middle].Value()))
			count += middleV
		}
	}
	return count
}

func Part2() int {
	in := string(common.Input())
	st := G.Parse(in)
	orders := st.Find("order")
	lists := st.Find("list")
	count := 0
	for _, list := range lists {
		numbers := list.Find("num")
        filter := make(map[string]bool)
        for _, n := range(numbers) {
            filter[n.Value()] = true
        }
		levels, err := getLevels(orders, filter)
		if err != nil {
			panic(err)
		}
		level := 0
		good := true
		for _, number := range numbers {
			l, ok := levels[number.Value()]
			if ok {
				if l < level {
					good = false
					break
				}
				level = l
			}
		}
		if !good {
            orders, err := getOrder(orders, filter)
            if err != nil {
                panic(err)
            }
            c := make([]string, 0)
            for _, level := range(orders) {
                for _, n := range(level) {
                    if filter[n.GetValue()] {
                        c = append(c, n.GetValue())
                    }
                }
            }
			middle := (len(c) / 2)
			middleV := common.Net(strconv.Atoi(c[middle]))
			count += middleV
		}
	}
	return count
}

func getLevels(orders []common.SyntaxTree, filter map[string]bool) (map[string]int, error) {
	g := common.CreateGraph()
	for _, order := range orders {
		numbers := order.Find("num")
		start := numbers[0].Value()
		end := numbers[1].Value()
        if !filter[start] {
            continue
        }
		var s common.Node
		var e common.Node
		if g.GetNode(start) == nil {
			s = common.CreateNode(start)
			g.AddNode(s)
		} else {
			s = g.GetNode(start)
		}
		if g.GetNode(end) == nil {
			e = common.CreateNode(end)
			g.AddNode(e)
		} else {
			e = g.GetNode(end)
		}
		e.Connect(s)
	}
	levels, err := common.TopographicalOrdering(g)
	if err != nil {
		return nil, err
	}
	m := make(map[string]int)
	for i, level := range levels {
		for _, n := range level {
			m[n.GetValue()] = i
		}
	}
	return m, nil
}

func getOrder(orders []common.SyntaxTree, filter map[string]bool) ([][]common.Node, error) {
	g := common.CreateGraph()
	for _, order := range orders {
		numbers := order.Find("num")
		start := numbers[0].Value()
		end := numbers[1].Value()
        if !filter[start] {
            continue
        }
		var s common.Node
		var e common.Node
		if g.GetNode(start) == nil {
			s = common.CreateNode(start)
			g.AddNode(s)
		} else {
			s = g.GetNode(start)
		}
		if g.GetNode(end) == nil {
			e = common.CreateNode(end)
			g.AddNode(e)
		} else {
			e = g.GetNode(end)
		}
		e.Connect(s)
	}
	levels, err := common.TopographicalOrdering(g)
	if err != nil {
		return nil, err
	}
    return levels, nil
}
