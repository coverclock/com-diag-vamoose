package gcra

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
)

func TestBurstTolerance(t *testing.T) {
    var bt ticks.Ticks = 0
    
    // BurstTolerance(peak ticks.Ticks, jittertolerance ticks.Ticks, sustained ticks.Ticks, burstsize Events)

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
