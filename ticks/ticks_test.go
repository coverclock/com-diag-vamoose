package ticks

import (
    "testing"
)

func TestFREQUENCY(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Error("FAILED!") }
}

func TestFrequency(t * testing.T) {
   if (Frequency() == 1000000000) {} else { t.Error("FAILED!") }	
}

func TestNow(t * testing.T) {
	before := Now()
	t.Logf("before=%dns\n", before);
	if (before > 0) {} else { t.Error("FAILED!") }
	after := Now()
	t.Logf("after=%dns\n", after);
	if (after > 0) {} else { t.Error("FAILED!") }
	t.Logf("delta=%dns\n", after - before);
	if (after > before) {} else { t.Error("FAILED!") }
}
