package main

import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 175700056 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 71668682 {
        t.Fail()
    }
}
