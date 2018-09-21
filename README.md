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
and possible useful packages.

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
I used Java in two product development efforts, one of which was
actually an embedded project for which we used a Java compiler, but that
was hardly mainstream either. I've done quite a bit of development in
Python, but that was strictly in the realm of building tools to support
my embedded work. I've been known to hack JavaScript in an emergency.

Why Go? Moore’s Law, based on an observation made in 1965 by Gordon
Moore, founder of Fairchild Semiconductor and Intel, on transistor
density in integrated circuits, came to predict a doubling of
microprocessor performance every eighteen months. This cadence of
introducing new microprocessor generations became so predictable
that over the past few decades it drove everything from hardware
systems architecture, to computer software and programming language
design, to consumer product roadmaps. In 2006, David Patterson, the
Turing Award-winning computer scientist who was in part responsible
for RAID disk arrays, RISC processors, and the classic books on
computer architecture by Patterson and Hennessy, observed that the
growth in microprocessor performance had stalled, and instead
semiconductor manufacturers had turned to increasing the number of
processing cores per chip. Today, Patterson says: “We are a factor
of 15 behind where we should be if Moore’s Law were still operative.
We are in the post-Moore’s Law era.” We can no longer throw faster
computers at our product development requirements.

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
more than two decades ago. The Generic Cell Rate Algorithm or GCRA, which
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
architectures and patterns has evolved over the years.

All of this was in turn is based on work I did on commercial products,
specifically, an ATM switch (A500), and an ATM interface card (TN2305),
during my time at Bell Labs in the latter half of the 1990s. On the
ATM switch, which applied the GCRA to hundreds of virtual circuits ingressing
on many OC-3 optical fiber ports, the GCRA was implemented in hardware and used
for traffic policing; my code merely computed its parameters. In the
ATM interface card, which had a few dozen virtual circuits egressing on a single
OC-3 port, the GCRA was used for traffic shaping, and I implemented it all in
firmware, writing in C++.

You would think that after having implemented the same basic algorithm,
described in a public standard, many times, I'd pretty much have it
down. But every time I revisit it, I learn something new. And by using
a different language, I encounter new challenges and have new insights.

I still have a lot of affection for C and C++ (and Java and Python,
in their place); virtually all of my paying work these days continues to
be in C. My productivity in that language has been greatly enhanced
by my use of my Diminuto C systems programming library, all or parts of which
ships in a handful of commercial products from several different clients.

## Repositories

<https://github.com/coverclock/com-diag-vamoose>

<https://github.com/coverclock/com-diag-diminuto>

<https://github.com/coverclock/com-diag-buckaroo>

<https://github.com/coverclock/com-diag-grandote>

## Articles

C. Overclock, "Traffic Management", 2006-12,
<http://coverclock.blogspot.com/2006/12/traffic-management.html>

C. Overclock, "Rate Control and Throttles", 2007-01,
<http://coverclock.blogspot.com/2007/01/rate-control-and-throttles.html>

C. Overclock, "Traffic Contracts", 2007-01,
<http://coverclock.blogspot.com/2007/01/traffic-contracts.html>

## References

<https://golang.org/doc/>

J. Sloan, "ATM Traffic Management", Digital Aggregates Corporation, 2005-08,
<http://www.diag.com/reports/ATMTrafficManagement.html>

N. Giroux et al., Traffic Management Specification Version 4.1, ATM Forum,
af-tm-0121.000, 1999-03

## Notes

I'm still trying to figure out the correct way to calculate the peak rate in
bytes per second, typically using the instantaneous rate calculated from the
interdeparture (in the shaper) and interarrival (in the policer) times. I
remain unhappy.

This is selected output from Throttle unit tests. Note that the sustained rates
closely match the respective 512B/s and 1024B/s Throttle settings.

    shaper: end total=30720B mean=32.20125786163522B/burst maximum=64B/burst delay=0.06272279126310272s/burst peak=3.8832519103094074e+06B/s sustained=510.7809393419717B/s.
    policer: end admitted=30720B policed=0B total=30720B mean=32.20125786163522B/burst maximum=64B/burst peak=30283.891949027326B/s sustained=510.78030067125894B/s.

    shaper: end total=61440B mean=32.82051282051282B/burst maximum=64B/burst delay=0.03194344442735043s/burst peak=4.32870208755149e+06B/s sustained=1019.3363099703025B/s.
    policer: end admitted=61440B policed=0B total=61440B mean=32.82051282051282B/burst maximum=64B/burst peak=57208.91333031246B/s sustained=1019.335670172402B/s.

This is selected output from Contract unit tests. Not only are the peak rates
(nominally 1024B/s) incorrect, but they suggest that the interdeparture and
interarrival times are are jittered all to heck. The sustained rates (nominally
512B/s) are quite close.

    shaper: end total=61440B mean=32.66347687400319B/burst maximum=64B/burst delay=0.06353551014141415s/burst peak=4.218533886583679e+06B/s sustained=511.9992536160214B/s.
    policer: end admitted=61440B policed=0B total=61440B mean=32.66347687400319B/burst maximum=64B/burst peak=32299.388239167594B/s sustained=511.9988639939872B/s.

Since the Throttle unit tests use the same parameters individually, and the
same core GCRA code, as the Contract unit tests use in composition, I'm
a little baffled as to what these numbers are really telling me. This would
suggested to me a bug in the Contract implementation, but I'm not seeing it
yet.