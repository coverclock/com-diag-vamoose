/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Shapes data stream read from standard input and emits it to standard output.
//
// USAGE
//
// shape [ -h ] [ -v ] [ -D ] [ -V ] [ -p PEAKBYTESPERSECOND ] [ -s SUSTAINEDBYTESPERSECOND ] [ -b BURSTBYTES ]
//
// EXAMPLES
//
// yes | head -100 | ./shape -V -p 64 -s 32 -b 32 > /dev/null
//
// dd if=/dev/zero count=10 | ./shape -V -p 2048 -s 1024 -b 512 | dd of=/dev/null
//
// dd if=/dev/urandom count=10 | ./fletch -V -b 512 | ./shape -V -p 2048 -s 1024 -b 512 | ./fletch -V -b 512 > /dev/null
//
package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/gcra"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/contract"
)

const APP_VERSION = "0.0"

//                                      "h",        "Print the help menu."
var versionFlag     * bool  = flag.Bool("v", false, "Print the version number.")
var debugFlag       * bool  = flag.Bool("D", false, "Enable debug output.")
var verboseFlag     * bool  = flag.Bool("V", false, "Enable verbose output.")
var peakFlag        * int64 = flag.Int64("p", 1, "Set the peak rate in bytes per second.")
var sustainedFlag   * int64 = flag.Int64("s", 1, "Set the sustained rate in bytes per second.")
var burstFlag       * int64 = flag.Int64("b", 1, "Set the maximum burst size in bytes.")

func main() {
    
    flag.Parse()

    if *versionFlag {
        fmt.Fprintf(os.Stderr, "Version: %s.\n", APP_VERSION)
    }

    var frequency = ticks.Frequency()
    var now = ticks.Now()

    peakrate := throttle.Events(*peakFlag)
    peakincrement := gcra.Increment(peakrate, 1, frequency)
    burstsize := throttle.Events(*burstFlag)
    jittertolerance := gcra.JitterTolerance(peakincrement, burstsize)
    sustainedrate := throttle.Events(*sustainedFlag)
    sustainedincrement := gcra.Increment(sustainedrate, 1, frequency)
    bursttolerance := contract.BurstTolerance(peakincrement, jittertolerance, sustainedincrement, burstsize)

    var shape throttle.Throttle = contract.New(peakincrement, 0, sustainedincrement, bursttolerance, now)

    if *verboseFlag {
        fmt.Fprintf(os.Stderr, "Contract: %v.\n", shape)
    }
    
    var buffer = make([] byte, int(burstsize))
    var delay ticks.Ticks = 0
    var read int = 0
    var written int = 0
    var total int64 = 0
    var count int64 = 0
    var eof error
    var admissable bool = false
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
        
        now = ticks.Now()

        delay = shape.Request(now)
        if *debugFlag {
            fmt.Fprintf(os.Stderr, "Delay: %vs.\n", float64(delay) / float64(frequency))
        }
        
        if delay > 0 { 
            ticks.Sleep(delay)
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
        
        then = now
        now = ticks.Now()
        admissable = shape.Admits(now, throttle.Events(written))
        if !admissable { fmt.Fprintf(os.Stderr, "Admissable: %v!\n", admissable) }
        
        if written <= 0 {
            // Should never happen.
        } else if count <= 0 {
            // Do nothing.
        } else if now <= then {
            // Should never happen.
        } else {
            rate = float64(written) * float64(frequency) / float64(now - then)
            if rate > peak {
                peak = rate
            }
        }
        
        total += int64(written)
        count += 1

    }
    
    now = ticks.Now()
    shape.Update(now)

    delay = shape.Expected()
    if *debugFlag {
        fmt.Fprintf(os.Stderr, "Delay: %vs.\n", float64(delay) / float64(frequency))
    }

    if delay > 0 {
        ticks.Sleep(delay)
    }
    
    after = ticks.Now()
    shape.Update(after)

    if *verboseFlag {
        fmt.Fprintf(os.Stderr, "Contract: %v.\n", shape)
        fmt.Fprintf(os.Stderr, "Total: %vB.\n", total)
        fmt.Fprintf(os.Stderr, "Average: %vB/io.\n", float64(total) / float64(count))
        fmt.Fprintf(os.Stderr, "Peak: %vBps.\n", peak)
        fmt.Fprintf(os.Stderr, "Sustained: %vBps.\n", float64(total) * float64(frequency) / float64(after - before))
    }

}

