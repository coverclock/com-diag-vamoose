package throttle

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
 	"github.com/coverclock/com-diag-vamoose/framework"
    "math/rand"
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
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }
 	
 	now = 0
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 1)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 1) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 2)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 2) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment - 1
 	delay = that.Request(now)
 	if (delay == 1) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 3)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 3) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 1) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 3)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 3) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 2)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 2) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
  	if (alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == (increment + 1)) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 1) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (!that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (!that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment + 1
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
 	
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (!that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

    now = now + increment
 	delay = that.Request(now)
 	if (delay == 0) {} else { t.Errorf("FAILED! %v", delay) }
 	
 	alarmed = !that.Commit()
 	if (!alarmed) {} else { t.Errorf("FAILED! %v", alarmed) }
    
	if (that.increment == increment) {} else { t.Errorf("FAILED! %v", that.increment) }
	if (that.limit == limit) {} else { t.Errorf("FAILED! %v", that.limit) }
	if (that.expected == increment) {} else { t.Errorf("FAILED! %v", that.expected) }
	if (that.deficit == 0) {} else { t.Errorf("FAILED! %v", that.deficit) }
	if (!that.full0) {} else { t.Errorf("FAILED! %v", that.full0) }
	if (!that.full1) {} else { t.Errorf("FAILED! %v", that.full1) }
	if (!that.full2) {} else { t.Errorf("FAILED! %v", that.full2) }
 	if (that.empty0) {} else { t.Errorf("FAILED! %v", that.empty0) }
  	if (that.empty1) {} else { t.Errorf("FAILED! %v", that.empty1) }
 	if (that.empty2) {} else { t.Errorf("FAILED! %v", that.empty2) }
 	if (!that.alarmed1) {} else { t.Errorf("FAILED! %v", that.alarmed1) }
 	if (!that.alarmed2) {} else { t.Errorf("FAILED! %v", that.alarmed2) }
 	
 	if (that.GetDeficit() == that.deficit) {} else { t.Errorf("FAILED! %v!=%v", that.GetDeficit(), that.deficit) }
 	if (that.IsFull() == that.full1) {} else { t.Errorf("FAILED! %v!=%v", that.IsFull(), that.full1) }
 	if (that.IsEmpty() == that.empty1) {} else { t.Errorf("FAILED! %v!=%v", that.IsEmpty(), that.empty1) }
 	if (that.IsAlarmed() == that.alarmed1) {} else { t.Errorf("FAILED! %v!=%v", that.IsAlarmed(), that.alarmed1) }

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
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
     
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
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + increment - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (!that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + increment
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commit()) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
   
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
      
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    
    now = now + (increment * ticks.Ticks(size)) - 1;
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (!that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Update(now)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

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
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))  
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size)) 
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }

    now = now + (increment * ticks.Ticks(size))
    if (that.Request(now) == 0) {} else { t.Error("FAILED!"); t.Log(that.String()) }
    size = gcra.Events(rand.Int63n(BLOCKSIZE)) + 1
    if (that.Commits(size)) {} else { t.Error("FAILED!"); t.Log(that.String()) }
 
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
	const BLOCKSIZE int = 32768
    const OPERATIONS int = 1000000
	const LIMIT ticks.Ticks = 0
     
    frequency := ticks.Frequency()
    increment := (frequency + BANDWIDTH - 1) / BANDWIDTH
    t.Logf("OPERATIONS=%v BANDWIDTH=%vB/s BLOCKSIZE=%vB frequency=%vHz\n", OPERATIONS, BANDWIDTH, BLOCKSIZE, frequency)

    interarrival := float64(increment) / float64(frequency)
    expected := float64(frequency) / float64(increment)
    now := ticks.Now()
    t.Logf("increment=%vt seconds=%vs expected=%vB/s LIMIT=%vt now=%vt\n", increment, interarrival, expected, LIMIT, now)
	that := New(increment, LIMIT, now)
	
	framework.SimulatedEventStream(t, that, BLOCKSIZE, OPERATIONS)
    
}

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

func TestThrottleActual(t * testing.T) {
    const BURST int = 64				// bytes
    const BANDWIDTH int = 1024			// bytes per second
    const TOTAL uint64 = 1024 * 60		// bytes

    supply := make(chan byte, BURST + 1) // Producer closes.
    demand := make(chan byte, BURST) // Policer closes.
       
    frequency := ticks.Frequency()
    increment := frequency / ticks.Ticks(BANDWIDTH)
    limit := frequency * ticks.Ticks(BURST) / ticks.Ticks(BANDWIDTH)
    now := ticks.Now()
    
    shape := New(increment, 0, now)
    police := New(increment, limit, now)
    
    framework.ActualEventStream(t, shape, police, supply, demand, TOTAL)

}
