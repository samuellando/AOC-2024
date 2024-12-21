
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 36758 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 76358113886726 {
        t.Fail()
    }
}
