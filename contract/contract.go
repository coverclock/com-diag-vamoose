/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Implements a traffic contract that is a composite of two GCRAs: one that
// describes the peak rate, and one that describes the sustainable rate. The
// event stream must conform to both GCRAs. The interface still appears to be
// a single GCRA from the point of view of the calling application. The
// implementation consists of two throttles, one for the peak GCRA, the other
// for the sustained GCRA. The peak throttle contains the peak increment, and
// the peak limit that is the jitter tolerance. The sustained throttle contains
// the sustained increment, and the sustained limit computed from maximum burst
// size and the jitter tolerance.
//
package contract

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
func (this * Contract) String() string {
    return fmt.Sprintf("Contract@%p[%d]:{p:(%s},s:{%s}}",
		unsafe.Pointer(this), unsafe.Sizeof(*this),
        this.peak.String(), this.sustained.String());
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a throttle back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (this * Contract) Reset(now ticks.Ticks) {
    this.sustained.Reset(now)
    this.peak.Reset(now)
}

/*******************************************************************************
 * CONSTRUCTORS
 ******************************************************************************/

// Init initialize a throttle, setting its traffic contract parameters: peak is
// the peak increment, jittertolerance is the peak limit, sustained is the
// sustained increment, and bursttolerance is the sustained limit.
func (this * Contract) Init(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, bursttolerance ticks.Ticks, now ticks.Ticks) {
    this.peak.Init(peak, jittertolerance, now)
    this.sustained.Init(sustained, bursttolerance, now)
	this.Reset(now)
}

/*******************************************************************************
 * DESTRUCTORS
 ******************************************************************************/

// Fini handles any cleanup necessary before a throttle is deallocated. It is
// deferred when the throttle is constructed by New. It is also callable as
// part of the API, although doing so may render the throttle unusable.
func (this * Contract) Fini() {
	// Do nothing.
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocates initialize a throttle, setting its traffic contract parameters:
// peak is the peak increment, jittertolerance is the peak limit, sustained is
// the sustained increment, and bursttolerance is the sustained limit.
func New(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, bursttolerance ticks.Ticks, now ticks.Ticks) * Contract {
	contract := new(Contract)
	defer contract.Fini()
	contract.Init(peak, jittertolerance, sustained, bursttolerance, now)
	return contract
}

/*******************************************************************************
 * MUTATORS
 ******************************************************************************/

// Request computes, given the current time in ticks how long of a delay in
// ticks
// would be necessary before the next event were emitted for that emission to be
// in compliance with the traffic contract.
func (this * Contract) Request(now ticks.Ticks) ticks.Ticks {
    var delay ticks.Ticks = 0
    
    p := this.peak.Request(now)
    s := this.sustained.Request(now)
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
func (this * Contract) Commits(events gcra.Events) bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Commits(events)
    s := this.sustained.Commits(events)
    return (p && s)
}

// Commit is equivalent to calling Commits with one event.
func (this * Contract) Commit() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Commit()
    s := this.sustained.Commit()
    return (p && s)
}

// Admits combines calling Request with the current time in ticks with calling
// and returning the value of Commits with the number of events.
func (this * Contract) Admits(now ticks.Ticks, events gcra.Events) bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Admits(now, events)
    s := this.sustained.Admits(now, events)
    return (p && s)
}

// Admit is equivalent to calling Admits with one event.
func (this * Contract) Admit(now ticks.Ticks) bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Admit(now)
    s := this.sustained.Admit(now)
    return (p && s)
}

// Update is equivalent to calling Admits with zero events. It is a way to
// update the throttle with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the throttle back into compliance with the traffic contract.
func (this * Contract) Update(now ticks.Ticks) bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Update(now)
    s := this.sustained.Update(now)
    return (p && s)
}

// Comply computes the number of ticks it would be necessary for the caller to
// delay for the event stream  to comply to the traffic contract with no limit
// penalty accumulated.
func (this * Contract) Comply() ticks.Ticks {
    var delay ticks.Ticks = 0
    
    p := this.peak.Comply()
    s := this.sustained.Comply()
    if (p > s) {
        delay = p
    } else {
        delay = s
    }
    
    return delay
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// isEmpty returns true if the throttle is empty, that is, it has no accumulated
// early ticks.
func (this * Contract) IsEmpty() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.IsEmpty()
    s := this.sustained.IsEmpty()
	return (p && s)
}

// IsFull returns true if the throttle is full, that is, its accumulated early
// ticks is greater than or equal to its limit.
func (this * Contract) IsFull() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.IsFull()
    s := this.sustained.IsFull()
	return (p || s)
}

// IsAlarmed returns true if the throttle is alarmed, that is, its accumulated
// early ticks is greater than its limit, indicating that the event emission
// stream is out of compliance with the traffic contract.
func (this * Contract) IsAlarmed() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.IsAlarmed()
    s := this.sustained.IsAlarmed()
	return (p || s)
}

/*******************************************************************************
 * SENSORS
 ******************************************************************************/

// Emptied returns true if the throttle just emptied in the last action.
func (this * Contract) Emptied() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Emptied()
    s := this.sustained.Emptied()
	return (p || s)
}

// Filled returns true if the throttle just filled in the last action.
func (this * Contract) Filled() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Filled()
    s := this.sustained.Filled()
	return (p || s)
}

// Alarmed returns true if the throttle just alarmed in the last action.
func (this * Contract) Alarmed() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Alarmed()
    s := this.sustained.Alarmed()
	return (p || s)
}

// Cleared returns true if the throttle just unalarmed in the last action,
// indicating that the event emission stream has returned to being compliant
// with the traffic contract.
func (this * Contract) Cleared() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Cleared()
    s := this.sustained.Cleared()
	return (p || s)
}
