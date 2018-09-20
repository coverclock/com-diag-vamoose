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
 	"fmt"
)

/*******************************************************************************
 * MANUAL
 ******************************************************************************/

func TestContractSandbox(t * testing.T) {
    const PEAK ticks.Ticks = 4
    const JITTER ticks.Ticks = 2
    const SUSTAINED ticks.Ticks = 8
    const BURST gcra.Events = 16
    var now ticks.Ticks = 1000000000
    var delay ticks.Ticks = 0
    var admissable bool = false
    var deficit ticks.Ticks = 0
    
    that := New(PEAK, JITTER, SUSTAINED, BURST, now)
    fmt.Printf("that=%s\n", that.String())
    
    now += 4
    delay = that.Request(now)
    fmt.Printf("delay=%v\n", delay);
    fmt.Printf("that=%s\n", that.String())
  
    admissable = that.Commits(BURST)
    fmt.Printf("admissable=%v\n", admissable)
    fmt.Printf("that=%s\n", that.String())
    
    now += 4
    delay = that.Request(now)
    fmt.Printf("delay=%v\n", delay);
    fmt.Printf("that=%s\n", that.String())
  
    admissable = that.Commits(BURST)
    fmt.Printf("admissable=%v\n", admissable)
    fmt.Printf("that=%s\n", that.String())
    
    deficit = that.GetDeficit()
    fmt.Printf("deficit=%v\n", deficit);

    now += deficit
    delay = that.Request(now)
    fmt.Printf("delay=%v\n", delay);
    fmt.Printf("that=%s\n", that.String())
  
    admissable = that.Commits(BURST)
    fmt.Printf("admissable=%v\n", admissable)
    fmt.Printf("that=%s\n", that.String())
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

