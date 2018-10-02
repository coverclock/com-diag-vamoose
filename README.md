# com-diag-vamoose

Musings with the Go programming language.

## Copyright

Copyright 2018 by the Digital Aggregates Corporation.

## License

Licensed under the terms of the Lesser GNU Public License version 2.1.

## Trademarks

"Digital Aggregates Corporation" is a registered trademark.

"Chip Overclock" is a registered trademark.

## Contact

Chip Overclock    
<mailto:coverclock@diag.com>    
Digital Aggregates Corporation    
<http://wwww.diag.com>    
3440 Youngfield St. #209    
Wheat Ridge CO 80033    

## Abstract

This repository contains the results of my attempts to learn the Go
programming language (a.k.a. golang) by implementing some non-trivial
and possible useful packages. One of these packages include yet another
implementation on my part of the Generic Cell Rate Algorithm (GCRA),
adapted as usual to bursts of variable length packets.

## Packages

* com-diag-vamoose/Vamoose/pkg/contract - Implements a traffic contract throttle consisting of peak and sustained GCRAs.
* com-diag-vamoose/Vamoose/pkg/fletcher - Implements the Fletcher sixteen-bit checksum algorithm.
* com-diag-vamoose/Vamoose/pkg/gcra - Implements a Generic Cell Rate Algorithm (GCRA) throttle using a virtual scheduler.
* com-diag-vamoose/Vamoose/pkg/harness - Provides at test harness for exercising throttles.
* com-diag-vamoose/Vamoose/pkg/throttle - Describes the interface for a rate control algorithm.
* com-diag-vamoose/Vamoose/pkg/ticks - Implements basic monotonic time functions for use in rate control.

## Commands

* com-diag-vamoose/Vamoose/cmd/fletch - Computes the Fletcher-16 checksum of a data stream admitted from standard input and emitted to standard output.
* com-diag-vamoose/Vamoose/cmd/shape - Shapes the data stream admitted from standard input and emitted to standard output.

## Remarks

My systems programming language of choice has changed over the decades,
depending on what I was getting paid to do and where I was doing it. In
the 1970s, it was IBM 360/370 Basic Assembler Language (BAL), and later
a structured language implemented in BAL macro language (which itself
was Turing complete), with an occasional foray into PL/1. In the 1980s,
it was PDP-11 Assembler Language (PAL). In the late 1980s and to the
mid-1990s it was C. In the mid-1990s to the mid-2000s it was C++, which
was mostly an artifact of the long history Bell Labs and its spinoffs
(where I was variously employed during that period) had for using C++
for firmware development. In the 2010s, I saw a significant reduction
in the use of C++ for systems programming, in part due to the evolution
of C++ into a langauge that was hard to learn, difficult to debug,
and hence not terribly productive to use.

During most of this time I cast about for an alternative to C and C++
for the kinds of real-time or close-to-bare-metal work I typically do. I
briefly considered D, but it didn't seem to catch on with the mainstream.
I used Java in two product development efforts, one of which was actually
an embedded project for which we used a Java compiler, but that was hardly
mainstream either. I've done quite a bit of development in Python, but
that was strictly in the realm of building tools to support my embedded
work. I've been known to hack JavaScript in an emergency.

Why Go? Moore’s Law, based on an observation made in 1965 by Gordon
Moore, founder of Fairchild Semiconductor and Intel, on transistor density
in integrated circuits, came to predict a doubling of microprocessor
performance every eighteen months. This cadence of introducing new
microprocessor generations became so predictable that over the past
few decades it drove everything from hardware systems architecture,
to computer software and programming language design, to consumer
product roadmaps. In 2006, David Patterson, the Turing Award-winning
computer scientist who was in part responsible for RAID disk arrays, RISC
processors, and the classic books on computer architecture by Patterson
and Hennessy, observed that the growth in microprocessor performance had
stalled, and instead semiconductor manufacturers had turned to increasing
the number of processing cores per chip. Today, Patterson says: “We
are a factor of 15 behind where we should be if Moore’s Law were
still operative.  We are in the post-Moore’s Law era.” We can no
longer throw faster computers at our product development requirements.

