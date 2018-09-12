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

Implements a Generic Cell Rate Algorithm (GCRA) using a Virtual Scheduler (VS).
This can in turn be used to implement a variety of traffic shaping and rate
control algorithms. The VS works by monitoring the inter-arrival interval of
events and comparing that interval to the expected value. When the cumulative
error in the inter-arrival interval exceeds a threshold, the throttle becomes
"alarmed" and the traffic stream is in violation of its contract. In the
original TM spec, an event was the emission (if traffic shaping) or arrival
(if traffic policing) of an ATM cell, but it could be data blocks, error
reports, or any other kind of real-time activity. In this implementation,
it can even be variable length data blocks, in which the traffic contract
describes the mean bandwidth of the traffic stream, not the instantaneous
bandwidth as with ATM. In the original TM spec, the variable "i" was the
increment or contracted inter-arrival interval, "l" was the limit or
threshold, "x" was the expected inter-arrival interval for the next event,
and "x1" was the aggregate early duration accumulated so far. A throttle can
be used to smooth out low frequency events over a long duration, or to
implement a leaky bucket algorithm.

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
