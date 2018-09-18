package ticks

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
)

func TestFREQUENCY(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Fatalf("FAILED!\n") }
}

func TestFrequency(t * testing.T) {
   if (Frequency() == 1000000000) {} else { t.Fatalf("FAILED!\n") }	
}

func TestNow(t * testing.T) {
	before := Now()
	t.Logf("before=%vns\n", before);
	if (before > 0) {} else { t.Fatalf("FAILED!\n") }
	delay := FREQUENCY
	Sleep(delay)
	after := Now()
	t.Logf("after=%vns\n", after);
	if (after > 0) {} else { t.Fatalf("FAILED!\n") }
	t.Logf("elapsed=%vns=%vs\n", after - before, float64(after - before) / float64(FREQUENCY));
	if (after > before) {} else { t.Fatalf("FAILED!\n") }
}
