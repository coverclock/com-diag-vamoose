/* vi: set ts=4 expandtab shiftwidth=4: */

package throttle

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
// error in the inter-arrival interval exceeds a threshold, the throttle becomes
// "alarmed" and the traffic stream is in violation of its contract. In the
// original TM spec, an event was the emission (if traffic shaping) or arrival
// (if traffic policing) of an ATM cell, but it could be data blocks, error
// reports, or any other kind of real-time activity. In this implementation,
// it can even be variable length data blocks, in which the traffic contract
// describes the mean bandwidth of the traffic stream, not the instantaneous
// bandwidth as with ATM. In the original TM spec, the variable "i" was the
// increment or contracted inter-arrival interval, "l" was the limit or
// threshold, "x" was the expected inter-arrival interval for the next event,
// and "x1" was the inter-arrival deficit accumulated so far. A throttle can
// be used to smooth out low frequency events over a long duration, or to
// implement a leaky bucket algorithm.

import (
	"fmt"
	"unsafe"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Throttle is the type that contains the state of the throttle. It is based
// on the Virtual Scheduler implementation for the Generic Cell Rate Algorithm
// (GCRA) in the ATM Forum Traffic Management (TM) 4.1 standard. Throttles can
// be used by a producer for traffic shaping, or by a consumer for traffic
// policing.
type Throttle struct {
	now			ticks.Ticks			// Current timestamp
	then		ticks.Ticks			// Prior timestamp
	increment	ticks.Ticks			// GCRA i: ticks per event
	limit		ticks.Ticks			// GCRA l: maximum deficit ticks
	expected	ticks.Ticks			// GCRA x: expected ticks until next event
	deficit		ticks.Ticks			// GCRA x1: aggregate deficit ticks
	full0		bool				// The leaky bucket will fill.
	full1		bool				// The leaky bucket is filling.
	full2		bool				// The leaky bucket was filled.
	empty0		bool				// The leaky bucket will empty.
	empty1		bool				// The leaky bucket is emptying.
	empty2		bool				// The leaky bucket was emptied.
	alarmed1	bool				// The throttle is alarmed.
	alarmed2	bool				// The throttle was alarmed.
}

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

// String returns a printable string showing the guts of the throttle.
func (this * Throttle) String() string {
	return fmt.Sprintf("Throttle@%p[%d]:{t:%d,i:%d,l:%d,e:%d,d:%d,r:%d,f:{%t,%t,%t},e:{%t,%t,%t},a:{%t,%t}}",
		unsafe.Pointer(this), unsafe.Sizeof(*this),
		this.now - this.then,
		this.increment, this.limit, this.expected, this.deficit,
		this.deficit - this.limit,
		this.full0, this.full1, this.full2,
		this.empty0, this.empty1, this.empty2,
		this.alarmed1, this.alarmed2);
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a throttle back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (this * Throttle) Reset(now ticks.Ticks) {
	this.now = now
	this.then = this.now - this.increment
	this.expected = this.increment
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

// Init initialize a throttle, setting its traffic contract parameters
// increment, which is the expected interarrival time in ticks for every event,
// and limit, which is the maximum aggreate deficit ticks that can be accumulated
// before the traffic contract is violated and the throttle becomes alarmed; and
// its dynamic state, which is the current monotonic time in ticks.
func (this * Throttle) Init(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) {
	this.increment = increment
	this.limit = limit
	this.Reset(now)
}

/*******************************************************************************
 * DESTRUCTORS
 ******************************************************************************/

// Fini handles any cleanup necessary before a throttle is deallocated. It is
// deferred when the throttle is constructed by New. It is also callable as
// part of the API, although doing so may render the throttle unusable.
func (this * Throttle) Fini() {
	// Do nothing.
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocate a new throttle. It initializes it with the traffic contract
// parameters increment, which is the expected interarrival time in ticks for
// every event, and limit, which is the maximum aggreate deficit ticks that can be
// accumulated before the traffic contract is violated and the throttle becomes
// alarmed; and its dynamic state, which is the current monotonic time in ticks.
func New(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) * Throttle {
	throttle := new(Throttle)
	defer throttle.Fini()
	throttle.Init(increment, limit, now)
	return throttle
}

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

// Request asks given the current time in ticks how long of a delay in ticks
// would be necessary before the next event were emitted for that emission to be
// in compliance with the traffic contract.
func (this * Throttle) Request(now ticks.Ticks) ticks.Ticks {
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

// Commits updates the throttle with the number of events having been emitted
// starting at the time specified in the previous Request, and returns false
// if the throttle is alarmed, indicating the application might want to slow it
// down a bit, true otherwise.
func (this * Throttle) Commits(events gcra.Events) bool {
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
func (this * Throttle) Commit() bool {
	return this.Commits(1)
}

// Admits combines calling Request with the current time in ticks with calling
// and returning the value of Commits with the number of events.
func (this * Throttle) Admits(now ticks.Ticks, events gcra.Events) bool {
	this.Request(now)
	return this.Commits(events)
}

// Admit is equivalent to calling Admits with one event.
func (this * Throttle) Admit(now ticks.Ticks) bool {
	return this.Admits(now, 1)
}

// Update is equivalent to calling Admits with zero events. It is a way to
// update the throttle with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the throttle back into compliance with the traffic contract.
func (this * Throttle) Update(now ticks.Ticks) bool {
	return this.Admits(now, 0)
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// GetDeficit returns the number of ticks that would be necessary for the caller
// to delay for the event stream  to comply to the traffic contract with no
// limit penalty accumulated.
func (this * Throttle) GetDeficit() ticks.Ticks {
    return this.deficit
}

// isEmpty returns true if the throttle is empty, that is, it has no accumulated
// deficit ticks.
func (this * Throttle) IsEmpty() bool {
	return this.empty1
}

// IsFull returns true if the throttle is full, that is, its accumulated deficit
// ticks is greater than or equal to its limit.
func (this * Throttle) IsFull() bool {
	return this.full1
}

// IsAlarmed returns true if the throttle is alarmed, that is, its accumulated
// deficit ticks is greater than its limit, indicating that the event emission
// stream is out of compliance with the traffic contract.
func (this * Throttle) IsAlarmed() bool {
	return this.alarmed1
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

// Emptied is true if the throttle just emptied in the last action.
func (this * Throttle) Emptied() bool {
	return (this.empty1 && (!this.empty2))
}

// Filled is true if the throttle just filled in the last action.
func (this * Throttle) Filled() bool {
	return (this.full1 && (!this.full2))
}

// Alarmed is true if the throttle just alarmed in the last action.
func (this * Throttle) Alarmed() bool {
	return (this.alarmed1 && (!this.alarmed2))
}

// Cleared is true if the throttle just unalarmed in the last action, indicating
// that the event emission stream has returned to being compliant with the
// traffic contract.
func (this * Throttle) Cleared() bool {
	return ((!this.alarmed1) && this.alarmed2)
}

