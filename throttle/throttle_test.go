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
	var now ticks.Ticks = 0
	var stuff string = ""
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error() }
    if (!that.IsFull()) {} else { t.Error() }
    if (!that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (!that.Filled()) {} else { t.Error() }
    if (!that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error() }
    if (!that.IsFull()) {} else { t.Error() }
    if (!that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (!that.Filled()) {} else { t.Error() }
    if (!that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
    // BURST
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
      
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)
 
    if (!that.IsEmpty()) {} else { t.Error() }
    if (!that.IsFull()) {} else { t.Error() }
    if (!that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (!that.Filled()) {} else { t.Error() }
    if (!that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
    // FILL
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error() }
    if (that.Commit()) {} else { t.Error() }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error() }
    if (that.IsFull()) {} else { t.Error() }
    if (that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (that.Filled()) {} else { t.Error() }
    if (that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
    // RECOVER
    
    now = now + increment + 1
    if (that.Request(now) == 1) {} else { t.Error() }
    if (that.Commit()) {} else { t.Error() }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error() }
    if (that.IsFull()) {} else { t.Error() }
    if (that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (!that.Filled()) {} else { t.Error() }
    if (!that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
    now = now + increment + 1
    if (that.Request(now) == 0) {} else { t.Error() }
    if (!that.Commit()) {} else { t.Error() }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error() }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error() }
    if (!that.IsFull()) {} else { t.Error() }
    if (that.IsAlarmed()) {} else { t.Error() }
    if (!that.Emptied()) {} else { t.Error() }
    if (!that.Filled()) {} else { t.Error() }
    if (!that.Alarmed()) {} else { t.Error() }
    if (!that.Cleared()) {} else { t.Error() }
    
}

