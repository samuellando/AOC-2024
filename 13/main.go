package main

import (
	"advent/common"
	"errors"
	"fmt"
	"math"
	"strconv"
)

var G = common.CreateGrammar(
	common.CreateRule("root",
		common.OneOrMoreExpression(
			common.OrExpression(
				common.ReferenceExpression("machine"),
				common.TerminalExpression("\n")))),
	common.CreateRule("machine",
		common.SeqExpression(
			common.OneOrMoreExpression(
				common.ReferenceExpression("buttonDef")),
			common.ReferenceExpression("prizeDef"))),
	common.CreateRule("buttonDef",
		common.SeqExpression(
			common.TerminalExpression("Button "),
			common.ReferenceExpression("buttonId"),
			common.TerminalExpression(": "),
			common.ReferenceExpression("movements"),
			common.TerminalExpression("\n"))),
	common.CreateRule("buttonId", common.RegexExpression(`(A|B)`)),
	common.CreateRule("movements",
		common.OneOrMoreExpression(
			common.OrExpression(
				common.SeqExpression(
					common.ReferenceExpression("movement"),
					common.TerminalExpression(", "),
					common.ReferenceExpression("movements")),
				common.ReferenceExpression("movement")))),
	common.CreateRule("movement",
		common.SeqExpression(
			common.ReferenceExpression("axis"),
			common.ReferenceExpression("value"))),
	common.CreateRule("value",
		common.RegexExpression(`((\+|-)*\d+)`)),
	common.CreateRule("axis",
		common.RegexExpression(`(X|Y)`)),
	common.CreateRule("prizeDef",
		common.SeqExpression(
			common.TerminalExpression("Prize: "),
			common.ReferenceExpression("position"),
			common.TerminalExpression("\n"))),
	common.CreateRule("position",
		common.SeqExpression(
			common.TerminalExpression("X="),
			common.ReferenceExpression("value"),
			common.TerminalExpression(", Y="),
			common.ReferenceExpression("value"))))

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

type button struct {
	dx int
	dy int
}

func loadButton(st common.SyntaxTree) *button {
	svalues := st.Find("value")
	dx := common.Net(strconv.Atoi(svalues[0].Value()))
	dy := common.Net(strconv.Atoi(svalues[1].Value()))
	return &button{dx, dy}
}

type position struct {
	x int
	y int
}

func loadPosition(st common.SyntaxTree) *position {
	svalues := st.Find("value")
	x := common.Net(strconv.Atoi(svalues[0].Value()))
	y := common.Net(strconv.Atoi(svalues[1].Value()))
	return &position{x, y}
}

type machine struct {
	a     *button
	b     *button
	prize *position
}

func loadMachine(st common.SyntaxTree) *machine {
	buttonConfigs := st.Find("buttonDef")
	prizeConfig := st.Find("prizeDef")[0]
	return &machine{loadButton(buttonConfigs[0]), loadButton(buttonConfigs[1]), loadPosition(prizeConfig)}
}

func Part1() int {
	input := string(common.Input())
	st := G.Parse(input)
	machineConfigs := st.Find("machine")
	cost := 0
	for _, config := range machineConfigs {
		machine := loadMachine(config)
		a, b, err := solve(machine.a.dx, machine.a.dy, machine.b.dx, machine.b.dy, machine.prize.x, machine.prize.y)
		if err == nil {
			aIsValid := a >= 0 && a <= 100
			bIsValid := b >= 0 && b <= 100
			if aIsValid && bIsValid {
				cost += 3*a + b
			}
		}
	}
	return cost
}

func Part2() int {
	input := string(common.Input())
	st := G.Parse(input)
	machineConfigs := st.Find("machine")
	cost := 0
    offset := 10000000000000
	for _, config := range machineConfigs {
		machine := loadMachine(config)
		a, b, err := solve(machine.a.dx, machine.a.dy, machine.b.dx, machine.b.dy, machine.prize.x + offset, machine.prize.y + offset)
		if err == nil {
			aIsValid := a >= 0 
			bIsValid := b >= 0
			if aIsValid && bIsValid {
				cost += 3*a + b
			}
		}
	}
	return cost
}

func solve(adxi, adyi, bdxi, bdyi, pxi, pyi int) (int, int, error) {
	adx := float64(adxi)
	ady := float64(adyi)
	bdx := float64(bdxi)
	bdy := float64(bdyi)
	px := float64(pxi)
	py := float64(pyi)

	a1 := adx / bdx
	a2 := ady / bdy
	if math.Abs(a1-a2) <= 0.0001 {
		// Colinear => more than one solution. Check the cost.
		fmt.Println("COLINEAR")
	} else {
		b := (py - (ady * px / adx)) / (bdy - (ady * bdx / adx))
		a := (px - bdx*b) / adx
		aIsInt := a-math.Floor(a) < 0.01 || math.Ceil(a)-a < 0.01
		bIsInt := b-math.Floor(b) < 0.01 || math.Ceil(b)-b < 0.01
		if bIsInt && aIsInt {
			return int(math.Round(a)), int(math.Round(b)), nil
		}
	}
	return 0, 0, errors.New("No int solution")
}
