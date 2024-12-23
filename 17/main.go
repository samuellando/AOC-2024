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
	common.CreateRule("root", common.SeqExpression(
		common.ReferenceExpression("register"),
		common.ReferenceExpression("register"),
		common.ReferenceExpression("register"),
		common.TerminalExpression("\n"),
		common.ReferenceExpression("program"))),
	common.CreateRule("register", common.SeqExpression(
		common.RegexExpression("Register (A|B|C): "),
		common.ReferenceExpression("num"),
		common.TerminalExpression("\n"))),
	common.CreateRule("program", common.SeqExpression(
		common.TerminalExpression("Program: "),
		common.ReferenceExpression("instructions"),
		common.TerminalExpression("\n"))),
	common.CreateRule("instructions", common.OneOrMoreExpression(
		common.OrExpression(
			common.SeqExpression(
				common.ReferenceExpression("instruction"),
				common.TerminalExpression(","),
				common.ReferenceExpression("instructions")),
			common.ReferenceExpression("instruction")))),
	common.CreateRule("instruction", common.SeqExpression(
		common.ReferenceExpression("num"),
		common.TerminalExpression(","),
		common.ReferenceExpression("num"))),
	common.CreateRule("num", common.RegexExpression(`\d+`)))

type computer struct {
	pc           int
	a            int
	b            int
	c            int
	instructions []instruction
}

type instruction struct {
	opcode  int
	operand int
}

func loadComputer(config common.SyntaxTree) computer {
	c := computer{0, 0, 0, 0, make([]instruction, 0)}
	registers := config.Find("register")
	instructions := config.Find("instruction")
	c.a = common.Net(strconv.Atoi(registers[0].Find("num")[0].Value()))
	c.b = common.Net(strconv.Atoi(registers[1].Find("num")[0].Value()))
	c.c = common.Net(strconv.Atoi(registers[2].Find("num")[0].Value()))

	for _, ic := range instructions {
		c.instructions = append(c.instructions, loadInstruction(ic))
	}
	return c
}

func loadInstruction(config common.SyntaxTree) instruction {
	nums := config.Find("num")
	code := common.Net(strconv.Atoi(nums[0].Value()))
	op := common.Net(strconv.Atoi(nums[1].Value()))
	return instruction{code, op}
}

func (c *computer) comboOperand(op int) int {
	switch op {
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		return op
	}
}

func (c *computer) execute() []int {
	output := make([]int, 0)
	for c.pc < len(c.instructions) {
		inst := c.instructions[c.pc]
		switch inst.opcode {
		case 0:
			// adv: division
			numerator := c.a
			denominator := 1 << c.comboOperand(inst.operand)
			c.a = numerator / denominator
		case 1:
			// bxl: xor with operand
			c.b = c.b ^ inst.operand
		case 2:
			// bst: modulo
			c.b = c.comboOperand(inst.operand) % 8
		case 3:
			// jnz: jump not zero
			if c.a != 0 {
				c.pc = inst.operand - 1
			}
		case 4:
			// bxc: xor with c
			c.b = c.b ^ c.c
		case 5:
			// out
			output = append(output, c.comboOperand(inst.operand)%8)
		case 6:
			// bdv: division to b
			numerator := c.a
			denominator := 1 << c.comboOperand(inst.operand)
			c.b = numerator / denominator
		case 7:
			// cdv: division to c
			numerator := c.a
			denominator := 1 << c.comboOperand(inst.operand)
			c.c = numerator / denominator
		}
		c.pc += 1
	}
	return output
}

func Part1() string {
	input := string(common.Input())
	st,_ := G.Parse(input)
	c := loadComputer(st)
	out := c.execute()
	return toString(out)
}

func Part2() int {
	input := string(common.Input())
	st,_ := G.Parse(input)
	expected := strings.Split(st.Find("instructions")[0].Value(), ",")
	c := loadComputer(st)
	return reverse(c, expected, 0, 0) 
}

func reverse(initial computer, expected []string, input, n int) int {
	for i := 0; i < 8; i++ {
		c := initial
		input2 := (input << 3) | i
		c.a = input2
		out := c.execute()
		if len(out) == n+1 && len(out) <= len(expected) {
            match := true
            for j, v := range expected[len(expected) - len(out):] {
                if out[j] != common.Net(strconv.Atoi(v)) {
                    match = false
                }
            }
            if !match {
                continue
            }
			if len(out) == len(expected) {
				if toString(out) == strings.Join(expected, ",") {
					return input2
				}
			} else {
				req := reverse(initial, expected, input2, n+1)
				if req != -1 {
					return req
				}
			}
		}
	}
	return -1

}

func toString(a []int) string {
	sout := make([]string, len(a))
	for i, v := range a {
		sout[i] = strconv.Itoa(v)
	}
	return strings.Join(sout, ",")
}
