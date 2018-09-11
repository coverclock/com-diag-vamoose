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
// and "x1" was the aggregate early duration accumulated so far. A throttle can
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
// C. Overclock, Grandote: Throttle.{h, cpp},
// https://github.com/coverclock/com-diag-grandote,
// 2005
//
// C. Overclock, Buckaroo: Throttle.java,
// https://github.com/coverclock/com-diag-buckaroo,
// 2006
//
// C. Overclock, Diminuto: diminuto_throttle.{h, c},
// https://github.com/coverclock/com-diag-diminuto,
// 2008

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
	limit		ticks.Ticks			// GCRA l: maximum early ticks
	expected	ticks.Ticks			// GCRA x: expected ticks until next event
	early		ticks.Ticks			// GCRA x1: aggregate early ticks
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
func (that * Throttle) String() string {
	return fmt.Sprintf("Throttle@%p[%d]: { e=%d i=%d l=%d x=%d x1=%d d=%d f=(%t,%t,%t) e=(%t,%t,%t) a=(%t,%t) }",
		unsafe.Pointer(that), unsafe.Sizeof(*that),
		that.now - that.then,
		that.increment, that.limit, that.expected, that.early,
		that.early - that.limit,
		that.full0, that.full1, that.full2,
		that.empty0, that.empty1, that.empty2,
		that.alarmed1, that.alarmed2);
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a throttle back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (that * Throttle) Reset(now ticks.Ticks) {
	that.now = now
	that.then = that.now - that.increment
	that.expected = that.increment
	that.early = 0
	that.full0 = false
	that.full1 = false
	that.full2 = false
	that.empty0 = true
	that.empty1 = true
	that.empty2 = true
	that.alarmed1 = false
	that.alarmed2 = false
}

/*******************************************************************************
 * Constructors
 ******************************************************************************/

// Init initialize a throttle, setting its traffic contract parameters
// increment, which is the expected interarrival time in ticks for every event,
// and limit, which is the maximum aggreate early ticks that can be accumulated
// before the traffic contract is violated and the throttle becomes alarmed; and
// its dynamic state, which is the current monotonic time in ticks.
func (that * Throttle) Init(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) {
	that.increment = increment
	that.limit = limit
	that.Reset(now)
}

/*******************************************************************************
 * DESTRUCTORS
 ******************************************************************************/

// Fini handles any cleanup necessary before a throttle is deallocated. It is
// deferred when the throttle is constructed by New. It is also callable as
// part of the API, although doing so may render the throttle unusable.
func (that * Throttle) Fini() {
	// Do nothing.
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocate a new throttle. It initializes it with the traffic contract
// parameters increment, which is the expected interarrival time in ticks for
// every event, and limit, which is the maximum aggreate early ticks that can be
// accumulated before the traffic contract is violated and the throttle becomes
// alarmed; and its dynamic state, which is the current monotonic time in ticks.
func New(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) * Throttle {
	throttle := new(Throttle)
	defer throttle.Fini()
	throttle.Init(increment, limit, now)
	return throttle
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// isEmpty returns true if the throttle is empty, that is, it has no accumulated
// early ticks.
func (that * Throttle) IsEmpty() bool {
	return that.empty1
}

// IsFull returns true if the throttle is full, that is, its accumulated early
// ticks is greater than or equal to its limit.
func (that * Throttle) IsFull() bool {
	return that.full1
}

// IsAlarmed returns true if the throttle is alarmed, that is, its accumulated
// early ticks is greater than its limit, indicating that the event emission
// stream is out of compliance with the traffic contract.
func (that * Throttle) IsAlarmed() bool {
	return that.alarmed1
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

// Emptied is true if the throttle just emptied in the last action.
func (that * Throttle) Emptied() bool {
	return (that.empty1 && (!that.empty2))
}

// Filled is true if the throttle just filled in the last action.
func (that * Throttle) Filled() bool {
	return (that.full1 && (!that.full2))
}

// Alarmed is true if the throttle just alarmed in the last action.
func (that * Throttle) Alarmed() bool {
	return (that.alarmed1 && (!that.alarmed2))
}

// Cleared is true if the throttle just unalarmed in the last action, indicating
// that the event emission stream has returned to being compliant with the
// traffic contract.
func (that * Throttle) Cleared() bool {
	return ((!that.alarmed1) && that.alarmed2)
}

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

// Request asks given the current time in ticks how long of a delay in ticks
// would be necessary before the next event were emitted for that emission to be
// in compliance with the traffic contract.
func (that * Throttle) Request(now ticks.Ticks) ticks.Ticks {
	var delay ticks.Ticks
	var elapsed ticks.Ticks
	
	that.now = now
	elapsed = that.now - that.then
	if (that.expected <= elapsed) {
		that.early = 0
		that.full0 = false
		that.empty0 = true
		delay = 0
	} else {
		that.early = that.expected - elapsed
		if (that.early <= that.limit) {
			that.full0 = false
			that.empty0 = false
			delay = 0
		} else {
			that.full0 = true
			that.empty0 = false
			delay = that.early - that.limit
		}
	}

	return delay
}

// Commits updates the throttle with the number of events having been emitted
// starting at the time specified in the previous Request, and returns true
// if the throttle is full, indicating the application might want to slow it
// down a bit.
func (that * Throttle) Commits(events gcra.Events) bool {
	that.then = that.now
	that.expected = that.early;
	if (events <= 0) {
	    // Do nothing.
	} else if (events == 1) {
	     that.expected += that.increment
	} else {
	     that.expected += that.increment * ticks.Ticks(events)
	}
	that.full2 = that.full1
	that.full1 = that.full0
	that.empty2 = that.empty1
	that.empty1 = that.empty0
	that.alarmed2 = that.alarmed1
	if (that.Emptied()) {
		that.alarmed1 = false
	} else if (that.Filled()) {
		that.alarmed1 = true
	} else {
		// Do nothing.
	}

	return that.full1
}

// Commit is equivalent to calling Commits with one event.
func (that * Throttle) Commit() bool {
	return that.Commits(1)
}

// Admits combines calling Request with the current time in ticks with calling
// and returning the value of Commits with the number of events.
func (that * Throttle) Admits(now ticks.Ticks, events gcra.Events) bool {
	that.Request(now)
	return that.Commits(events)
}

// Admit is equivalent to calling Admits with one event.
func (that * Throttle) Admit(now ticks.Ticks) bool {
	return that.Admits(now, 1)
}

// Update is equivalent to calling Admits with zero events. It is a way to
// update the throttle with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the throttle back into compliance with the traffic contract.
func (that * Throttle) Update(now ticks.Ticks) bool {
	return that.Admits(now, 0)
}
