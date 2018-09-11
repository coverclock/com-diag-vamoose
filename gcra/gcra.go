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
// Generic Cell Rate Algorithm. A typical implementation of a GCRA is in fact
// a composite of two GCRAs: one that describes the peak rate, and one that
// describes the sustainable rate; the event stream must conform to both GCRAs.
// Yet the interface still appears to be a single GCRA from the point of view
// of the calling application.

import (
	"github.com/coverclock/com-diag-vamoose/ticks"
)

// Events is the type used to indicate how many events have been emitted since
// the last update of the throttle. An event can be the emission of a single
// packet, or a single byte, or a single bit, etc. It is up to the application
// to define what an event is. The throttle is defined in terms of ticks per
// event.
type Events uint32

// Gcra is the type that describes the interface to any implementation of the
// Generic Cell Rate Algorithm.
type Gcra interface {

    String() string

    Reset(now ticks.Ticks)

    Init(increment ticks.Ticks, limit ticks.Ticks, now ticks.Ticks)

    Fini()
    
    IsEmpty() bool

    IsFull() bool

    IsAlarmed() bool

    Emptied() bool

    Filled() bool
 
    Alarmed() bool

    Cleared() bool

    Request(now ticks.Ticks) ticks.Ticks

    Commits(events Events) bool

    Commit() bool

    Admits(now ticks.Ticks, events Events) bool

    Admit(now ticks.Ticks) bool

    Update(now ticks.Ticks) bool

}
