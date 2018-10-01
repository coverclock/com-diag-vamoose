/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Provides a test harnesses for testing GCRA implementations.
//
package harness

import (
    "testing"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/ticks"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/throttle"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/fletcher"
    "math/rand"
    "net"
    "fmt"
    "sync"
)

var Debug bool = false

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

// SimulatedEventStream provides a virtual-time test of a GCRA given the
// shaping GCRA, the policing GCRA, the maximum size of an event burst, and
// the total number of iterations to perform.
func SimulatedEventStream(t * testing.T, shape throttle.Throttle, police throttle.Throttle, burst int, iterations int) {
    var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var size throttle.Events = 0
    var maximum throttle.Events = 0
    var total uint64 = 0
    var admissable bool = false
    var admitted bool = false
    var rate float64 = 0
    var peak float64 = 0

    fmt.Printf("Simulated: shape=%v.\n", shape);
    fmt.Printf("Simulated: police=%v.\n", police);

    frequency := float64(ticks.Frequency())

    for ii := 0; ii < iterations; ii += 1 {

        delay = shape.Request(now)
        now += delay
        if now >= 0 {} else { t.Fatalf("Simulated: OVERFLOW! %v\n", now) }
        duration += delay
        if duration >= 0 {} else { t.Log(shape.String()); t.Fatalf("Simulated: OVERFLOW! %v\n", duration) }

        if ii <= 0 {
            // Do nothing.
        } else if delay <= 0 {
            // Do nothing.
        } else {
            rate = float64(size) * frequency / float64(delay)
            if rate > peak {
                peak = rate
            }
        }

        delay = shape.Request(now)
        if delay == 0 {} else { t.Log(shape.String()); t.Fatalf("Simulated: FAILED! %v\n", delay);  }

        size = throttle.Events(rand.Int63n(int64(burst))) + 1
        if 0 < size {} else { t.Log(shape.String()); t.Fatalf("Simulated: FAILED! %v\n", size) }
        if size <= throttle.Events(burst) {} else { t.Fatalf("Simulated: FAILED! %v\n", size) }
        if size > maximum { maximum = size }
        total += uint64(size)
        if total > 0 {} else { t.Fatalf("Simulated: OVERFLOW! %v\n", total) }

        admissable = shape.Commits(size)
        if admissable {} else { t.Log(shape.String); t.Fatalf("Simulated: FAILED! %v\n", admissable) }

        admitted = police.Admits(now, size)
        if admitted {} else { t.Log(police.String); t.Fatalf("Simulated: FAILED! %v\n", admitted) }

    }

    delay = shape.Comply()
    now += delay
    if now >= 0 {} else { t.Fatalf("Simulated: OVERFLOW! %v\n", now) }
    duration += delay
    if duration >= 0 {} else { t.Fatalf("Simulated: OVERFLOW! %v\n", duration) }

    admissable = shape.Update(now)
    if admissable {} else { t.Log(shape.String); t.Fatalf("Simulated: FAILED! %v\n", admissable) }

    admitted = police.Update(now)
    if admitted {} else { t.Log(police.String); t.Fatalf("Simulated: FAILED! %v\n", admitted) }

    fmt.Printf("Simulated: shape=%v.\n", shape);
    fmt.Printf("Simulated: police=%v.\n", police);

    average := float64(total) / float64(iterations)
    seconds := float64(duration) / frequency
    mean := seconds / float64(iterations)
    sustained := float64(total) * frequency / float64(duration)

    fmt.Printf("Simulated: total=%vB mean=%vB/io maximum=%vB/io latency=%vs/io peak=%vB/s sustained=%vB/s\n", total, average, maximum, mean, peak, sustained)

}

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

var mutex sync.Mutex

