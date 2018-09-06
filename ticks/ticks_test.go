package ticks

import (
    "testing"
)

func TestTicks(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Error("FAILED!") }

    if (Frequency() == 1000000000) {} else { t.Error("FAILED!") }	

	before := Now()
	if (before > 0) {} else { t.Error("FAILED!") }

	after := Now()
	if (after > 0) {} else { t.Error("FAILED!") }
	if (after > before) {} else { t.Error("FAILED!") }
}
