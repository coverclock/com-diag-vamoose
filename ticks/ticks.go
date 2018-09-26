/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Basic time-related functions for use in Vamoose. We use the go Time types
// and functions but put a thin abstraction layer around them so it will be
// easily changed in the future if necessary.
//
package ticks

import (
	"time"
	"runtime"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Ticks is a type big enough to contain a monotonic elapsed time value.
type Ticks time.Duration

/*******************************************************************************
 * VALUES
 ******************************************************************************/

// FREQUENCY is the resolution of the time value that can be stored in a Tick
// in units of Hertz, or in otherwords, the number of Ticks in a second.
const FREQUENCY Ticks = 1000000000

// epoch is the Time Zero origin from which all Tick values are a duration.
var epoch time.Time = time.Now()

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

// Frequency returns the resolution of the time value that can be stored in a
// Tick in units of cycles per second or Hertz.
func Frequency() Ticks {
	return Ticks(FREQUENCY)
}

// Now returns the current value of Ticks for the monotonically increasing time
// that is now.
func Now() Ticks {
	return Ticks(time.Now().Sub(epoch))
}

// Sleep delays the caller for at least as many ticks as specified. If zero
// ticks are specified, the caller yields the processor.
func Sleep(ticks Ticks) {
    if ticks > 0 {
        time.Sleep(time.Duration(ticks))
    } else {
        runtime.Gosched()
    }
}