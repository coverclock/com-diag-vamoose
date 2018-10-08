/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// This portion of the Throttles package contains helper functions that
// are useful for some applications. It is seperated out for the sole purpose
// of exercising multi-file package support.
//
package throttle

import (
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
)

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

// Increment calculates an interarrival (or interdeparture) time increment
// from the event rate in the form of a numerator and a denominator, and the
// frequency in ticks representing the time base in Hertz. Representing the
// event rate in fractional form allows for rates lower than one Hertz, or for
// non-integer rates, to be specified.
func Increment(numerator Events, denominator Events, frequency ticks.Ticks) ticks.Ticks {
    var i ticks.Ticks
    var n ticks.Ticks = ticks.Ticks(numerator)
    var d ticks.Ticks = ticks.Ticks(denominator)
        
    i = frequency
    if d > 1 {
        i *= d
    }
    if n <= 1 {
        // Do nothing.
    } else if (i % n) > 0 {
        i /= n
        i += 1
    } else {
        i /= n
    }

    return i
}

// JitterTolerance computes a jitter tolerance (peak limit) from the peak
// increment (minimum interarrival time) and maximum burst size.
func JitterTolerance(peak ticks.Ticks, burstsize Events) ticks.Ticks {
    var limit ticks.Ticks = 0

    if burstsize > 1 {
        limit = ticks.Ticks(burstsize - 1) * peak
    }
    
    return limit
}

// BurstTolerance computes a burst tolerance (sustained limit) from the peak
// increment (minimum interarrival time), jittertolerance (peak limit),
// sustained increment (mean interarrival time), and maximum burst size.
func BurstTolerance(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, burstsize Events) ticks.Ticks {
    var bt ticks.Ticks
    
    bt = jittertolerance

    if burstsize <= 1 {
        // Do nothing.
    } else if peak >= sustained {
        // Do nothing.
    } else {
        bt += ticks.Ticks(burstsize - 1) * (sustained - peak)
    }
    
    return bt
}

