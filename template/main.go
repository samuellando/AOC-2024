package main

import (
	"advent/common"
	"fmt"
)

func main() {
    fmt.Println("Part 1:")
	part1()
    fmt.Println("Part 2:")
	part2()
}

func part1() {
    for line := range(common.AsInts(common.InputLines())) {
        fmt.Println(line)
    }
}
func part2() {
	// lines := common.GetInputs()
	count := 0
	fmt.Println(count)
}
