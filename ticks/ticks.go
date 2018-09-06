/* vi: set ts=4 expandtab shiftwidth=4: */

package ticks

/**
 * @file
 *
 * Copyright 2018 Digital Aggregates Corporation, Colorado, USA<BR>
 * Licensed under the terms in LICENSE.txt<BR>
 * Chip Overclock <coverclock@diag.com><BR>
 * https://github.com/coverclock/com-diag-vamoose<BR>
 *
 * ABSTRACT
 *
 * Basic time-related functions for use in Vamoose. We use the go Time types
 * and functions but put a thin abstraction layer around them so it will be
 * easily changed in the future if necessary. (For example, I might end up
 * using the low level runtime now() function instead.)
 */

import (
	"time"
)

/*******************************************************************************
 * TYPES
 ******************************************************************************/

type Ticks time.Duration

/*******************************************************************************
 * VALUES
 ******************************************************************************/

const (
	FREQUENCY Ticks = 1000000000
)

var epoch time.Time = time.Now()

/*******************************************************************************
 * ACTIONS
 ******************************************************************************/

func Frequency() Ticks {
	return Ticks(FREQUENCY)
}

func Now() Ticks {
	return Ticks(time.Now().Sub(epoch))
}
