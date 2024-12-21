
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 1396562 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 844132 {
        t.Fail()
    }
}
