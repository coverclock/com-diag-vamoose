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

func TestThrottleSanity(t * testing.T) {
	const increment ticks.Ticks = 10
	const limit ticks.Ticks = 2
	var now ticks.Ticks = 0
	var delay ticks.Ticks = 0
	var elapsed ticks.Ticks = 0
	var alarmed bool = false

	that := New(increment, limit, ticks.Ticks(0))
	
    stuff := that.String()
    if (len(stuff) > 0) { t.Log(stuff) } else { t.Errorf("FAILED! \"%v\"", stuff) }

	elapsed = that.now - that.then
	if (elapsed == increment) {} else { t.Errorf("FAILED! %v", elapsed) }

	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	now = 0
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 1)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 1) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 2)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 2) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 1) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 3)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 3) {} else { t.Errorf("FAILED! %v", that.early) }
	if (that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 1) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 3)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 3) {} else { t.Errorf("FAILED! %v", that.early) }
	if (that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 2)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 2) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
  	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 1)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 1) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.early == 0) {} else { t.Errorf("FAILED! %v", that.early) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }

}

func TestThrottleOne(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
     
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // BURST
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // FILL
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // RECOVER
    
    now = now + increment + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + increment - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 2 
    if (!that.Admit(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // UPDATE
    
    now = now + increment + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED AGAIN
        
    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

}

func TestThrottleFixed(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	var size Events = 10;
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // BURST
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // FILL
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // RECOVER
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 2 
    if (!that.Admits(now, size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // UPDATE
    
    now = now + (increment * ticks.Ticks(size)) + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED AGAIN
        
    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

}

func TestThrottleVariable(t * testing.T) {
	var increment ticks.Ticks = 100
	var limit ticks.Ticks = 10
	var now ticks.Ticks = 0
	var size Events = 10;
	const BLOCKSIZE int64 = 32768
	
	// CONSTRUCTORS

	that := New(increment, limit, ticks.Ticks(0))

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED
    
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // BURST
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // FILL
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // RECOVER
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, RE-REQUEST, COMMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 1
    if (that.Request(now) == 1) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + 1
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // REQUEST, DELAY, ADMIT
    
    now = now + (increment * ticks.Ticks(size)) - 2
    if (that.Request(now) == 2) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + 2 
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Admits(now, size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (!that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // UPDATE
    
    now = now + (increment * ticks.Ticks(size)) + 10
    if (!that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    // SUSTAINED AGAIN
        
    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

}

func TestThrottleSimulated(t * testing.T) {
    const BANDWIDTH ticks.Ticks = 1024 // Bytes per second.
	const BLOCKSIZE Events = 32768
    const OPERATIONS uint = 1000000
	const MARGIN ticks.Ticks = 200 // 0.5%
	const LIMIT ticks.Ticks = 0
	var frequency ticks.Ticks = 0
	var increment ticks.Ticks = 0
	var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var seconds ticks.Ticks = 0
    var bandwidth ticks.Ticks = 0
    var delta ticks.Ticks = 0
    var margin ticks.Ticks = 0
	var size Events = 0
    var blocksize Events = 0
    var total uint64 = 0
    var inadmissable bool = false
    var iops uint = 0
    
    frequency = ticks.Frequency()
    increment = (frequency + BANDWIDTH - 1) / BANDWIDTH
    blocksize = BLOCKSIZE / 2
    seconds = (increment * ticks.Ticks(blocksize)) / frequency
    t.Logf("OPERATIONS=%d BANDWIDTH=%dB/s BLOCKSIZE=%dB mean=%dB/io frequency=%dHz\n", OPERATIONS, BANDWIDTH, BLOCKSIZE, blocksize, frequency)

    t.Logf("increment=%dt mean=%ds/io LIMIT=%dt now=%dt\n", increment, seconds, LIMIT, now)
	that := New(increment, LIMIT, now)
	t.Log(that.String())
	
	for iops = 0; iops < OPERATIONS; iops += 1 {
	    delay = that.Request(now)
	    now += delay
	    if (now >= 0) {} else { t.Errorf("OVERFLOW! %v", now) }
	    duration += delay
	    if (duration >= 0) {} else { t.Errorf("OVERFLOW! %v", duration) }
	    delay = that.Request(now)
	    if (delay <= 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }
        size = Events(rand.Int63n(int64(BLOCKSIZE))) + 1
	    if (0 < size) {} else { t.Errorf("FAILED! %v", size) }
	    if (size <= Events(BLOCKSIZE)) {} else { t.Errorf("FAILED! %v", size) }
	    total += uint64(size)
	    if (total > 0) {} else { t.Errorf("OVERFLOW! %v", total) }
	    inadmissable = that.Commits(size)
	    if (!inadmissable) {} else { t.Errorf("FAILED! %v", inadmissable); t.Log(that.String()) }
	}
	
	blocksize = Events(total / uint64(OPERATIONS))
	seconds = duration / frequency
	delay = seconds / ticks.Ticks(OPERATIONS)
	t.Logf("total=%dB mean=%dB/io duration=%dt=%ds mean=%ds/io\n", total, blocksize, duration, seconds, delay)
	if (duration > frequency) {} else { t.Errorf("FAILED! %v", duration) }

	bandwidth = ticks.Ticks(total) / seconds
	delta = bandwidth - BANDWIDTH
	if (delta < 0) { delta = -delta }
    margin = BANDWIDTH / MARGIN
	t.Logf("bandwidth=%dB/s delta=%dB/s margin=%dB/s\n", bandwidth, delta, margin)
	if (delta < margin) {} else { t.Errorf("FAILED! %v", delta) }
    
}

