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
    "math/rand"
	"github.com/coverclock/com-diag-vamoose/ticks"
)

func TestThrottle1(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	var stuff string = ""
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // BURST
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
      
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // FILL
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    if (that.Commit()) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (that.Filled()) {} else { t.Error("FAILED!") }
    if (that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // RECOVER
    
    now = now + increment + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
    if (that.Commit()) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    now = now + increment + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 2 
    if (!that.Admit(now)) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // UPDATE
    
    now = now + increment + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED AGAIN
        
    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commit()) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }

}

func TestThrottle10(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	var size Events = 10;
	var stuff string = ""
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // BURST
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // FILL
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    if (that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (that.Filled()) {} else { t.Error("FAILED!") }
    if (that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // RECOVER
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
    if (that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 2 
    if (!that.Admits(now, size)) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // UPDATE
    
    now = now + (increment * ticks.Ticks(size)) + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED AGAIN
        
    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }

}

func TestThrottleN(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	var size Events = 10;
	var stuff string = ""
	const BLOCKSIZE int64 = 32768
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // BURST
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // FILL
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (that.Filled()) {} else { t.Error("FAILED!") }
    if (that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // RECOVER
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
     
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!") }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!") }
    
    now = now + 2 
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Admits(now, size)) {} else { t.Error("FAILED!") }
 
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (!that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }
    
    // UPDATE
    
    now = now + (increment * ticks.Ticks(size)) + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)

    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (that.Cleared()) {} else { t.Error("FAILED!") }
    
    // SUSTAINED AGAIN
        
    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!") }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!") }
    
    stuff = that.String()
    if (len(stuff) > 0) {} else { t.Error("FAILED!") }
    t.Log(stuff)
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!") }
    if (!that.IsFull()) {} else { t.Error("FAILED!") }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!") }
    if (!that.Emptied()) {} else { t.Error("FAILED!") }
    if (!that.Filled()) {} else { t.Error("FAILED!") }
    if (!that.Alarmed()) {} else { t.Error("FAILED!") }
    if (!that.Cleared()) {} else { t.Error("FAILED!") }

}

func TestThrottleLoop(t * testing.T) {
    const BANDWIDTH ticks.Ticks = 1024 // Bytes per second.
    const OPERATIONS int = 1000000
	const BLOCKSIZE int64 = 32768
	const MARGIN ticks.Ticks = 200
	var frequency ticks.Ticks = 0
	var increment ticks.Ticks = 0
	var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var bandwidth ticks.Ticks = 0  
    var measured ticks.Ticks = 0
    var delta ticks.Ticks = 0
    var margin ticks.Ticks = 0
	var size Events = 10;
    var total Events = 0
    var inadmissable bool = false
    
    frequency = ticks.Frequency()
    increment = frequency / BANDWIDTH
   
	that := New(increment, ticks.Ticks(0), now)
	t.Log(that.String())
	
	for iops := 0; iops < OPERATIONS; iops += 1 {
	    delay = that.Request(now)
	    now += delay
	    duration += delay
	    delay = that.Request(now)
	    if (delay == 0) {} else { t.Error("FAILED!") }
        size = Events(rand.Int63n(BLOCKSIZE)) + 1
	    if (0 < size) {} else { t.Error("FAILED!") }
	    if (size <= Events(BLOCKSIZE)) {} else { t.Error("FAILED!") }
	    total += size
	    inadmissable = that.Commits(size)
	    if (!inadmissable) {} else { t.Error("FAILED!") }
	}
	
	if (total > 0) {} else { t.Error("FAILED!") }
	if (duration > frequency) {} else { t.Error("FAILED!") }
	bandwidth = duration / frequency
	measured = ticks.Ticks(total) / bandwidth
	delta = measured - BANDWIDTH
	if (delta < 0) { delta = -delta }
    margin = BANDWIDTH / MARGIN
	t.Logf("BANDWIDTH=%dbps total=%dbytes duration=%dticks=%dseconds measured=%dbps delta=%dbps margin=%dbps\n", BANDWIDTH, total, duration, bandwidth, measured, delta, margin)
	if (delta < margin) {} else { t.Error("FAILED!") }
    
}