func producer(t * testing.T, limit uint64, output chan <- byte, totalp * uint64, checksump * uint16, done chan<- bool) {
    var total uint64 = 0
    var size int = 0
    var count int = 0
    var largest int = 0
    var datum [1] byte
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0

    // The output channel is one byte larger than the burst size to allow for
    // the end of record character. Multiple records can be written to the
    // channel if the Producer gets ahead of the Shaper, as the two operate
    // asynchronously. We want the Producer to generate data as quickly as the
    // Shaper can consume it. We don't want the Shaper to block other than to
    // manage the traffic stream to the Policer.

    burst := cap(output) - 1

    mutex.Lock()
    fmt.Printf("producer: begin burst=%vB.\n", burst)
    mutex.Unlock()

    for limit > 0 {

        // Choose a random size somewhere in the range of one to the maximum
        // burst size.

        size = rand.Intn(burst) + 1
        if uint64(size) > limit {
            size = int(limit)
        }

        // Producer only generates printable characters. There's no special
        // reason for this other maybe than convenience for debugging. All
        // it really needs to do is not generate the end of record character
        // 0x00 as part of the payload.

        for remain := size; remain > 0; remain -= 1 {
            datum[0] = byte(rand.Int31n(int32('~') - int32(' ') + 1) + int32(' '))
            c = fletcher.Checksum16(datum[:], &a, &b)
            output <- datum[0]
        }
        total += uint64(size)

        if (size > largest) { largest = size }

        count += 1

        // Write an 0x00 end of record indicator to the channel. That allows
        // the Shaper to delimit the data that goes into the datagram it sends
        // to the Policer.

        datum[0] = 0x00
        output <- datum[0]

        if Debug {
            mutex.Lock()
            fmt.Printf("producer: produced=%vB total=%vB remaining=%vB.\n", size, total, limit)
            mutex.Unlock()
        }

        limit -= uint64(size)

        ticks.Sleep(0)

    }

    close(output)

    *totalp = total
    *checksump = c

    mean := float64(total) / float64(count)

    mutex.Lock()
    fmt.Printf("producer: end total=%vB mean=%vB/burst maximum=%vB/burst.\n", total, mean, largest);
    mutex.Unlock()

    done <- true
}

func shaper(t * testing.T, input <- chan byte, that throttle.Throttle, output net.PacketConn, address net.Addr, done chan<- bool) {
    var total uint64 = 0
    var datum byte = 0
    var okay bool = true
    var size int = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var accumulated ticks.Ticks = 0
    var alarmed bool = false
    var count int = 0
    var largest int = 0
    var now ticks.Ticks = 0
    var rate float64 = 0.0
    var peak float64 = 0.0

    burst := cap(input) - 1

    mutex.Lock()
    fmt.Printf("shaper: begin burst=%vB.\n", burst);
    mutex.Unlock()

    buffer := make([] byte, burst)

    frequency := float64(ticks.Frequency())

    before := ticks.Now()

    for {

        // The Shaper assumes that the Producer has filled the channel with
        // data. So it delays now assuming there will be data in the channel
        // to consume. Otherwise, there will be an additional delay that will
        // affect the traffic management.

        now = ticks.Now()
        delay = that.Request(now)
        if delay < 0 {
            t.Fatalf("shaper: delay=%v!\n", delay)
        }

        duration = delay
        accumulated += delay

        ticks.Sleep(delay)

        now = ticks.Now()
        delay = that.Request(now)
        if delay != 0 {
            t.Fatalf("shaper: delay=%v!\n", delay)
        }

        // Calculate the peak data rate, which is based on the instantaneous
        // interarrival time between records from the Producer, and the size
        // of those records. The interarrival time is actually the traffic
        // management delay enforced by the traffic contract, since the Producer
        // constantly generates data for the Shaper to consume.

        if count == 0 {
            // Do nothing.
        } else if duration <= 0 {
            // Do nothing.
        } else {
            rate = float64(size) * frequency / float64(duration)
            if rate > peak { 
                peak = rate
            }
        }

        datum, okay = <- input
        if !okay {
            break
        }
        if datum == 0 {
            t.Fatalf("shaper: buffer[0]=%v\n", datum)
        }
        buffer[0] = datum
        size = 1

        for {
            datum, okay = <- input
            if !okay {
                t.Fatalf("shaper: !okay!\n")
            }
            if datum == 0 {
                break
            }
            if size >= len(buffer) {
                t.Fatalf("shaper: size=%v!\n", size)
            }
            buffer[size] = datum
            size += 1
        }

        total += uint64(size)  
        if (size > largest) { largest = size }

        // Write a datagram to the Policer. We use UDP so that if we get
        // to far ahead of the Policer (which would indicate a failure in
        // the traffic shaping), datagrams are discarded instead of the
        // Shaper blocking. This will cause the byte count and checksum
        // computed by the Consumer to fail.

        written, failure := output.WriteTo(buffer[:size], address)
        if failure != nil {
            t.Fatalf("shaper: failure=%v!\n", failure);
        }
        if written != size {
            t.Fatalf("shaper: written=%v size=%v!\n", written, size);
        }

        alarmed = !that.Commits(throttle.Events(size))
        if alarmed {
            t.Logf("shaper: contract=%v!\n", that)
            t.Fatalf("shaper: alarmed=%v!\n", alarmed);
        }

        if Debug {
            mutex.Lock()
            fmt.Printf("shaper: delay=%vs written=%vB total=%vB.\n", float64(duration) / frequency, written, total);
            mutex.Unlock()
        }

        count += 1

    }

    now = ticks.Now()
    that.Update(now)
    delay = that.Comply()

    mutex.Lock()
    fmt.Printf("shaper: delay=%vs.\n", float64(delay) / frequency);
    mutex.Unlock()

    ticks.Sleep(delay)
    now = ticks.Now()
    that.Update(now)

    after := now

    // Write the end of file indicator to the Policer by sending a datagram
    // with just a 0x00 in it.

    buffer[0] = 0x00
    size = 1
    written, failure := output.WriteTo(buffer[0:size], address)
    if failure != nil {
        t.Fatalf("shaper: failure=%v!\n", failure);
    } 
    if written != 1 {
        t.Fatalf("shaper: written=%v size=%v!\n", written, 1);
    }

    // Calculate the sustained rate.

    average := float64(accumulated) / float64(count) / frequency
    mean := float64(total) / float64(count)
    sustained := float64(total) * frequency / float64(after - before)

    mutex.Lock()
    fmt.Printf("shaper: end total=%vB mean=%vB/burst maximum=%vB/burst delay=%vs/burst peak=%vB/s sustained=%vB/s.\n", total, mean, largest, average, peak, sustained);
    mutex.Unlock()

    done <- true
}

