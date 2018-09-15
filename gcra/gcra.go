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
