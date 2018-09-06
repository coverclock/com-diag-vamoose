/* vi: set ts=4 expandtab shiftwidth=4: */

package throttle

/**
 * @file
 *
 * Copyright 2018 Digital Aggregates Corporation, Colorado, USA<BR>
 * Licensed under the terms in LICENSE.txt<BR>
 * Chip Overclock <coverclock@diag.com><BR>
 * https://github.com/coverclock/com-diag-vamoose<BR>
 *
 * ABSTRACT
 *
 * Implements a Generic Cell Rate Algorithm (GCRA) using a Virtual Scheduler.
 * This can in turn be used to implement a variety of traffic shaping and rate
 * control algorithms. The VS works by monitoring the inter-arrival interval of
 * events and comparing that interval to the expected value. When the cumulative
 * error in the inter-arrival interval exceeds a threshold, the throttle becomes
 * "alarmed" and the traffic stream is in violation of its contract. In the
 * original TM spec, an event was the emission (if traffic shaping) or arrival
 * (if traffic policing) of an ATM cell, but it could be data blocks, error
 * reports, or any other kind of real-time activity. In this implementation,
 * it can even be variable length data blocks, in which the traffic contract
 * describes the mean bandwidth of the traffic stream, not the instantaneous
 * bandwidth as with ATM. In the original TM spec, the variable "i" was the
 * increment or contracted inter-arrival interval, "l" was the limit or
 * threshold, "x" was the expected inter-arrival interval for the next event,
 * and "x1" was the actual inter-arrival interval of that event. A throttle can
 * be used to smooth out low frequency events over a long duration, or to
 * implement a leaky bucket algorithm.
 *
 * REFERENCES
 *
 * ATM Forum, Traffic Management Specification Version 4.1, af-tm-0121.000,
 * 1999-03
 *
 * Chip Overclock, "Traffic Management", 2006-12,
 * http://coverclock.blogspot.com/2006/12/traffic-management.html
 *
 * Chip Overclock, "Rate Control Using Throttles", 2007-01,
 * http://coverclock.blogspot.com/2007/01/rate-control-and-throttles.html
 *
 * Chip Overclock, "Traffic Contracts", 2007-01,
 * http://coverclock.blogspot.com/2007/01/traffic-contracts.html
 *
 * Chip Overclock, Diminuto, https://github.com/coverclock/com-diag-diminuto
 */

import (
	"fmt"
	"unsafe"
	"github.com/coverclock/com-diag-vamoose/ticks"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

type Events uint32

type Throttle struct {
	now			ticks.Ticks			// Current timestamp
	then		ticks.Ticks			// Prior timestamp
	increment	ticks.Ticks			// GCRA i
	limit		ticks.Ticks			// GCRA l
	expected	ticks.Ticks			// GCRA x
	actual		ticks.Ticks			// GCRA x1
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

func (that * Throttle) String() string {
	return fmt.Sprintf("Throttle@%p[%d]: { iat=%d i=%d l=%d x=%d x1=%d f=(%t,%t,%t) e=(%t,%t,%t) a=(%t,%t) }",
		unsafe.Pointer(that), unsafe.Sizeof(*that),
		that.now - that.then,
		that.increment, that.limit, that.expected, that.actual,
		that.full0, that.full1, that.full2,
		that.empty0, that.empty1, that.empty2,
		that.alarmed1, that.alarmed2);
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

func (that * Throttle) Reset(now ticks.Ticks) {
	that.now = now
	that.then = that.now - that.increment
	that.expected = that.increment
	that.actual = 0
	that.full0 = false
	that.full1 = false
	that.full2 = false
	that.empty0 = true
	that.empty1 = true
	that.empty2 = true
	that.alarmed1 = false
	that.alarmed2 = false
}

func (that * Throttle) Init(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) {
	that.increment = increment
	that.limit = limit
	that.Reset(now)
}

/*******************************************************************************
 * DESTRUCTORS
 ******************************************************************************/

func (that * Throttle) fini() {
	// Do nothing.
}

/*******************************************************************************
 * CONSTRUCTORS
 ******************************************************************************/

func New(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks) * Throttle {
	throttle := new(Throttle)
	defer throttle.fini()
	throttle.Init(increment, limit, now)
	return throttle
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

func (that * Throttle) IsEmpty() bool {
	return that.empty1
}

func (that * Throttle) IsFull() bool {
	return that.full1
}

func (that * Throttle) IsAlarmed() bool {
	return that.alarmed1
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

func (that * Throttle) Emptied() bool {
	return (that.empty1 && (!that.empty2))
}

func (that * Throttle) Filled() bool {
	return (that.full1 && (!that.full2))
}

func (that * Throttle) Alarmed() bool {
	return (that.alarmed1 && (!that.alarmed2))
}

func (that * Throttle) Cleared() bool {
	return ((!that.alarmed1) && that.alarmed2)
}

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

func (that * Throttle) Request(now ticks.Ticks) ticks.Ticks {
	var delay ticks.Ticks
	var elapsed ticks.Ticks
	
	that.now = now
	elapsed = that.now - that.then
	if (that.expected <= elapsed) {
		that.actual = 0
		that.full0 = false
		that.empty0 = true
		delay = 0
	} else {
		that.actual = that.expected - elapsed
		if (that.actual <= that.limit) {
			that.full0 = false
			that.empty0 = false
			delay = 0
		} else {
			that.full0 = true
			that.empty0 = false
			delay = that.actual - that.limit
		}
	}

	return delay
}

func (that * Throttle) Commits(events Events) bool {
	that.then = that.now
	that.expected = that.actual;
	if (events > 0) {
	     that.expected = that.expected + (that.increment * ticks.Ticks(events))
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

func (that * Throttle) Commit() bool {
	return that.Commits(1)
}

func (that * Throttle) Admits(now ticks.Ticks, events Events) bool {
	that.Request(now)
	return that.Commits(events)
}

func (that * Throttle) Admit(now ticks.Ticks) bool {
	return that.Admits(now, 1)
}

func (that * Throttle) Update(now ticks.Ticks) bool {
	return that.Admits(now, 0)
}
