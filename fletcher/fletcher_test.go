/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
package fletcher

import (
    "testing"
)

func TestChecksum(t *testing.T) {
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0
    var buffer [6] byte = [...] byte { 'q', 'w', 'e', 'r', 't', 'y' }
    
    c = Checksum16(buffer[:], &a, &b)
    
    if (a == uint8(0xae)) {} else { t.Errorf("a=0x%x\n", a) }
    if (b == uint8(0x4d)) {} else { t.Errorf("b=0x%x\n", b) }
    if (c == uint16(0x4dae)) {} else { t.Errorf("c=0x%x\n", c) }
}
