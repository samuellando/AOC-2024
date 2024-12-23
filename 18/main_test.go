
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 250 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != "56,8" {
        t.Fail()
    }
}
