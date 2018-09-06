package ticks

import (
    "testing"
)

func TestTicks(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Error() }

    if (Frequency() == 1000000000) {} else { t.Error() }	

	before := Now()
	if (before > 0) {} else { t.Error() }

	after := Now()
	if (after > 0) {} else { t.Error() }
	if (after > before) {} else { t.Error() }
}
