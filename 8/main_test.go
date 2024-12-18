
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 323 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 1077 {
        t.Fail()
    }
}
