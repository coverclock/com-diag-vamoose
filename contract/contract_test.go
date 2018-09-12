package contract

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
    "math/rand"
)

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func TestContractSimulated(t * testing.T) {
    const PEAK ticks.Ticks = 1024 // Bytes per second.
    const TOLERANCE ticks.Ticks = 64
    const SUSTAINED ticks.Ticks = 512 // Bytes per second.
	const BURST gcra.Events = 32768
    const OPERATIONS uint = 1000000
	const MARGIN ticks.Ticks = 200 // 0.5%
	var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
	var size gcra.Events = 0
    var total uint64 = 0
    var admissable bool = false
    var iops uint = 0
    var largest gcra.Events = 0
    var that gcra.Gcra
    
    frequency := ticks.Frequency()
    peak := (frequency + PEAK - 1) / PEAK
    sustained := (frequency + SUSTAINED - 1) / SUSTAINED
    now := ticks.Now()
    
	that = New(peak, TOLERANCE, sustained, BURST, now)
	t.Log(that.String())
	
	for iops = 0; iops < OPERATIONS; iops += 1 {

	    delay = that.Request(now)
	    if (delay >= 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }
	    now += delay
	    if (now >= 0) {} else { t.Errorf("OVERFLOW! %v", now) }
	    duration += delay
	    if (duration >= 0) {} else { t.Errorf("OVERFLOW! %v", duration) }
	    
	    delay = that.Request(now)
	    if (delay == 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }

        size = gcra.Events(rand.Int63n(int64(BURST))) + 1
	    if (0 < size) {} else { t.Errorf("FAILED! %v", size) }
	    if (size <= gcra.Events(BURST)) {} else { t.Errorf("FAILED! %v", size) }
	    if (size > largest) { largest = size }
	    total += uint64(size)
	    if (total > 0) {} else { t.Errorf("OVERFLOW! %v", total) }

	    admissable = that.Commits(size)
	    if (admissable) {} else { t.Errorf("FAILED! %v", admissable); t.Log(that.String()) }

	}
	
	blocksize := float64(total) / float64(OPERATIONS)
	seconds := float64(duration) / float64(frequency)
	interarrival := seconds / float64(OPERATIONS)
	t.Logf("total=%vB largest=%vB/io mean=%vB/io mean=%vs/io\n", total, largest, blocksize, interarrival)
	if (duration > frequency) {} else { t.Errorf("FAILED! %v", duration) }

	bandwidth := float64(total) / float64(seconds)
	delta := bandwidth - float64(SUSTAINED)
	if (delta < 0) { delta = -delta }
    margin := float64(SUSTAINED) / float64(MARGIN)
	t.Logf("sustained=%vB/s delta=%vB/s margin=%vB/s\n", bandwidth, delta, margin)
	if (delta < margin) {} else { t.Errorf("FAILED! %v", delta) }
    
}

