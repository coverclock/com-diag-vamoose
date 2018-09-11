package throttle

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
    "math/rand"
    "time"
    "net"
    "fmt"
    "sync"
)

/*******************************************************************************
 * SANITY
 ******************************************************************************/

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

/*******************************************************************************
 * SINGLE EVENTS
 ******************************************************************************/

func TestThrottleOne(t * testing.T) {
	const increment ticks.Ticks = 100
	const limit ticks.Ticks = 10
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

/*******************************************************************************
 * FIXED EVENTS
 ******************************************************************************/

func TestThrottleFixed(t * testing.T) {
	const increment ticks.Ticks = 100
	const limit ticks.Ticks = 10
	const size gcra.Events = 10;
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

/*******************************************************************************
 * VARIABLE EVENTS
 ******************************************************************************/

func TestThrottleVariable(t * testing.T) {
	const BLOCKSIZE int64 = 32768
	const increment ticks.Ticks = 100
	const limit ticks.Ticks = 10
	var size gcra.Events = 0;
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
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
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
    if (that.IsEmpty()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsFull()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.IsAlarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Emptied()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Filled()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Alarmed()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (!that.Cleared()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

}

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func TestThrottleSimulated(t * testing.T) {
    const BANDWIDTH ticks.Ticks = 1024 // Bytes per second.
	const BLOCKSIZE gcra.Events = 32768
    const OPERATIONS uint = 1000000
	const MARGIN ticks.Ticks = 200 // 0.5%
	const limit ticks.Ticks = 0
	var increment ticks.Ticks = 0
	var frequency ticks.Ticks = 0
	var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var seconds ticks.Ticks = 0
    var bandwidth ticks.Ticks = 0
    var delta ticks.Ticks = 0
    var margin ticks.Ticks = 0
	var size gcra.Events = 0
    var blocksize gcra.Events = 0
    var total uint64 = 0
    var inadmissable bool = false
    var iops uint = 0
    
    frequency = ticks.Frequency()
    increment = (frequency + BANDWIDTH - 1) / BANDWIDTH
    blocksize = BLOCKSIZE / 2
    seconds = (increment * ticks.Ticks(blocksize)) / frequency
    t.Logf("OPERATIONS=%d BANDWIDTH=%dB/s BLOCKSIZE=%dB mean=%dB/io frequency=%dHz\n", OPERATIONS, BANDWIDTH, BLOCKSIZE, blocksize, frequency)

    t.Logf("increment=%dt mean=%ds/io LIMIT=%dt now=%dt\n", increment, seconds, limit, now)
	that := New(increment, limit, now)
	t.Log(that.String())
	
	for iops = 0; iops < OPERATIONS; iops += 1 {
	    delay = that.Request(now)
	    now += delay
	    if (now >= 0) {} else { t.Errorf("OVERFLOW! %v", now) }
	    duration += delay
	    if (duration >= 0) {} else { t.Errorf("OVERFLOW! %v", duration) }
	    delay = that.Request(now)
	    if (delay <= 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }
        size = gcra.Events(rand.Int63n(int64(BLOCKSIZE))) + 1
	    if (0 < size) {} else { t.Errorf("FAILED! %v", size) }
	    if (size <= gcra.Events(BLOCKSIZE)) {} else { t.Errorf("FAILED! %v", size) }
	    total += uint64(size)
	    if (total > 0) {} else { t.Errorf("OVERFLOW! %v", total) }
	    inadmissable = that.Commits(size)
	    if (!inadmissable) {} else { t.Errorf("FAILED! %v", inadmissable); t.Log(that.String()) }
	}
	
	blocksize = gcra.Events(total / uint64(OPERATIONS))
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

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

var mutex sync.Mutex

func producer(t * testing.T, limit uint64, burst int, delay time.Duration, output chan <- byte, done chan<- bool) {
    var total uint64 = 0
    var size int = 0
    var datum byte = 0
    
    mutex.Lock()
    fmt.Println("producer: begin.")
    mutex.Unlock()

    then := ticks.Now()
    
    for limit > 0 {
        
        size = rand.Intn(burst) + 1
        if uint64(size) > limit {
            size = int(limit)
        }
        
        for index := size; index > 0; index -= 1 {
            datum = byte(rand.Int31n(255))
            output <- datum
            total += 1  
        }
        
        limit -= uint64(size)
         
        mutex.Lock()
        fmt.Printf("producer: produced=%vB total=%vB remaining=%vB.\n", size, total, limit)
        mutex.Unlock()
       
        time.Sleep(delay)
        
    }
    
    close(output)
       
    now := ticks.Now()
    frequency := float64(ticks.Frequency())
    duration := float64(now - then) / frequency
    bandwidth := float64(total) / duration
    
    mutex.Lock()
    fmt.Printf("producer: end total=%vB duration=%vs bandwidth=%vB/s.\n", total, duration, bandwidth);
    mutex.Unlock()
    
    done <- true
}

func shaper(t * testing.T, burst int, input <- chan byte, that gcra.Gcra, output net.PacketConn, address net.Addr, done chan<- bool) {
    var total uint64 = 0
    var okay bool = true
    var size int = 0
    var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var alarmed bool = false
        
    mutex.Lock()
    fmt.Println("shaper: begin.");
    mutex.Unlock()

    buffer := make([] byte, burst)
    
    frequency := float64(ticks.Frequency())
    
    for {

        buffer[0], okay = <- input
        if !okay {
            break
        }
        total += 1

        for size = 1; (size < burst) && (len(input) > 0); size +=1 {
            buffer[size], okay = <- input
            if !okay {
                // Should never happen.
                break
            }
            total += 1
        }
        
        now = ticks.Now()
        delay = that.Request(now)
        duration = 0
        for delay > 0 {
            duration += delay                       
            time.Sleep(time.Duration(delay))
            now = ticks.Now()
            delay = that.Request(now) 
        }

        written, failure := output.WriteTo(buffer[0:size - 1], address)
        if failure != nil {
            mutex.Lock()
            fmt.Printf("shaper: failure=%v!\n", failure);
            mutex.Unlock()
            break
        }
        
        fmt.Printf("shaper: delay=%vs written=%vB total=%vB.\n", float64(duration) / frequency, written, total);

        alarmed = that.Commits(gcra.Events(size))
        if alarmed {
            mutex.Lock()
            fmt.Printf("shaper: alarmed=%v!\n", alarmed);
            mutex.Unlock()
            break
        }

    }
          
    mutex.Lock()
    fmt.Println("shaper: end");
    mutex.Unlock()
    
    done <- true
}

func policer(t * testing.T, burst int, input net.PacketConn, that gcra.Gcra, output chan<- byte, done chan<- bool) {
    var total uint64 = 0
    var now ticks.Ticks = 0
    var admissable bool = false
    
    mutex.Lock()
    fmt.Println("policer: begin.");
    mutex.Unlock()
   
    buffer := make([] byte, burst)
    
    for {
    
        read, _, failure := input.ReadFrom(buffer)
        if failure != nil {
            mutex.Lock()
            fmt.Printf("policer: failure=%v!\n", failure);
            mutex.Unlock()
            break
        }
        total += uint64(read)

        now = ticks.Now()
        admissable = that.Admits(now, gcra.Events(read))
        if admissable {
            mutex.Lock()
            fmt.Printf("policer: admitted=%vB total=%vB.\n", read, total)
            mutex.Unlock()
            for index := 0; index < read; index += 1 {
                output <- buffer[index]
            }
        } else {
            mutex.Lock()
            fmt.Printf("policer: policed=%vB total=%vB?\n", read, total);         
            mutex.Unlock()
        }
    
    }
    
    close(output)
    
    mutex.Lock()
    fmt.Println("policer: end");
    mutex.Unlock()
    
    done <- true
}

func consumer(t * testing.T, input <-chan byte, done chan<- bool) {
    var total uint64 = 0
    var okay bool = true
    
    mutex.Lock()
    fmt.Println("consumer: begin.");
    mutex.Unlock()
    
    then := ticks.Now()
    
    for {

        _, okay = <- input
        if !okay {
            break
        }
        total += 1

    }
     
    now := ticks.Now()
    frequency := float64(ticks.Frequency())
    duration := float64(now - then) / frequency
    bandwidth := float64(total) / duration

    mutex.Lock()
    fmt.Printf("consumer: end total=%vB duration=%vs bandwidth=%vB/s.\n", total, duration, bandwidth);
    mutex.Unlock()
    
    done <- true
}

func TestThrottleActual(t * testing.T) {
    const BURST int = 64				// bytes
    const BANDWIDTH int = 1024			// bytes per second
    const DURATION int = 1				// seconds
    const DELAY time.Duration = 1000	// nanoseconds
    var failure error
    
    mutex.Lock()
    fmt.Println("Beginning.")
    mutex.Unlock()
    
    done := make(chan bool, 4)
    defer close(done)
    
    supply := make(chan byte, BURST)
    defer close(supply)
    
    demand := make(chan byte, BURST)
    defer close(demand)
        
    source, failure := net.ListenPacket("udp", ":5555")
    if failure != nil {
        t.Fatal(failure)
    }
    defer source.Close()
           
    sink, failure := net.ListenPacket("udp", ":0")
    if failure != nil {
        t.Fatal(failure)
    }
    defer sink.Close()
 
    destination, failure := net.ResolveUDPAddr("udp", "localhost:5556")
    if failure != nil {
        t.Fatal(failure)
    }
    
    frequency := ticks.Frequency()
    increment := frequency / ticks.Ticks(BANDWIDTH)
    limit := frequency * ticks.Ticks(BURST) / ticks.Ticks(BANDWIDTH)
    now := ticks.Now()
    shape := New(increment, 0, now)
    police := New(increment, limit, now)
    
    mutex.Lock()
    fmt.Println("Starting.")
    mutex.Unlock()
   
    go consumer(t, demand, done)
    go policer(t, BURST, source, police, demand, done)
    go shaper(t, BURST, supply, shape, sink, destination, done)
    go producer(t, uint64(DURATION) * uint64(BANDWIDTH), BURST, DELAY, supply, done)
    
    mutex.Lock()
    fmt.Println("Waiting.")
    mutex.Unlock()
    
    if false {
        time.Sleep(time.Duration(DURATION) * 1000000000 * 5)
    } else {
        <- done
        <- done
        <- done
        <- done
    }
   
    mutex.Lock()
    fmt.Println("Ending.")
    mutex.Unlock()
}
