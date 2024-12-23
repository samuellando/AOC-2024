
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 135536 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 583 {
        t.Fail()
    }
}
