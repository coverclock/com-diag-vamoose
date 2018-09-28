/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Computes the Fletcher-16 checksum of the data stream read from standard
// input and emits it to standard output.
//
// USAGE
//
// EXAMPLES
//
// dd if=/dev/urandom count=10 | ./fletch -V -b 512 | ./shape -V -p 2048 -s 1024 -b 512 | ./fletch -V -b 512 > /dev/null
//
package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/coverclock/com-diag-vamoose/ticks"
    "github.com/coverclock/com-diag-vamoose/fletcher"
)

const APP_VERSION = "0.0"

//                                      "h",        "Print the help menu."
var versionFlag     * bool  = flag.Bool("v", false, "Print the version number.")
var debugFlag       * bool  = flag.Bool("D", false, "Enable debug output.")
var verboseFlag     * bool  = flag.Bool("V", false, "Enable verbose output.")
var burstFlag       * int   = flag.Int("b", 1, "Set the buffer size in bytes.")

func main() {
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0
    
    flag.Parse()

    if *versionFlag {
        fmt.Fprintf(os.Stderr, "Version: %s.\n", APP_VERSION)
    }
 
    var frequency = ticks.Frequency()
    var now = ticks.Now()
    var buffer = make([] byte, int(*burstFlag))
    var read int = 0
    var written int = 0
    var total int64 = 0
    var count int64 = 0
    var eof error
    var before ticks.Ticks = 0
    var after ticks.Ticks = 0
    var then ticks.Ticks = 0
    var rate float64 = 0.0
    var peak float64 = 0.0
    
    before = ticks.Now()

    for {
        
        read, eof = os.Stdin.Read(buffer)
        if eof != nil {
            if *debugFlag { fmt.Fprintf(os.Stderr, "Read: EOF.\n") }
            break
        }
        if *debugFlag { fmt.Fprintf(os.Stderr, "Read: %vB.\n", read) }
        
        c = fletcher.Checksum16(buffer[:read], &a, &b)
        
        then = now
        now = ticks.Now()

        if read <= 0 {
            // Should never happen.
        } else if count <= 0 {
            // Do nothing.
        } else if now <= then {
            // Should never happen.
        } else {
            rate = float64(read) * float64(frequency) / float64(now - then)
            if rate > peak {
                peak = rate
            }
        }
                
        written, eof = os.Stdout.Write(buffer[:read])
        if eof != nil {
            fmt.Fprintf(os.Stderr, "Error: %v!\n", eof)
            break
        }
        if *debugFlag { fmt.Fprintf(os.Stderr, "Written: %vB.\n", written) }
        if written != read {
            fmt.Fprintf(os.Stderr, "Short: %v:%v!\n", read, written)
        }
                
        total += int64(written)
        count += 1

    }
    
    after = ticks.Now()
    
    if *verboseFlag {
        fmt.Fprintf(os.Stderr, "Total: %vB.\n", total)
        fmt.Fprintf(os.Stderr, "Average: %vB/io.\n", float64(total) / float64(count))
        fmt.Fprintf(os.Stderr, "Peak: %vBps.\n", peak)
        fmt.Fprintf(os.Stderr, "Sustained: %vBps.\n", float64(total) * float64(frequency) / float64(after - before))
        fmt.Fprintf(os.Stderr, "Checksum: %#04x.\n", c)
    }

}

