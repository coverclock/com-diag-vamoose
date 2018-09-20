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

My systems programming language of choice has changed over the decades,
depending on what I was getting paid to do and where I was doing it. In
the 1970s, it was IBM 360/370 Basic Assembler Language (BAL), and later
a structured language implemented in BAL macro language (which itself
was Turing complete), with an occasional foray into PL/1. In the 1980s,
it was PDP-11 Assembler Language (PAL). In the late 1980s and to the
mid-1990s it was C. In the mid-1990s to the mid-2000s it was C++, which
was mostly an artifact of the long history Bell Labs and its spinoffs
- where I was variously employed at the time - had for using C++ for
firmware development. In the 2010s, I saw a significant reduction in the
use of C++ for systems programming, in part due to the evolution of C++
into a langauge that was hard to learn, difficult to debug, and hence
not terribly productive to use.

During most of this time I cast about for an alternative to C and C++
for the kinds of real-time or close-to-bare-metal work I typically do. I
briefly considered D, but it didn't seem to catch on with the mainstream.
I used Java in two product development efforts, one of which was
actually an embedded project for which we used a Java compiler, but that
was hardly mainstream either. I've done quite a bit of development in
Python, but that was strictly in the realm of building tools to support
my embedded work. I've been known to hack JavaScript in an emergency.

When Google began promoting Go, it seemed like a reasonable thing to try
out, although it took me a while to get to it. I still have a lot of
affection for C and C++ (and Java and Python, in their place); virtually
all of my work these days continues to be in C. My productivity in C has
been greatly enhanced by my using my Diminuto C systems programming library.

My work here in Go has been based my prior work in earlier languages,
libraries, and repositories, from oldest to newest: C++ in Grandote
(forked from Desperadito, which was forked from Desperado) from 2005, Java
in Buckaroo from 2006, C in Diminuto from 2008, and finally Go in Vamoose.
They are not strictly ports from one another, because my own understanding of
the underlying algorithms, architectures and patterns has evolved over
the years. All of this was in turn based on work I did on commercial products
- specifically, an ATM switch, and an ATM interface card, both using OC-3
optical fiber - during my time at Bell Labs in the mid-to-late 1990s.

## Repositories

<https://github.com/coverclock/com-diag-vamoose>

<https://github.com/coverclock/com-diag-diminuto>

<https://github.com/coverclock/com-diag-buckaroo>

<https://github.com/coverclock/com-diag-grandote>

## Articles

<http://coverclock.blogspot.com/2006/12/traffic-management.html>

<http://coverclock.blogspot.com/2007/01/rate-control-and-throttles.html>

<http://coverclock.blogspot.com/2007/01/traffic-contracts.html>

## References

<https://golang.org/doc/>

J. Sloan, "ATM Traffic Management", 2005-08,
<http://www.diag.com/reports/ATMTrafficManagement.html>

N. Giroux et al., Traffic Management Specification Version 4.1, ATM Forum,
af-tm-0121.000, 1999-03