The stalling of single core performance, and the surprising (to me
anyway) growth of multi-core processors, leads me to believe we
need programming languages that are compiled to squeeze more
performance from single cores, that natively support efficient
multi-threading to leverage large numbers of parallel cores, and
that developers can productively write and debug. Google's Go, with
its compiled performance, its super lightweight "goroutines" threads
based on the Communicating Sequence Process (CSP) model, and its
simpler syntax and semantics than C++, seemed like good choice to evaluate.

My work here in Go has been based my prior work on traffic scheduling
more than two decades ago. The Generic Cell Rate Algorithm, or GCRA, which
I originally encountered in the ATM Forum document "Traffic Management
4.0", has become my go-to (so to speak) example with which to evalute
the real-time capabilities of a new programming language. Typically
implemented using either a "virtual scheduler" or a "leaky bucket"
approach, the GCRA is a tool that can be used to control the rate at
which events (an abstract term which the developer can interpret as
bytes, packets, log messages, what have you) are emitted (written, sent,
logged, etc.).

From oldest to newest, I have developed open-source implementations of
the GCRA in: C++ for Desperado (forked into Desperadito, which was later
forked into Grandote) in 2005; Java for Buckaroo in 2006; C for Diminuto
in 2008; and finally Go for Vamoose. They are not strictly ports from
one another, because my own understanding of the underlying algorithms,
architectures, and patterns has evolved over the years.

All of this was in turn is based on work I did on commercial products,
specifically, an ATM switch (A500), and an ATM interface card (TN2305),
during my time at Bell Labs in the latter half of the 1990s. On the
ATM switch, which applied the GCRA to hundreds of virtual circuits
ingressing on many OC-3 optical fiber ports, the GCRA was implemented
in hardware and used for traffic policing; my code merely computed its
parameters. On the ATM interface card, which had a few dozen virtual
circuits egressing on a single OC-3 port, the GCRA was used for traffic
shaping, and I implemented it all in firmware, writing in C++.

You would think that after having implemented the same basic algorithm,
described in a public standard, many times, I'd pretty much have it
down. But every time I revisit it, I learn something new. And by using
a different language, I encounter new challenges and have new insights.
This kind of deliberate practice has served me well throughout my career.

I still have a lot of affection for C and C++ (and Java and Python, in
their place); virtually all of my paying work these days continues to be
in C. My productivity in that language has been greatly enhanced by my
use of my Diminuto C systems programming library, all or parts of which
ships in a handful of commercial products from several different clients.
It remains to be seen if Go will yield the same kind of success for me.

## Repositories

<https://github.com/coverclock/com-diag-vamoose>

<https://github.com/coverclock/com-diag-diminuto>

<https://github.com/coverclock/com-diag-buckaroo>

<https://github.com/coverclock/com-diag-grandote>

## Articles

C. Overclock, "Traffic Management", 2006-12-25,
<http://coverclock.blogspot.com/2006/12/traffic-management.html>

C. Overclock, "Rate Control and Throttles", 2007-01-12,
<http://coverclock.blogspot.com/2007/01/rate-control-and-throttles.html>

C. Overclock, "Traffic Contracts", 2007-01-17,
<http://coverclock.blogspot.com/2007/01/traffic-contracts.html>

## Presentations

J. Sloan, "Going, Going, Gone: Learning A Systems Programming
Language for the Post-Moore's Law World", Gogo Business Aviation,
Broomfield Colorado, 2018-10-05,
<https://www.dropbox.com/s/mudrxf8vwf6og2r/Vamoose.pdf?dl=0>

## References

"The Go Programming Language - Documentation",
<https://golang.org/doc/>

W. Kennedy, "Scheduling In Go - Part I", 2018-08-12,
<https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html>

W. Kennedy, "Scheduling In Go - Part II", 2018-09-27,
<https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html>

J. Sloan, "ATM Traffic Management", Digital Aggregates Corporation, 2005-08-29,
<http://www.diag.com/reports/ATMTrafficManagement.html>

N. Giroux et al., "Traffic Management Specification Version 4.1", ATM Forum,
af-tm-0121.000, 1999-03

## Miscellaneous

