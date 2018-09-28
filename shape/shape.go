/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// USAGE
//
// shape [ -h ] [ -v ] [ -D ] [ -V ] [ -p PEAKBYTESPERSECOND ] [ -s SUSTAINEDBYTESPERSECOND ] [ -b BURSTBYTES ]
//
// EXAMPLES
//
package main

import (
    "flag"
    "fmt"
    "os"
    "io"
    "github.com/coverclock/com-diag-vamoose/ticks"
    "github.com/coverclock/com-diag-vamoose/gcra"
    "github.com/coverclock/com-diag-vamoose/contract"
)

const APP_VERSION = "0.0"

//                                      "h",        "Print the help menu."
var versionFlag     * bool  = flag.Bool("v", false, "Print the version number.")
var debugFlag       * bool  = flag.Bool("D", false, "Print the version number.")
var verboseFlag     * bool  = flag.Bool("V", false, "Print the version number.")
var peakFlag        * int64 = flag.Int64("p", 1, "Set the peak rate in bytes per second.")
var sustainedFlag   * int64 = flag.Int64("s", 1, "Set the sustained rate in bytes per second.")
var burstFlag       * int64 = flag.Int64("b", 1, "Set the maximum burst size in bytes.")

func main() {

    flag.Parse()

    if *versionFlag {
        fmt.Fprintf(os.Stderr, "Version: %s.\n", APP_VERSION)
    }

    frequency := ticks.Frequency()

    peakrate := gcra.Events(*peakFlag)
    peakincrement := gcra.Increment(peakrate, 1, frequency)
    burstsize := gcra.Events(*burstFlag)
    jittertolerance := gcra.JitterTolerance(peakincrement, burstsize)
    sustainedrate := gcra.Events(*sustainedFlag)
    sustainedincrement := gcra.Increment(sustainedrate, 1, frequency)
    bursttolerance := gcra.BurstTolerance(peakincrement, jittertolerance, sustainedincrement, burstsize)

    now := ticks.Now()

    shape := contract.New(peakincrement, 0, sustainedincrement, bursttolerance, now)

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

    for {
        
        read, eof = io.ReadAtLeast(os.Stdin, buffer[:], 1)
        if eof != nil {
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
        
        now = ticks.Now()
        admissable = shape.Admits(now, gcra.Events(written))
        if !admissable { fmt.Fprintf(os.Stderr, "Admissable: %v!\n", admissable) }
        
        total += int64(written)
        count += 1

    }
    
    now = ticks.Now()
    shape.Update(now)
    delay = shape.Comply()
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
        fmt.Fprintf(os.Stderr, "Sustained: %vBps.\n", float64(total) * float64(frequency) / float64(after - before))
    }

}

