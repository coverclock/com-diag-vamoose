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
// implementation consists of two gcras, one for the peak GCRA, the other
// for the sustained GCRA. The peak gcra contains the peak increment, and
// the peak limit that is the jitter tolerance. The sustained gcra contains
// the sustained increment, and the sustained limit computed from maximum burst
// size and the jitter tolerance.
//
// This package was based on the C implementation diminuto_shaper in the
// Diminuto repository. Some improvements were made in the Go implementation,
// and then those changes were back-ported from Vamoose to Diminuto 53.0.1.
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
package contract

import (
    "fmt"
    "unsafe"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/gcra"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Contract is a gcra that is a composite of a sustained gcra
// and a peak gcra.
type Contract struct {
    peak gcra.Gcra
    sustained gcra.Gcra
}

/*******************************************************************************
 * HELPERS
 ******************************************************************************/

// String returns a printable string showing the guts of the gcra.
func (this * Contract) String() string {
    return fmt.Sprintf("Contract@%p[%d]:{p:{%s},s:{%s}}",
        unsafe.Pointer(this), unsafe.Sizeof(*this),
        this.peak.String(), this.sustained.String());
}

/*******************************************************************************
 * SETTERS
 ******************************************************************************/

// Reset a gcra back to its initial state. This is used during construction,
// but can also be used by an application when a calamitous happenstance
// occurs, like the far end disconnecting and reconnecting.
func (this * Contract) Reset(now ticks.Ticks) {
    this.sustained.Reset(now)
    this.peak.Reset(now)
}

/*******************************************************************************
 * CONSTRUCTORS
 ******************************************************************************/

// Init initialize a gcra, setting its traffic contract parameters: peak is
// the peak increment, jittertolerance is the peak limit, sustained is the
// sustained increment, and bursttolerance is the sustained limit.
func (this * Contract) Init(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, bursttolerance ticks.Ticks, now ticks.Ticks) {
    this.peak.Init(peak, jittertolerance, now)
    this.sustained.Init(sustained, bursttolerance, now)
}

/*******************************************************************************
 * ALLOCATORS
 ******************************************************************************/

// New allocates initialize a gcra, setting its traffic contract parameters:
// peak is the peak increment, jittertolerance is the peak limit, sustained is
// the sustained increment, and bursttolerance is the sustained limit.
func New(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, bursttolerance ticks.Ticks, now ticks.Ticks) * Contract {
    contract := new(Contract)
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

// Commits updates the gcra with the number of events having been emitted
// starting at the time specified in the previous Request, and returns false
// if the gcra is alarmed, indicating the application might want to slow it
// down a bit, true otherwise.
func (this * Contract) Commits(events throttle.Events) bool {
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
func (this * Contract) Admits(now ticks.Ticks, events throttle.Events) bool {
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
// update the gcra with the current time, with no event emission. This
// marks the passage of time during which the emission stream is idle, which
// may bring the gcra back into compliance with the traffic contract.
func (this * Contract) Update(now ticks.Ticks) bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Update(now)
    s := this.sustained.Update(now)
    return (p && s)
}

/*******************************************************************************
 * GETTERS
 ******************************************************************************/

// Expected returns the number of ticks that would be necessary for the
// caller to delay for the event stream  to comply to the traffic contract with
// no limit penalty accumulated given the current state of the gcra.
func (this * Contract) Expected() ticks.Ticks {
    var delay ticks.Ticks = 0

    p := this.peak.Expected()
    s := this.sustained.Expected()
    if (p > s) {
        delay = p
    } else {
        delay = s
    }

    return delay
}

// isEmpty returns true if both gcras are empty, that is, neother has accumulated
// early ticks.
func (this * Contract) IsEmpty() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.IsEmpty()
    s := this.sustained.IsEmpty()
    return (p && s)
}

// IsFull returns true if either gcra is full, that is, its accumulated early
// ticks is greater than or equal to its limit.
func (this * Contract) IsFull() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.IsFull()
    s := this.sustained.IsFull()
    return (p || s)
}

// IsAlarmed returns true if either gcra is alarmed, that is, its accumulated
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

// Emptied returns true if either gcra just emptied in the last action.
func (this * Contract) Emptied() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Emptied()
    s := this.sustained.Emptied()
    return (p || s)
}

// Filled returns true if either gcra just filled in the last action.
func (this * Contract) Filled() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Filled()
    s := this.sustained.Filled()
    return (p || s)
}

// Alarmed returns true if either gcra just alarmed in the last action.
func (this * Contract) Alarmed() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Alarmed()
    s := this.sustained.Alarmed()
    return (p || s)
}

// Cleared returns true if either gcra just unalarmed in the last action,
// indicating that the event emission stream has returned to being compliant
// with the traffic contract.
func (this * Contract) Cleared() bool {
    // Both calls must be executed! Beware refactoring!
    p := this.peak.Cleared()
    s := this.sustained.Cleared()
    return (p || s)
}
