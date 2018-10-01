/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Implements a Generic Cell Rate Algorithm (GCRA) using a Virtual Scheduler.
// This can in turn be used to implement a variety of traffic shaping and rate
// control algorithms. The VS works by monitoring the inter-arrival interval of
// events and comparing that interval to the expected value. When the cumulative
// error in the inter-arrival interval exceeds a threshold, the gcra becomes
// "alarmed" and the traffic stream is in violation of its contract. In the
// original TM spec, an event was the emission (if traffic shaping) or arrival
// (if traffic policing) of an ATM cell, but it could be data blocks, error
// reports, or any other kind of real-time activity. In this implementation,
// it can even be variable length data blocks, in which the traffic contract
// describes the mean bandwidth of the traffic stream, not the instantaneous
// bandwidth as with ATM. In the original TM spec, the variable "i" was the
// increment or contracted inter-arrival interval, "l" was the limit or
// threshold, "x" was the expected inter-arrival interval for the next event,
// and "x1" was the inter-arrival deficit accumulated so far. A gcra can
// be used to smooth out low frequency events over a long duration, or to
// implement a leaky bucket algorithm.
//
// REFERENCES
//
// N. Giroux et al., Traffic Management Specification Version 4.1, ATM Forum,
// af-tm-0121.000, 1999-03
//
// C. Overclock, "Traffic Management", 2006-12,
// http://coverclock.blogspot.com/2006/12/traffic-management.html
//
// C. Overclock, "Rate Control Using Throttles", 2007-01,
// http://coverclock.blogspot.com/2007/01/rate-control-and-throttles.html
//
// C. Overclock, "Traffic Contracts", 2007-01,
// http://coverclock.blogspot.com/2007/01/traffic-contracts.html
//
// J. Sloan, "ATM Traffic Management", 2005-08,
// http://www.diag.com/reports/ATMTrafficManagement.html
//
package gcra

