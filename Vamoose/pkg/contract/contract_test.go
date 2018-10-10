/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
package contract

import (
    "testing"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/gcra"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/harness"
    "fmt"
)

/*******************************************************************************
 * ANCILLARY
 ******************************************************************************/

func TestBurstTolerance(t * testing.T) {
    var bt ticks.Ticks

    // func BurstTolerance(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, burstsize Events) ticks.Ticks

    bt = BurstTolerance(2, 3, 7, 5) // Nominal
    if (bt == ticks.Ticks(23)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 0, 7, 5) // No CDVT
    if (bt == ticks.Ticks(20)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(0, 0, 2, 5) // No PCR, CDVT
    if (bt == ticks.Ticks(8)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(7, 3, 7, 5) // PCR == SCR
    if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 3, 7, 1) // MBS <= 1
    if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 3, 7, 0) // MBS <= 1
    if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

}

/*******************************************************************************
 * MANUAL
 ******************************************************************************/

func TestContractSandbox(t * testing.T) {
    const PEAK ticks.Ticks = 4
    const JITTER ticks.Ticks = 2
    const SUSTAINED ticks.Ticks = 8
    const BURST throttle.Events = 16
    var now ticks.Ticks = 1000000000
    var delay ticks.Ticks = 0
    var admissable bool = false
    var expected ticks.Ticks = 0

    fmt.Printf("now=%v\n", now)
    bt := BurstTolerance(PEAK, JITTER, SUSTAINED, BURST)
    that := New(PEAK, JITTER, SUSTAINED, bt, now)
    fmt.Printf("that=%s\n", that.String())

    now += 4
    fmt.Printf("now=%v\n", now)
    delay = that.Request(now)
    fmt.Printf("Request=%v\n", delay);
    fmt.Printf("that=%s\n", that.String())

    admissable = that.Commits(BURST)
    fmt.Printf("Commits=%v\n", admissable)
    fmt.Printf("that=%s\n", that.String())

    now += 4
    fmt.Printf("now=%v\n", now)
    delay = that.Request(now)
    fmt.Printf("Request=%v\n", delay);
    fmt.Printf("that=%s\n", that.String())

    admissable = that.Commits(BURST)
    fmt.Printf("Commits=%v\n", admissable)
    fmt.Printf("that=%s\n", that.String())

    expected = that.Expected()
    fmt.Printf("Comply=%v\n", expected);

    now += expected
    fmt.Printf("now=%v\n", now)
    admissable = that.Update(now)
    fmt.Printf("Update=%v\n", admissable);
    fmt.Printf("that=%s\n", that.String())
}

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func TestContractSimulated1(t * testing.T) {
    const PEAK ticks.Ticks = 1024 // Bytes per second.
    const SUSTAINED ticks.Ticks = 512 // Bytes per second.
    const BURST int = 32768// Bytes
    const OPERATIONS int = 1000000

    frequency := ticks.Frequency()
    peak := gcra.Increment(throttle.Events(PEAK), 1, frequency)
    burst := throttle.Events(BURST)
    jt := gcra.JitterTolerance(peak, burst)
    sustained := gcra.Increment(throttle.Events(SUSTAINED), 1, frequency)
    bt := BurstTolerance(peak, jt, sustained, burst)
    now := ticks.Now()

    shape := New(peak, 0, sustained, bt, now)
    police := New(peak, jt, sustained, bt, now)

    harness.SimulatedEventStream(t, shape, police, BURST, OPERATIONS)

}

func TestContractSimulated2(t * testing.T) {
    const PEAK ticks.Ticks = 2048 // Bytes per second.
    const SUSTAINED ticks.Ticks = 1024 // Bytes per second.
    const BURST int = 512// Bytes
    const OPERATIONS int = 1000000

    frequency := ticks.Frequency()
    peak := gcra.Increment(throttle.Events(PEAK), 1, frequency)
    burst := throttle.Events(BURST)
    jt := gcra.JitterTolerance(peak, burst)
    sustained := gcra.Increment(throttle.Events(SUSTAINED), 1, frequency)
    bt := BurstTolerance(peak, jt, sustained, burst)
    now := ticks.Now()

    shape := New(peak, 0 /* jt */, sustained, bt, now)
    police := New(peak, jt, sustained, bt, now)

    harness.SimulatedEventStream(t, shape, police, BURST, OPERATIONS)

}

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

func TestContractActual(t * testing.T) {
    const PEAK int = 1024               // bytes per second
    const SUSTAINED int = 512           // bytes per second
    const BURST int = 64                // bytes
    const TOTAL uint64 = 1024 * 60      // bytes

    supply := make(chan byte, BURST + 1) // +1 for EOR indicator. Producer closes.
    demand := make(chan byte, BURST) // Policer closes.

    frequency := ticks.Frequency()
    peak := gcra.Increment(throttle.Events(PEAK), 1, frequency)
    burst := throttle.Events(BURST)
    jt := gcra.JitterTolerance(peak, burst)
    sustained := gcra.Increment(throttle.Events(SUSTAINED), 1, frequency)
    bt := BurstTolerance(peak, jt, sustained, burst)
    now := ticks.Now()

    shape := New(peak, 0, sustained, bt, now)
    police := New(peak, jt, sustained, bt, now)

    harness.ActualEventStream(t, shape, police, supply, demand, TOTAL)

}
