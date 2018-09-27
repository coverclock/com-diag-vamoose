/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
package ticks

import (
    "testing"
    "unsafe"
    "time"
)

func TestTime(t * testing.T) {
    t.Logf("before=%s\n", time.Now().String());
	Sleep(Frequency())
    t.Logf("after=%s\n", time.Now().String());
}

func TestTypes(t * testing.T) {
    t.Logf("Ticks: Alignof=%v Sizeof=%v\n", unsafe.Alignof(Ticks(0)), unsafe.Sizeof(Ticks(0)));
}

func TestFREQUENCY(t * testing.T) {
    if (FREQUENCY == 1000000000) {} else { t.Fatalf("FAILED!\n") }
}

func TestFrequency(t * testing.T) {
   if (Frequency() == 1000000000) {} else { t.Fatalf("FAILED!\n") }	
}

func TestSleepOne(t * testing.T) {
	before := Now()
	t.Logf("before=%vns\n", before);
	if (before > 0) {} else { t.Fatalf("FAILED!\n") }
	delay := Frequency()
	Sleep(delay)
	after := Now()
	t.Logf("after=%vns\n", after);
	if (after > 0) {} else { t.Fatalf("FAILED!\n") }
	if after > before {} else { t.Fatalf("FAILED!\n") }
	ticks := after - before
	seconds := float64(ticks) / float64(FREQUENCY)
	t.Logf("elapsed=%vns=%vs\n", ticks, seconds);
	if seconds >= 1.0 {} else { t.Fatalf("FAILED!\n") }
	if seconds < 2.0 {} else { t.Fatalf("FAILED!\n") }
}

func TestSleepZero(t * testing.T) {
	before := Now()
	t.Logf("before=%vns\n", before);
	if (before > 0) {} else { t.Fatalf("FAILED!\n") }
	Sleep(0)
	after := Now()
	t.Logf("after=%vns\n", after);
	if (after > 0) {} else { t.Fatalf("FAILED!\n") }
	if after > before {} else { t.Fatalf("FAILED!\n") }
	ticks := after - before
	seconds := float64(ticks) / float64(FREQUENCY)
	t.Logf("elapsed=%vns=%vs\n", ticks, seconds);
}
