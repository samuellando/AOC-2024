package main

import (
	"fmt"
	"sort"
    "advent/common"
)

func main() {
    fmt.Println("Part 1:")
    part1()
    fmt.Println("Part 2:")
    part2()
}

func part1() {
    lines := common.GetInputs()
    r := make([]int, 0)
    l := make([]int, 0)
    for _, line := range(lines) {
        r = append(r, line[0])
        l = append(l, line[len(line) - 1])
    }
    sort.Sort(sort.IntSlice(r))
    sort.Sort(sort.IntSlice(l))
    sum := 0
    for i := 0; i < len(r); i++ {
        sum += common.Abs(r[i] - l[i])
    }
    fmt.Println(sum)
}

func part2() {
    lines := common.GetInputs()
    r := make(map[int]int)
    l := make([]int, 0)
    for _, line := range(lines) {
        ri := line[0]
        li := line[len(line) - 1]
        if _, ok := r[ri]; ok {
            r[ri] += 1
        } else {
            r[ri] = 1
        }
        l = append(l, li)
    }
    sum := 0
    for i := 0; i < len(l); i++ {
        if val, ok := r[l[i]]; ok {
            sum += val * l[i]
        } 
    }
    fmt.Println(sum)
}