<https://gopherize.me>

## Targets

Verious versions of this software has at one time or another been installed
and tested with the following combinations of hardware and software. Your
mileage may vary.

"Nickel"
Intel NUC5i7RYH    
Intel x86_64 64-bit    
Intel Core i7-5557U @ 3.10GHz x 8    
Ubuntu 18.04 "bionic"    
Linux 4.15.0    
go version go1.11 linux/amd64    

"Gold"    
Raspberry Pi 3B+    
ARM ARMv7 64-bit    
Broadcom BCM2837B0 Cortex-A53 @ 1.4GHz x 4      
Raspbian 9.4 "stretch"    
Linux 4.14.34    
go version go1.11 linux/arm    

## Installation

This is a mash-up of the directory structure expected by the standard Go
toolchain and the directory structure I use for my Digital Aggregates
projects. Your mileage may vary.

    export GOPATH="${HOME}/go"
    mkdir -p ${HOME}/src
    cd ${HOME}/src
    git clone https://github.com/coverclock/com-diag-vamoose
    mkdir -p ${GOPATH}/bin ${GOPATH}/pkg ${GOPATH}/src/github.com/coverclock
    cd ${GOPATH}/src/github.com/coverclock
    ln -s ${HOME}/src/com-diag-vamoose

## Unit Tests

    export GOPATH="${HOME}/go"
    cd ${GOPATH}/src
    go test -test.v github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks
    go test -test.v github.com/coverclock/com-diag-vamoose/Vamoose/pkg/fletcher
    go test -test.v github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle
    go test -test.v github.com/coverclock/com-diag-vamoose/Vamoose/pkg/gcra
    go test -test.v github.com/coverclock/com-diag-vamoose/Vamoose/pkg/contract

## Functional Tests

    export GOPATH="${HOME}/go"
    cd ${GOPATH}/src
    go build github.com/coverclock/com-diag-vamoose/Vamoose/cmd/fletch
    go build github.com/coverclock/com-diag-vamoose/Vamoose/cmd/shape
    dd if=/dev/urandom count=1000 | ./fletch -V -b 512 | ./shape -V -p 2048 -s 1024 -b 512 | ./fletch -V -b 512 > /dev/null

## Notes

### Policing

In the contract unit test, the jitter introduced by both the UDP
connection between the producer/shaper side and the policer/consumer
side of the Contract test can be seen in the traffic measurements. The
shaper measures something very close to the contract, about 1024Bps peak
and 512Bps sustained. The policer on the other hand measures a 32kBps
peak yet a 512Bps sustained. I haven't ruled out some boneheaded bug on
my part. But this attempt on my part to adapt the per-cell ATM GCRA to
event streams containing variable length packets makes me think it might
not be suitable for policing. (It's telling that the measured peak rate
by the policer always seems to be around thirty-two times the actual peak
rate provided by the shaper, even as their sustained rates are virtually
the same. I'm guessing this has something to do with either UDP datagram
queueing in the kernel and/or some artifact of the Go scheduler.)

    producer: end total=61440B mean=32.66347687400319B/burst maximum=64B/burst.
    shaper: end total=61440B mean=32.66347687400319B/burst maximum=64B/burst delay=0.06352381948059542s/burst peak=1024.7240767399094B/s sustained=511.9994023345643B/s.
    policer: end admitted=61440B policed=0B total=61440B mean=32.66347687400319B/burst maximum=64B/burst peak=31424.227065382667B/s sustained=511.999044093518B/s.
    consumer: end total=61440B.
    Actual: produced=61440:0x1a6d
    Actual: consumed=61440:0x1a6d

### gccgo

There are two Go compilers: the official Google compiler accessed via the
"go" command, and the Go front-end to the GNU compiler suite used via
the "gccgo" command (but it also has a "go" front end you can use). The
gccgo compiler has the potential to generate better code since the
GCC backend has generally good optimization. But the GCC Go run-time
library currently lags behind the offical compiler by several releases.

I did all my development using the official Google compiler, and that's
what I recommend you do too. But I have compiled and run functional tests
using gccgo via the Makefile that is part of this repo. I consider this
experimental.

