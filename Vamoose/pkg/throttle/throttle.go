/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// This is an interface that describes the API for any implementation of a
// Throttle. Throttles are mechanisms that shape event emission rates or
// police event admission rates. Frequently, throttles are implemented using
// a virtual scheduler or a leaky bucket.
//
package throttle

import (
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

// Events is the type used to indicate how many events have been emitted since
// the last update of the throttle. An event can be the emission of a single
// packet, or a single byte, or a single bit, etc. It is up to the application
// to define what an event is. The throttle is defined in terms of ticks per
// event.
type Events uint64

// Throttle is the type that describes the interface to any implementation of the
// Generic Cell Rate Algorithm.
type Throttle interface {

    /***************************************************************************
     * HELPERS
     **************************************************************************/

    // String returns a printable string showing the guts of the throttle.
    String() string

    /***************************************************************************
     * SETTERS
     **************************************************************************/

    // Reset a throttle back to its initial state. This is used during construction,
    // but can also be used by an application when a calamitous happenstance
    // occurs, like the far end disconnecting and reconnecting.
    Reset(now ticks.Ticks)

    /***************************************************************************
     * MUTATORS
     **************************************************************************/

    // Request computes, given the current time in ticks, how long of a delay
    // in ticks would be necessary before the next event were emitted for that
    // emission to be in compliance with the traffic contract.
    Request(now ticks.Ticks) ticks.Ticks

    // Commits updates the throttle with the number of events having been emitted
    // starting at the time specified in the previous Request, and returns false
    // if the throttle is alarmed, indicating the application might want to slow it
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
    // update the throttle with the current time, with no event emission. This
    // marks the passage of time during which the emission stream is idle, which
    // may bring the throttle back into compliance with the traffic contract (and
    // will do so if time has advanced at least as much as the value returned by
    // Expected).
    Update(now ticks.Ticks) bool

    /***************************************************************************
     * GETTERS
     **************************************************************************/
    
    // Expected returns the number of ticks that would be necessary for the
    // caller to delay for the event stream  to comply to the traffic contract with
    // no limit penalty accumulated given the current state of the throttle. For
    // throttles whose implementations differ from that of the Generic Cell Rate
    // Algorithm, the value returned may be the same as that returned by Request
    // given the current state of the throttle, or some other value entirely.
    Expected() ticks.Ticks
    
    // isEmpty returns true if the throttle is empty, that is, it has no accumulated
    // deficit ticks.
    IsEmpty() bool

    // IsFull returns true if the throttle is full, that is, its accumulated deficit
    // ticks is greater than or equal to its limit.
    IsFull() bool

    // IsAlarmed returns true if the throttle is alarmed, that is, its accumulated
    // deficit ticks is greater than its limit, indicating that the event
    // emission stream is out of compliance with the traffic contract.
    IsAlarmed() bool

    /***************************************************************************
     * SENSORS
     **************************************************************************/

    // Emptied returns true if the throttle just emptied in the last action.
    Emptied() bool

    // Filled returns true if the throttle just filled in the last action.
    Filled() bool
 
    // Alarmed returns true if the throttle just alarmed in the last action.
    Alarmed() bool

    // Cleared returns true if the throttle just unalarmed in the last action,
    // indicating that the event emission stream has returned to being
    // compliant with the traffic contract.
    Cleared() bool

}
