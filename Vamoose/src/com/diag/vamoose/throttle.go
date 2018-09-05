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
 */

import (
	"fmt"
	"time"
)

type Ticks time.Duration

const (
	FREQUENCY Ticks = 1000000000
)

type Throttle struct {
	now			Ticks				//
	then		Ticks				//
	increment	Ticks				// GCRA i
	limit		Ticks				// GCRA l
	expected	Ticks				// GCRA x
	actual		Ticks				// GCRA x1
	full0		bool				// The leaky bucket will fill.
	full1		bool				// The leaky bucket is filling.
	full2		bool				// The leaky bucket was filled.
	empty0		bool				// The leaky bucket will empty.
	empty1		bool				// The leaky bucket is emptying.
	empty2		bool				// The leaky bucket was emptied.
	alarmed1	bool				// The throttle is alarmed.
	alarmed2	bool				// The throttle was alarmed.
}

var NOW time.Time = time.Now()

func (that * Throttle) Reset() {
	
}

func (that * Throttle) Init() {
	
}

func (that * Throttle) Frequency() Ticks {
	return Ticks(FREQUENCY)
}

func (that * Throttle) Now() Ticks {
	return Ticks(time.Now().Sub(NOW))
}

func (that * Throttle) Request() {
	
}

func (that * Throttle) Commits() {
	
}

func (that * Throttle) Commit() {
	
}

func (that * Throttle) Admits() {
	
}

func (that * Throttle) Admit() {
	
}

func (that * Throttle) Update() {
	
}

func (that * Throttle) Fini() {
	
}