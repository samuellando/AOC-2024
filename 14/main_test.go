
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 209409792 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 8006 {
        t.Fail()
    }
}