func policer(t * testing.T, input net.PacketConn, that throttle.Throttle, output chan<- byte, done chan<- bool) {
    var eof bool = false
    var read int = 0
    var failure error
    var total uint64 = 0
    var admitted uint64 = 0
    var policed uint64 = 0
    var admissable bool = false
    var count int = 0
    var largest int = 0
    var now ticks.Ticks = 0
    var then ticks.Ticks = 0
    var rate float64 = 0.0
    var peak float64 = 0.0

    burst := cap(output)

    mutex.Lock()
    fmt.Printf("policer: begin burst=%vB.\n", burst);
    mutex.Unlock()

    buffer := make([] byte, burst)

    frequency := float64(ticks.Frequency())

    before := ticks.Now()

    for !eof {

        read, _, failure = input.ReadFrom(buffer)
        if failure != nil {
            t.Fatalf("policer: failure=%v!\n", failure);
        }
        if read <= 0 {
            t.Fatalf("policer: read=%v!\n", read);
        }
        if read > burst {
            t.Fatalf("policer: read=%v!\n", read);
        }
        if buffer[read - 1] == 0x00 {
            eof = true
            read -= 1
        }

        then = now
        now = ticks.Now()

        // Calculate the peak rate based on the datagram size and the
        // interarrival time.

        if count == 0 {
            // Do nothing.
        } else if read == 0 {
            // Do nothing.
        } else if now <= then {
            // Do nothing.
        } else {
            rate = float64(read) * frequency / float64(now - then)
            if rate > peak { peak = rate }
        }

        // If the datagram consisted of just the end of file indicator (which
        // we don't assume above, but that will be the case because of how the
        // Shaper sends it), than the read size of the datagram will be zero.

        if read > 0 {

            total += uint64(read)
            if (read > largest) { largest = read }

            admissable = that.Admits(now, throttle.Events(read))
            if admissable {
                admitted += uint64(read)
                if Debug {
                    mutex.Lock()
                    fmt.Printf("policer: read=%vB admitted=%vB total=%vB.\n", read, admitted, total)
                    mutex.Unlock()
                }
            } else {
                policed += uint64(read)
                if Debug {
                    mutex.Lock()
                    fmt.Printf("policer: read=%vB policed=%vB total=%vB?\n", read, policed, total)  
                    mutex.Unlock()
                }
            }

             for index := 0; index < read; index += 1 {
                output <- buffer[index]
            }

            count += 1

        } else if eof {

            that.Update(now)

        } else {

            // Should never happen. Should probably be fatal if it does.

        }

        ticks.Sleep(0)

    }

    after := ticks.Now()

    close(output)

    // Calculate the sustained rate.

    mean := float64(total) / float64(count)
    sustained := float64(total) * frequency / float64(after - before)

    // We kinda hope no data is ever policed. But it can happen because of
    // jitter introduced by the use of UDP and by the Goroutine scheduler.
    // It is possible for datagrams to pile up in the kernel, and so we
    // receive two consecutively with very little time in betweeen. This seems
    // to be non-deterministic. But I'm not ruling out some boneheaded
    // mistake on my part. But ATM (from whence the GCRA came) applies
    // policing on a cell by call basis, not on a packet by packet basis,
    // where all cells are the same size, but packets may be differently
    // sized.

    mutex.Lock()
    if policed > 0 { fmt.Printf("policer: POLICED!\n") }
    fmt.Printf("policer: end admitted=%vB policed=%vB total=%vB mean=%vB/burst maximum=%vB/burst peak=%vB/s sustained=%vB/s.\n", admitted, policed, total, mean, largest, peak, sustained)
    mutex.Unlock()

    done <- true
}

