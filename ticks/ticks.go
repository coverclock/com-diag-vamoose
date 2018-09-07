/* vi: set ts=4 expandtab shiftwidth=4: */

package ticks

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Basic time-related functions for use in Vamoose. We use the go Time types
// and functions but put a thin abstraction layer around them so it will be
// easily changed in the future if necessary. (For example, I might end up
// using the low level runtime now() function instead.)

import (
	"time"
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

// Return the resolution of the time value that can be stored in a Tick in
// units of Hertz.
func Frequency() Ticks {
	return Ticks(FREQUENCY)
}

// Return the current value of Ticks for the monotonically increasing time that
// is now.
func Now() Ticks {
	return Ticks(time.Now().Sub(epoch))
}
