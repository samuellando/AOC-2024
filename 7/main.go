package main

import (
	"advent/common"
	"fmt"
	"slices"
	"strconv"
)

var G = common.CreateGrammar(
	common.CreateRule("root", common.OneOrMoreExpression(common.SeqExpression(
		common.ReferenceExpression("equation"),
		common.TerminalExpression("\n")))),
	common.CreateRule("equation", common.SeqExpression(
		common.ReferenceExpression("solution"),
		common.TerminalExpression(":"),
		common.ReferenceExpression("nums"))),
	common.CreateRule("nums", common.SeqExpression(
		common.TerminalExpression(" "),
		common.ReferenceExpression("num"),
		common.OptionalExpression(
			common.ReferenceExpression("nums")))),
	common.CreateRule("solution", common.ReferenceExpression("num")),
	common.CreateRule("num", common.RegexExpression(`\d+`)))

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
    ops := []func(int, int) int {
        func(a, b int) int {return a + b}, 
        func(a, b int) int {return a * b}, 
    }
	input := string(common.Input())
	st := G.Parse(input)
    sum := 0
	for _, eq := range st.Find("equation") {
		solution := common.Net(strconv.Atoi(eq.Find("solution")[0].Value()))
		nums := common.Map(common.ToSeq(eq.Find("nums")[0].Find("num")), func(v common.SyntaxTree) int { return common.Net(strconv.Atoi(v.Value())) })
        if findOps(solution, ops,  slices.Collect(nums)...) {
            sum += solution
        }

	}
	return sum
}

func findOps(solution int, ops []func(int, int) int, nums ...int) bool {
    // Base conditions
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if solution == nums[0] {
			return true
		} else {
            return false
        }
	}
    // recursive condition
    for _, op := range(ops) {
        new := slices.Concat([]int{op(nums[0],nums[1])}, nums[2:])
        if findOps(solution, ops, new...) {
            return true
        }
    }
    return false
}

func Part2() int {
    ops := []func(int, int) int {
        func(a, b int) int {return a + b}, 
        func(a, b int) int {return a * b}, 
        func(a, b int) int {
            return common.Net(strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b)))
        }, 
    }
	input := string(common.Input())
	st := G.Parse(input)
    sum := 0
	for _, eq := range st.Find("equation") {
		solution := common.Net(strconv.Atoi(eq.Find("solution")[0].Value()))
		nums := common.Map(common.ToSeq(eq.Find("nums")[0].Find("num")), func(v common.SyntaxTree) int { return common.Net(strconv.Atoi(v.Value())) })
        if findOps(solution, ops,  slices.Collect(nums)...) {
            sum += solution
        }

	}
	return sum
}
