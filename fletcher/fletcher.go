/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Implements the computationally simple Fletcher checksum algorithm.
//
// REFERENCES
//
// J. Zweig, C. Partridge, "TCP Alternate Checksum Options", RFC 1146,
// https://tools.ietf.org/html/rfc1146, IETF, February 1990
//
// "Fletcher's checksum", Wikipedia,
// https://en.wikipedia.org/wiki/Fletcher's_checksum, 2016-12-21
//
// J. Fletcher, "An Arithmetic Checksum for Serial Transmissions",
// IEEE Transactions on Communication, COM-30, No. 1, pp. 247-252,
// January 1982
//
package fletcher

// Checksum16 computes a running sixteen-bit Fletcher checksum based on
// a slice of a byte buffer and the two eight-bit running checksum variables.
// The current sixteen-bit checksum is returned by concatenating the two
// eight-bit running values.
func Checksum16(buffer [] byte, ap * uint8, bp * uint8) uint16 {
    var a uint16
    var b uint16
    var c uint16
    
    a = uint16(*ap)
    b = uint16(*bp)
    
    for _, bb := range buffer {
        a = (a + uint16(bb)) % 255
        b = (b + a) % 255
    }
    
    *ap = uint8(a)
    *bp = uint8(b)
    
    c = b
    c <<= 8
    c |= a
    
    return c
}
