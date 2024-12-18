
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 66343330034722 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 637696070419031 {
        t.Fail()
    }
}
