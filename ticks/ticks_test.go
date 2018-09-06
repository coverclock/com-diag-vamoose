package ticks

import (
    "testing"
)

func TestFREQUENCY(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Error("FREQUENCY") }
}

func testFrequency(t * testing.T) {
    if (Frequency() == 1000000000) {} else { t.Error("FREQUENCY") }	
}

func testNow(t * testing.T) {
	ticks := Now()
	if (ticks > 0) {} else { t.Error("Now") }
}
