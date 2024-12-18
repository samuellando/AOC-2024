
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 187738 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 223767210249237 {
        t.Fail()
    }
}
