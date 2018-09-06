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
	before := Now()
	if (before > 0) {} else { t.Error("Now before") }
	after := Now()
	if (after > 0) {} else { t.Error("Now after") }
	if (after > before) {} else { t.Error("Now elapsed") }
}
