package main

import (
	"advent/common"
	"fmt"
	"strconv"
)

var G = common.CreateGrammar(
	common.CreateRule("root",
		common.OneOrMoreExpression(
			common.OrExpression(
				common.ReferenceExpression("function"),
				common.ReferenceExpression("curuption")))),
	common.CreateRule("curuption",
		common.RegexExpression(`.|\n`)),
	common.CreateRule("function",
		common.SeqExpression(
			common.ReferenceExpression("identifier"),
			common.TerminalExpression("("),
			common.ReferenceExpression("args"),
			common.TerminalExpression(")"))),
	common.CreateRule("identifier", common.RegexExpression(`(don't|do|mul)`)),
	common.CreateRule("args",
		common.OptionalExpression(common.ReferenceExpression("oargs"))),
	common.CreateRule("oargs",
		common.OneOrMoreExpression(
			common.OrExpression(
				common.SeqExpression(
					common.ReferenceExpression("arg"),
					common.TerminalExpression(","),
					common.ReferenceExpression("oargs")),
				common.ReferenceExpression("arg")))),
	common.CreateRule("arg", common.RegexExpression(`\d*`)))

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	input := common.Input()
	t := G.Parse(string(input))
	fs := t.Filter("function")
	sum := 0
	for _, f := range fs {
		id := f.Filter("identifier")[0].Value()
		if id == "mul" {
            sum += evalMul(f)
		}
	}
	return sum
}

func Part2() int {
	common.Input()
	input := common.Input()
	t := G.Parse(string(input))
	fs := t.Filter("function")
	sum := 0
    do := true
	for _, f := range fs {
		id := f.Filter("identifier")[0].Value()
		if id == "mul" {
            if do {
                sum += evalMul(f)
            }
		} else if id == "do" && evalDo(f) {
            do = true
		} else if id == "don't" && evalDont(f) {
            do = false
		}
	}
	return sum
}

func evalMul(f common.SyntaxTree) int {
	args := f.Find("args")
	if len(args) > 0 {
		vs := args[0].Find("arg")
		if len(vs) == 2 {
			r := common.Net(strconv.Atoi(vs[0].Value()))
			l := common.Net(strconv.Atoi(vs[1].Value()))
			return r * l
		}
	}
	return 0
}

func evalDo(f common.SyntaxTree) bool {
	args := f.Find("args")
	return len(args) == 0 
}

func evalDont(f common.SyntaxTree) bool {
	args := f.Find("args")
	return len(args) == 0 
}