func consumer(t * testing.T, input <-chan byte, totalp * uint64, checksump * uint16, done chan<- bool) {
    var total uint64 = 0
    var buffer [1] byte
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0

    burst := cap(input)

    mutex.Lock()
    fmt.Printf("consumer: begin burst=%vB.\n", burst);
    mutex.Unlock()

    for buffer[0] = range input {

        total += 1          
        c = fletcher.Checksum16(buffer[:], &a, &b)

        if Debug && ((total % uint64(burst)) == 0) {
            mutex.Lock()
            fmt.Printf("consumer: total=%vB.\n", total)
            mutex.Unlock()            
        }

        ticks.Sleep(0)
    }

    *totalp = total
    *checksump = c

    mutex.Lock()
    fmt.Printf("consumer: end total=%vB.\n", total);
    mutex.Unlock()

    done <- true
}

// ActualEventStream provides a real-time test of a GCRA given a GCRA
// used for traffic shaping, a GCRA used for traffic policing, a supply
// channel used for supply-side bursts, a demand channel used for demand-
// side bursts, and the total number of events in the event stream.
func ActualEventStream(t * testing.T, shape throttle.Throttle, police throttle.Throttle, supply chan byte, demand chan byte, total uint64) {
    var failure error
    var producertotal uint64 = 1
    var producerchecksum uint16 = 2
    var consumertotal uint64 = 3
    var consumerchecksum uint16 = 4

    fmt.Println("Actual: Beginning.")

    fmt.Printf("Actual: shape=%v.\n", shape);
    fmt.Printf("Actual: police=%v.\n", police);

    done := make(chan bool, 4)
    defer close(done)

    source, failure := net.ListenPacket("udp", ":5555")
    if failure != nil {
        t.Fatalf("Actual: FAILED! %v\n", failure)
    }
    defer source.Close()
    fmt.Printf("Actual: source=%+v.\n", source);

    sink, failure := net.ListenPacket("udp", ":0")
    if failure != nil {
        t.Fatalf("Actual: FAILED! %v\n", failure)
    }
    defer sink.Close()
    fmt.Printf("Actual: sink=%+v.\n", sink);

    destination, failure := net.ResolveUDPAddr("udp", "localhost:5555")
    if failure != nil {
        t.Fatalf("Actual: FAILED! %v\n", failure)
    }
    fmt.Printf("Actual: destination=%+v.\n", destination);

    fmt.Println("Actual: Starting.")

    go consumer(t, demand, &consumertotal, &consumerchecksum, done)
    go policer(t, source, police, demand, done)
    go shaper(t, supply, shape, sink, destination, done)
    go producer(t, total, supply, &producertotal, &producerchecksum, done)

    mutex.Lock()
    fmt.Println("Actual: Waiting.")
    mutex.Unlock()

    <- done
    <- done
    <- done
    <- done

    fmt.Println("Actual: Checking.")

    fmt.Printf("Actual: produced=%v:%#04x\n", producertotal, producerchecksum);
    fmt.Printf("Actual: consumed=%v:%#04x\n", consumertotal, consumerchecksum);

    fmt.Printf("Actual: shape=%v.\n", shape);
    fmt.Printf("Actual: police=%v.\n", police);

    if (consumertotal == producertotal) {} else { t.Fatalf("Actual: FAILED! consumertotal=%v producertotal=%v\n", consumertotal, producertotal) }
    if (consumerchecksum == producerchecksum) {} else { t.Fatalf("Actual: FAILED! consumerchecksum=%v producerchecksum=%v\n", consumerchecksum, producerchecksum) }

    fmt.Println("Actual: Ending.")
}

