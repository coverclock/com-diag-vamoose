/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
package throttle

import (
    "testing"
    "unsafe"
    "fmt"
)

func TestTypes(t * testing.T) {
    fmt.Printf("Events: Alignof=%v Sizeof=%v\n", unsafe.Alignof(Events(0)), unsafe.Sizeof(Events(0)));
}
