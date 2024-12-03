package main

import (
	"advent/common"
	"fmt"
)

func main() {
    fmt.Println("Part 1:")
	fmt.Println(Part1())
    fmt.Println("Part 2:")
    fmt.Println(Part2())
}

func Part1() int {
    for line := range(common.AsInts(common.InputLines())) {
        fmt.Println(line)
    }
    return 0
}
func Part2() int {
	// lines := common.GetInputs()
    return 0
}
