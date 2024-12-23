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
	common.CreateRule("root", common.ReferenceExpression("patterns")),
	common.CreateRule("patterns", common.OneOrMoreExpression(
		common.OrExpression(
			common.SeqExpression(
				common.ReferenceExpression("pattern"),
				common.TerminalExpression(", "),
				common.ReferenceExpression("patterns")),
			common.ReferenceExpression("pattern")))),
	common.CreateRule("pattern", common.RegexExpression(`([^,]*)`)))

func Part1() int {
	input := string(common.Input())
	lines := strings.Split(input, "\n")
	st, err := G.Parse(lines[0])
	if err != nil {
		panic(err)
	}
	patternsSt := st.Find("pattern")
	patternRules := make([]common.Rule, len(patternsSt))
	referenceExpressions := make([]common.Expression, len(patternsSt))
	for i, pst := range patternsSt {
		patternRules[i] = common.CreateRule(pst.Value(), common.SeqExpression(
				common.TerminalExpression(pst.Value()),
				common.OptionalExpression(common.ReferenceExpression("pat"))))
		referenceExpressions[i] = common.ReferenceExpression(pst.Value())
	}
	patternGrammar := common.CreateGrammar(
		common.CreateRule("pat", common.GreedyOrExpression(
			referenceExpressions...)),
		patternRules...)
	count := 0
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
		_, err = patternGrammar.Parse(line)
		if err == nil {
			count++
		}
	}
	return count
}

func Part2() int {
	input := string(common.Input())
	lines := strings.Split(input, "\n")
	st, err := G.Parse(lines[0])
	if err != nil {
		panic(err)
	}
	patternsSt := st.Find("pattern")
	patternRules := make([]common.Rule, len(patternsSt))
	referenceExpressions := make([]common.Expression, len(patternsSt))
	for i, pst := range patternsSt {
		patternRules[i] = common.CreateRule(pst.Value(), common.SeqExpression(
				common.TerminalExpression(pst.Value()),
				common.OptionalExpression(common.ReferenceExpression("pat"))))
		referenceExpressions[i] = common.ReferenceExpression(pst.Value())
	}
	patternGrammar := common.CreateGrammar(
		common.CreateRule("pat", common.GreedyOrExpressionCounter(
			referenceExpressions...)),
		patternRules...)
	count := 0
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}
        t, err := patternGrammar.Parse(line)
		if err == nil {
            count += common.Net(strconv.Atoi(t.Value()))
		}
	}
	return count
}
