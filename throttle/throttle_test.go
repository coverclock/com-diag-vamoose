package throttle

/**
 * @file
 *
 * Copyright 2018 Digital Aggregates Corporation, Colorado, USA<BR>
 * Licensed under the terms in LICENSE.txt<BR>
 * Chip Overclock <coverclock@diag.com><BR>
 * https://github.com/coverclock/com-diag-vamoose<BR>
 */

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
)

func TestThrottle(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	
	now := ticks.Now()

	that := New(increment, limit, now)

    if (that.IsEmpty())		{} else { t.Error("IsEmpty") }
    if (!that.IsFull())		{} else { t.Error("IsFull") }
    if (!that.IsAlarmed())	{} else { t.Error("IsAlarmed") }
    if (!that.Emptied())	{} else { t.Error("Emptied") }
    if (!that.Filled())		{} else { t.Error("Filled") }
    if (!that.Alarmed())	{} else { t.Error("Alarmed") }
    if (!that.Cleared())	{} else { t.Error("Cleared") }
    
    that.Error()
    
}

