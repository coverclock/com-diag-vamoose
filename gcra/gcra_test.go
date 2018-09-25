/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
package gcra

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
)

func TestIncrement(t * testing.T) {
    var i ticks.Ticks
    
    // func Increment(numerator Events, denominator Events, frequency ticks.Ticks) ticks.Ticks
       
    i = Increment(2, 1, 4) // 2/s @ 4Hz = 2t
	if (i == ticks.Ticks(2)) {} else { t.Errorf("FAILED! i=%v", i) }
    
    i = Increment(1, 2, 4) // 0.5/s @ 4Hz = 8t
	if (i == ticks.Ticks(8)) {} else { t.Errorf("FAILED! i=%v", i) }

    i = Increment(2, 1, 5) // 2/s @ 5Hz = 3t
	if (i == ticks.Ticks(3)) {} else { t.Errorf("FAILED! i=%v", i) }
    
    i = Increment(1, 2, 5) // 0.5/s @ 5Hz = 10t
	if (i == ticks.Ticks(10)) {} else { t.Errorf("FAILED! i=%v", i) }
	
}

func TestJitterTolerance(t * testing.T) {
    var jt ticks.Ticks
    
    // func JitterTolerance(peak ticks.Ticks, burstsize Events) ticks.Ticks

    jt = JitterTolerance(2, 3) // Nominal
	if (jt == ticks.Ticks(4)) {} else { t.Errorf("FAILED! jt=%v", jt) }

    jt = JitterTolerance(2, 0) // MBS <= 1
	if (jt == ticks.Ticks(0)) {} else { t.Errorf("FAILED! jt=%v", jt) }

    jt = JitterTolerance(2, 1) // MSB <= 1
	if (jt == ticks.Ticks(0)) {} else { t.Errorf("FAILED! jt=%v", jt) }

    jt = JitterTolerance(3, 2) // Nominal
	if (jt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! jt=%v", jt) }
}

func TestBurstTolerance(t * testing.T) {
    var bt ticks.Ticks
    
    // func BurstTolerance(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, burstsize Events) ticks.Ticks

    bt = BurstTolerance(2, 3, 7, 5) // Nominal
	if (bt == ticks.Ticks(23)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 0, 7, 5) // No CDVT
	if (bt == ticks.Ticks(20)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(0, 0, 2, 5) // No PCR, CDVT
	if (bt == ticks.Ticks(8)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(7, 3, 7, 5) // PCR == SCR
	if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 3, 7, 1) // MBS <= 1
	if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

    bt = BurstTolerance(2, 3, 7, 0) // MBS <= 1
	if (bt == ticks.Ticks(3)) {} else { t.Errorf("FAILED! bt=%v", bt) }

}
