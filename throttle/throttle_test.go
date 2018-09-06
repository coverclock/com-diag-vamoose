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
)

func TestThrottle(t * testing.T) {
	var increment Ticks = 100
	var limit Ticks = 10

	that := throttle.New(increment, limit)

    if (that.IsEmpty())		{} else { t.Error("IsEmpty") }
    if (!that.IsFull())		{} else { t.Error("IsFull") }
    if (!that.IsAlarmed())	{} else { t.Error("IsAlarmed") }
    if (!that.Emptied())	{} else { t.Error("Emptied") }
    if (!that.Filled())		{} else { t.Error("Filled") }
    if (!that.Alarmed())	{} else { t.Error("Alarmed") }
    if (!that.Cleared())	{} else { t.Error("Cleared") }
    
    that.Error()
    
}

