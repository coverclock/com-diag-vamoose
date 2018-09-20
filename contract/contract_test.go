package contract

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
 	"github.com/coverclock/com-diag-vamoose/harness"
)

/*******************************************************************************
 * SANITY
 ******************************************************************************/

func TestContractSanity(t * testing.T) {
    const PEAK ticks.Ticks = 1
    const JITTER ticks.Ticks = 2
    const SUSTAINED ticks.Ticks = 3
    const BURST ticks.Ticks = 4
    var now ticks.Ticks = 0
    
    that := New(PEAK, JITTER, SUSTAINED, BURST, now)
    fmt.Printf(that.String())
    
}

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func TestContractSimulated(t * testing.T) {
    const PEAK ticks.Ticks = 1024 // Bytes per second.
    const JITTER ticks.Ticks = 64 // Ticks
    const SUSTAINED ticks.Ticks = 512 // Bytes per second.
	const BURST int = 32768
    const OPERATIONS int = 1000000

    frequency := ticks.Frequency()
    peak := (frequency + PEAK - 1) / PEAK
    sustained := (frequency + SUSTAINED - 1) / SUSTAINED
    burst := gcra.Events(BURST)
    now := ticks.Now()
    
	that := New(peak, JITTER, sustained, burst, now)
	t.Log(that.String())
	    	
	harness.SimulatedEventStream(t, that, BURST, OPERATIONS)
    
}

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

func TestContractActual(t * testing.T) {
    const PEAK int = 1024				// bytes per second
    const SUSTAINED int = 512			// bytes per second
    const BURST int = 64				// bytes
    const TOTAL uint64 = 1024 * 60		// bytes
    
    supply := make(chan byte, BURST + 1) // +1 for EOR indicator. Producer closes.
    demand := make(chan byte, BURST) // Policer closes.
           
    frequency := ticks.Frequency()
    peak := (frequency + ticks.Ticks(PEAK) - 1) / ticks.Ticks(PEAK)
    jitter := peak / 100
    sustained := (frequency + ticks.Ticks(SUSTAINED) - 1) / ticks.Ticks(SUSTAINED)
    burst := gcra.Events(BURST)
    now := ticks.Now()
    
    shape := New(peak, 0, sustained, burst, now)
    police := New(peak, jitter, sustained, burst, now)
    
    harness.ActualEventStream(t, shape, police, supply, demand, TOTAL)

}

