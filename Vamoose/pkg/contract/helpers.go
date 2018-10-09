/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// This portion of the Contract package contains helper functions that
// are useful for some applications. It is seperated out for the sole purpose
// of exercising multi-file package support.
//
package contract

import (
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
)

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

// BurstTolerance computes a burst tolerance (sustained limit) from the peak
// increment (minimum interarrival time), jittertolerance (peak limit),
// sustained increment (mean interarrival time), and maximum burst size.
func BurstTolerance(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, burstsize throttle.Events) ticks.Ticks {
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

