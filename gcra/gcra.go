/* vi: set ts=4 expandtab shiftwidth=4: */

package gcra

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// This is an interface that describes the API for any implementation of a
// Generic Cell Rate Algorithm (GCRA) as defined in the ATM Forum TM 4.1
// specification. While TM 4.1 was all about ATM traffic management, the GCRA,
// and its implementation here, can be used for rate control, traffic shaping,
// and traffic policing, of any kind of event stream: bytes, packets, button
// presses in a model-view-controller pattern, etc.
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

import (
	"github.com/coverclock/com-diag-vamoose/ticks"
)

// Events is the type used to indicate how many events have been emitted since
// the last update of the Gcra. An event can be the emission of a single
// packet, or a single byte, or a single bit, etc. It is up to the application
// to define what an event is. The Gcra is defined in terms of ticks per
// event.
type Events uint32

// Gcra is the type that describes the interface to any implementation of the
// Generic Cell Rate Algorithm.
type Gcra interface {

    // String returns a printable string showing the guts of the Gcra.
    String() string

    // Reset a Gcra back to its initial state. This is used during construction,
    // but can also be used by an application when a calamitous happenstance
    // occurs, like the far end disconnecting and reconnecting.
    Reset(now ticks.Ticks)

    // Fini handles any cleanup necessary before a Gcra is deallocated. It is
    // deferred when the Gcra is constructed by New. It is also callable as
    // part of the API, although doing so may render the Gcra unusable.
    Fini()
    
    // GetDeficit returns the number of ticks it would be necessary for the
    // caller to delay for the event stream  to comply to the traffic contract
    // with no limit penalty accumulated.
    GetDeficit() ticks.Ticks
    
    // isEmpty returns true if the Gcra is empty, that is, it has no accumulated
    // deficit ticks.
    IsEmpty() bool

    // IsFull returns true if the Gcra is full, that is, its accumulated deficit
    // ticks is greater than or equal to its limit.
    IsFull() bool

    // IsAlarmed returns true if the Gcra is alarmed, that is, its accumulated
    // deficit ticks is greater than its limit, indicating that the event
    // emission stream is out of compliance with the traffic contract.
    IsAlarmed() bool

    // Emptied is true if the Gcra just emptied in the last action.
    Emptied() bool

    // Filled is true if the Gcra just filled in the last action.
    Filled() bool
 
    // Alarmed is true if the Gcra just alarmed in the last action.
    Alarmed() bool

    // Cleared is true if the Gcra just unalarmed in the last action, indicating
    // that the event emission stream has returned to being compliant with the
    // traffic contract.
    Cleared() bool

    // Request asks given the current time in ticks how long of a delay in ticks
    // would be necessary before the next event were emitted for that emission
    // to be in compliance with the traffic contract.
    Request(now ticks.Ticks) ticks.Ticks

    // Commits updates the Gcra with the number of events having been emitted
    // starting at the time specified in the previous Request, and returns false
    // if the Gcra is alarmed, indicating the application might want to slow it
    // down a bit, true otherwise.
    Commits(events Events) bool

    // Commit is equivalent to calling Commits with one event.
    Commit() bool

    // Admits combines calling Request with the current time in ticks with
    // calling and returning the value of Commits with the number of events.
    Admits(now ticks.Ticks, events Events) bool

    // Admit is equivalent to calling Admits with one event.
    Admit(now ticks.Ticks) bool

    // Update is equivalent to calling Admits with zero events. It is a way to
    // update the Gcra with the current time, with no event emission. This
    // marks the passage of time during which the emission stream is idle, which
    // may bring the Gcra back into compliance with the traffic contract (and
    // will do so if time has advanced at least as much as the value returned by
    // GetDeficit).
    Update(now ticks.Ticks) bool

}
