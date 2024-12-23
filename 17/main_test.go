
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != "7,5,4,3,4,5,3,4,6" {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 164278899142333 {
        t.Fail()
    }
}
