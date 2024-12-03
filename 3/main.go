package main

import (
	"advent/common"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	pattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	input := common.Input()
	ops := pattern.FindAll(input, -1)
    sum := 0
	for _, op := range ops {
        var d1 int
        var d2 int
        fmt.Sscanf(string(op), "mul(%d,%d)", &d1, &d2)
        sum += d1 * d2
	}
    return sum
}
func Part2() int {
	pattern := regexp.MustCompile(`don't\(\)|do\(\)|mul\(\d{1,3},\d{1,3}\)`)
	input := common.Input()
	ops := pattern.FindAll(input, -1)
    sum := 0
    enabled := true
	for _, op := range ops {
        switch string(op) {
        case "do()":
            enabled = true
        case "don't()":
            enabled = false
        default:
            if enabled {
                var d1 int
                var d2 int
                fmt.Sscanf(string(op), "mul(%d,%d)", &d1, &d2)
                sum += d1 * d2
            }
        }
	}
    return sum
}
