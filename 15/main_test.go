
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 1406392 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 1429013 {
        t.Fail()
    }
}
