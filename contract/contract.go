/* vi: set ts=4 expandtab shiftwidth=4: */

package contract

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Implements a composite throttle consisting of two throttles, the first
// containing the sustained rate contract, the second containing the peak
// rate contract. To be admissable, an event must be admitted to both
// throttles. The peak contract consists of the peak increment and the
// peak limit that is the jitter tolerance. The sustained contract
// consists of the sustained increment and the sustained limit computed from
// maximum burst size and the jitter tolerance.

import (
	"fmt"
	"unsafe"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
	"github.com/coverclock/com-diag-vamoose/throttle"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Contract is a gcra that is a composite of a sustained throttle
// and a peak throttle.
type Contract struct {
    peak throttle.Throttle
    sustained throttle.Throttle
}

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

// String returns a printable string showing the guts of the throttle.
func (that * Contract) String() string {
    return fmt.Sprintf("Contract@%p[%d]:{p:(%s},s:{%s}}",
		unsafe.Pointer(that), unsafe.Sizeof(*that),
        that.peak.String(), that.sustained.String());
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a throttle back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (that * Contract) Reset(now ticks.Ticks) {
    that.sustained.Reset(now)
    that.peak.Reset(now)
}

/*******************************************************************************
 * CONSTRUCTORS
 ******************************************************************************/

// Init initialize a throttle, setting its traffic contract parameters: peak is
// the peak increment, jitter is the peak limit, sustained is the sustained
// increment, and burst is the maximum burst size from which is computed the
// sustained limit.
func (that * Contract) Init(peak ticks.Ticks, jitter ticks.Ticks, sustained ticks.Ticks, burst gcra.Events, now ticks.Ticks) {
    that.peak.Init(peak, jitter, now)
    limit := jitter
    if (burst <= 1) {
        // Do nothing.
    } else if (peak >= sustained) {
        // Do nothing.
    } else {
        limit += ticks.Ticks(burst - 1) * (sustained - peak)
    }
    that.sustained.Init(sustained, limit, now)
	that.Reset(now)
}

/*******************************************************************************
 * DESTRUCTORS
 ******************************************************************************/

// Fini handles any cleanup necessary before a throttle is deallocated. It is
// deferred when the throttle is constructed by New. It is also callable as
// part of the API, although doing so may render the throttle unusable.
func (that * Contract) Fini() {
	// Do nothing.
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocates initialize a throttle, setting its traffic contract parameters:
// peak is the peak increment, jitter is the peak limit, sustained is the
// sustained increment, and burst is the maximum burst size from which is
// computed the sustained limit.
func New(peak ticks.Ticks, jitter ticks.Ticks, sustained ticks.Ticks, burst gcra.Events, now ticks.Ticks) * Contract {
	throttle := new(Contract)
	defer throttle.Fini()
	throttle.Init(peak, jitter, sustained, burst, now)
	return throttle
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// isEmpty returns true if the throttle is empty, that is, it has no accumulated
// early ticks.
func (that * Contract) IsEmpty() bool {
    p := that.peak.IsEmpty()
    s := that.sustained.IsEmpty()
	return (p && s)
}

// IsFull returns true if the throttle is full, that is, its accumulated early
// ticks is greater than or equal to its limit.
func (that * Contract) IsFull() bool {
    p := that.peak.IsFull()
    s := that.sustained.IsFull()
	return (p || s)
}

// IsAlarmed returns true if the throttle is alarmed, that is, its accumulated
// early ticks is greater than its limit, indicating that the event emission
// stream is out of compliance with the traffic contract.
func (that * Contract) IsAlarmed() bool {
    p := that.peak.IsAlarmed()
    s := that.sustained.IsAlarmed()
	return (p || s)
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

// Emptied is true if the throttle just emptied in the last action.
func (that * Contract) Emptied() bool {
    p := that.peak.Emptied()
    s := that.sustained.Emptied()
	return (p && s)
}

// Filled is true if the throttle just filled in the last action.
func (that * Contract) Filled() bool {
    p := that.peak.Filled()
    s := that.sustained.Filled()
	return (p || s)
}

// Alarmed is true if the throttle just alarmed in the last action.
func (that * Contract) Alarmed() bool {
    p := that.peak.Alarmed()
    s := that.sustained.Alarmed()
	return (p || s)
}

// Cleared is true if the throttle just unalarmed in the last action, indicating
// that the event emission stream has returned to being compliant with the
// traffic contract.
func (that * Contract) Cleared() bool {
    p := that.peak.Cleared()
    s := that.sustained.Cleared()
	return (p && s)
}

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

// Request asks given the current time in ticks how long of a delay in ticks
// would be necessary before the next event were emitted for that emission to be
// in compliance with the traffic contract.
func (that * Contract) Request(now ticks.Ticks) ticks.Ticks {
    var delay ticks.Ticks = 0
    
    p := that.peak.Request(now)
    s := that.sustained.Request(now)
    if (p > s) {
        delay = p
    } else {
        delay = s
    }
    
    return delay
}

// Commits updates the throttle with the number of events having been emitted
// starting at the time specified in the previous Request, and returns false
// if the throttle is alarmed, indicating the application might want to slow it
// down a bit, true otherwise.
func (that * Contract) Commits(events gcra.Events) bool {
    p := that.peak.Commits(events)
    s := that.sustained.Commits(events)
    
    return (p && s)
}

// Commit is equivalent to calling Commits with one event.
func (that * Contract) Commit() bool {
    p := that.peak.Commit()
    s := that.sustained.Commit()
    
    return (p && s)
}

// Admits combines calling Request with the current time in ticks with calling
// and returning the value of Commits with the number of events.
func (that * Contract) Admits(now ticks.Ticks, events gcra.Events) bool {
    p := that.peak.Admits(now, events)
    s := that.sustained.Admits(now, events)
    
    return (p && s)
}

// Admit is equivalent to calling Admits with one event.
func (that * Contract) Admit(now ticks.Ticks) bool {
    p := that.peak.Admit(now)
    s := that.sustained.Admit(now)
    
    return (p && s)
}

// Update is equivalent to calling Admits with zero events. It is a way to
// update the throttle with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the throttle back into compliance with the traffic contract.
func (that * Contract) Update(now ticks.Ticks) bool {
    p := that.peak.Update(now)
    s := that.sustained.Update(now)
    
    return (p && s)
}