import (
    "fmt"
    "unsafe"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Gcra is the type that contains the state of the gcra. It is based
// on the Virtual Scheduler implementation for the Generic Cell Rate Algorithm
// (GCRA) in the ATM Forum Traffic Management (TM) 4.1 standard. Gcras can
// be used by a producer for traffic shaping, or by a consumer for traffic
// policing.
type Gcra struct {
    now         ticks.Ticks         // Current timestamp
    then        ticks.Ticks         // Prior timestamp
    increment   ticks.Ticks         // GCRA i: ticks per event
    limit       ticks.Ticks         // GCRA l: maximum deficit ticks
    expected    ticks.Ticks         // GCRA x: expected ticks until next event
    deficit     ticks.Ticks         // GCRA x1: current deficit ticks
    full0       bool                // The leaky bucket will fill.
    full1       bool                // The leaky bucket is filling.
    full2       bool                // The leaky bucket was filled.
    empty0      bool                // The leaky bucket will empty.
    empty1      bool                // The leaky bucket is emptying.
    empty2      bool                // The leaky bucket was emptied.
    alarmed1    bool                // The gcra is alarmed.
    alarmed2    bool                // The gcra was alarmed.
}

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

func btoi(value bool) int {
    if value { return 1 } else { return 0 }
}

// String returns a printable string showing the guts of the gcra.
func (this * Gcra) String() string {
    return fmt.Sprintf("Gcra@%p[%d]:{T:%d,i:%d,l:%d,x:%d,d:%d,D:%d,f:{%d,%d,%d},e:{%d,%d,%d},a:{%d,%d}}",
        unsafe.Pointer(this), unsafe.Sizeof(*this),
        this.now - this.then,
        this.increment, this.limit,
        this.expected, this.deficit,
        this.deficit - this.limit,
        btoi(this.full0), btoi(this.full1), btoi(this.full2),
        btoi(this.empty0), btoi(this.empty1), btoi(this.empty2),
        btoi(this.alarmed1), btoi(this.alarmed2))
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a gcra back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (this * Gcra) Reset(now ticks.Ticks) {
    this.now = now
    this.then = this.now - this.increment
    this.expected = 0
    this.deficit = 0
    this.full0 = false
    this.full1 = false
    this.full2 = false
    this.empty0 = true
    this.empty1 = true
    this.empty2 = true
    this.alarmed1 = false
    this.alarmed2 = false
}

/*******************************************************************************
 * CONSTRUCTORS
 ******************************************************************************/

// Init initialize a gcra, setting its traffic contract parameters
// increment, which is the expected interarrival time in ticks for every event,
// and limit, which is the maximum aggreate deficit ticks that can be accumulated
// before the traffic contract is violated and the gcra becomes alarmed; and
// its dynamic state, which is the current monotonic time in ticks.
func (this * Gcra) Init(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) {
    this.increment = increment
    this.limit = limit
    this.Reset(now)
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocate a new gcra. It initializes it with the traffic contract
// parameters increment, which is the expected interarrival time in ticks for
// every event, and limit, which is the maximum aggreate deficit ticks that can be
// accumulated before the traffic contract is violated and the gcra becomes
// alarmed; and its dynamic state, which is the current monotonic time in ticks.
func New(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) * Gcra {
    gcra := new(Gcra)
    gcra.Init(increment, limit, now)
    return gcra
}

/*******************************************************************************
 * MUTATORS
 ******************************************************************************/

// Request computes, given the current time in ticks, how long of a delay in
// ticks would be necessary before the next event were emitted for that
// emission to be in compliance with the traffic contract.
func (this * Gcra) Request(now ticks.Ticks) ticks.Ticks {
    var delay ticks.Ticks
    var elapsed ticks.Ticks

    this.now = now
    elapsed = this.now - this.then
    if (this.expected <= elapsed) {
        this.deficit = 0
        this.full0 = false
        this.empty0 = true
        delay = 0
    } else {
        this.deficit = this.expected - elapsed
        if (this.deficit <= this.limit) {
            this.full0 = false
            this.empty0 = false
            delay = 0
        } else {
            this.full0 = true
            this.empty0 = false
            delay = this.deficit - this.limit
        }
    }

    return delay
}

// Commits updates the gcra with the number of events having been emitted
// starting at the time specified in the previous Request, and returns false
// if the gcra is alarmed, indicating the application might want to slow it
// down a bit, true otherwise.
func (this * Gcra) Commits(events throttle.Events) bool {
    this.then = this.now
    this.expected = this.deficit;
    if (events <= 0) {
        // Do nothing.
    } else if (events == 1) {
         this.expected += this.increment
    } else {
         this.expected += this.increment * ticks.Ticks(events)
    }
    this.full2 = this.full1
    this.full1 = this.full0
    this.empty2 = this.empty1
    this.empty1 = this.empty0
    this.alarmed2 = this.alarmed1
    if (this.Emptied()) {
        this.alarmed1 = false
    } else if (this.Filled()) {
        this.alarmed1 = true
    } else {
        // Do nothing.
    }

    return !this.alarmed1
}

// Commit is equivalent to calling Commits with one event.
func (this * Gcra) Commit() bool {
    return this.Commits(1)
}

// Admits combines calling Request with the current time in ticks with calling
// and returning the value of Commits with the number of events.
func (this * Gcra) Admits(now ticks.Ticks, events throttle.Events) bool {
    this.Request(now)
    return this.Commits(events)
}

// Admit is equivalent to calling Admits with one event.
func (this * Gcra) Admit(now ticks.Ticks) bool {
    return this.Admits(now, 1)
}

// Update is equivalent to calling Admits with zero events. It is a way to
// update the gcra with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the gcra back into compliance with the traffic contract.
func (this * Gcra) Update(now ticks.Ticks) bool {
    return this.Admits(now, 0)
}

// Comply computes the number of ticks that would be necessary for the caller
// to delay for the event stream  to comply to the traffic contract with no
// limit penalty accumulated given the current state of the gcra.
func (this * Gcra) Comply() ticks.Ticks {
    return this.expected
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// isEmpty returns true if the gcra is empty, that is, it has no accumulated
// deficit ticks.
func (this * Gcra) IsEmpty() bool {
    return this.empty1
}

// IsFull returns true if the gcra is full, that is, its accumulated deficit
// ticks is greater than or equal to its limit.
func (this * Gcra) IsFull() bool {
    return this.full1
}

// IsAlarmed returns true if the gcra is alarmed, that is, its accumulated
// deficit ticks is greater than its limit, indicating that the event emission
// stream is out of compliance with the traffic contract.
func (this * Gcra) IsAlarmed() bool {
    return this.alarmed1
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

// Emptied is true if the gcra just emptied in the last action.
func (this * Gcra) Emptied() bool {
    return (this.empty1 && (!this.empty2))
}

// Filled is true if the gcra just filled in the last action.
func (this * Gcra) Filled() bool {
    return (this.full1 && (!this.full2))
}

// Alarmed is true if the gcra just alarmed in the last action.
func (this * Gcra) Alarmed() bool {
    return (this.alarmed1 && (!this.alarmed2))
}

// Cleared is true if the gcra just unalarmed in the last action, indicating
// that the event emission stream has returned to being compliant with the
// traffic contract.
func (this * Gcra) Cleared() bool {
    return ((!this.alarmed1) && this.alarmed2)
}

