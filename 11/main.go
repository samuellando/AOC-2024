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

func Part1() int {
	lines := common.AsInts(common.InputLines())
	var state []int
	for line := range lines {
		state = line
	}
	return blinker(state, 25)
}

func Part2() int {
	lines := common.AsInts(common.InputLines())
	var state []int
	for line := range lines {
		state = line
	}
	return blinker(state, 75)
}

func blinker(state []int, n int) int {
	counts := make(map[int]int)
	for _, stone := range state {
		counts[stone] = 1
	}
	for range n {
        next := make(map[int]int)
		for v, c := range counts {
			if c == 0 {
				continue
			}
			digits := len(strconv.Itoa(v))
			if v == 0 {
				addOrInsert(next, 1, c)
			} else if digits%2 == 0 {
				sv := strconv.Itoa(v)
				firstHalf := common.Net(strconv.Atoi(sv[:digits/2]))
				nextHalf := common.Net(strconv.Atoi(sv[digits/2:]))
				addOrInsert(next, firstHalf, c)
				addOrInsert(next, nextHalf, c)
			} else {
				addOrInsert(next, v*2024, c)
			}
		}
        counts = next
	}
	total := 0
	for _, c := range counts {
		total += c
	}
	return total
}

func addOrInsert(m map[int]int, k, c int) {
	if _, ok := m[k]; ok {
		m[k] += c
	} else {
		m[k] = c
	}
}

