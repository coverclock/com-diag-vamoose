package contract

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose

import (
    "testing"
	"github.com/coverclock/com-diag-vamoose/ticks"
	"github.com/coverclock/com-diag-vamoose/gcra"
 	"github.com/coverclock/com-diag-vamoose/fletcher"
    "math/rand"
    "time"
    "net"
    "fmt"
    "sync"
)

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func TestContractSimulated(t * testing.T) {
    const PEAK ticks.Ticks = 1024 // Bytes per second.
    const TOLERANCE ticks.Ticks = 64
    const SUSTAINED ticks.Ticks = 512 // Bytes per second.
	const BURST gcra.Events = 32768
    const OPERATIONS uint = 1000000
	const MARGIN ticks.Ticks = 200 // 0.5%
	var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
	var size gcra.Events = 0
    var total uint64 = 0
    var admissable bool = false
    var iops uint = 0
    var largest gcra.Events = 0
    var that gcra.Gcra
    
    frequency := ticks.Frequency()
    peak := (frequency + PEAK - 1) / PEAK
    sustained := (frequency + SUSTAINED - 1) / SUSTAINED
    now := ticks.Now()
    
	that = New(peak, TOLERANCE, sustained, BURST, now)
	t.Log(that.String())
	
	for iops = 0; iops < OPERATIONS; iops += 1 {

	    delay = that.Request(now)
	    if (delay >= 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }
	    now += delay
	    if (now >= 0) {} else { t.Errorf("OVERFLOW! %v", now) }
	    duration += delay
	    if (duration >= 0) {} else { t.Errorf("OVERFLOW! %v", duration) }
	    
	    delay = that.Request(now)
	    if (delay == 0) {} else { t.Errorf("FAILED! %v", delay); t.Log(that.String()) }

        size = gcra.Events(rand.Int63n(int64(BURST))) + 1
	    if (0 < size) {} else { t.Errorf("FAILED! %v", size) }
	    if (size <= gcra.Events(BURST)) {} else { t.Errorf("FAILED! %v", size) }
	    if (size > largest) { largest = size }
	    total += uint64(size)
	    if (total > 0) {} else { t.Errorf("OVERFLOW! %v", total) }

	    admissable = that.Commits(size)
	    if (admissable) {} else { t.Errorf("FAILED! %v", admissable); t.Log(that.String()) }

	}
	
	blocksize := float64(total) / float64(OPERATIONS)
	seconds := float64(duration) / float64(frequency)
	interarrival := seconds / float64(OPERATIONS)
	t.Logf("total=%vB largest=%vB/io mean=%vB/io mean=%vs/io\n", total, largest, blocksize, interarrival)
	if (duration > frequency) {} else { t.Errorf("FAILED! %v", duration) }

	bandwidth := float64(total) / float64(seconds)
	delta := bandwidth - float64(SUSTAINED)
	if (delta < 0) { delta = -delta }
    margin := float64(SUSTAINED) / float64(MARGIN)
	t.Logf("sustained=%vB/s delta=%vB/s margin=%vB/s\n", bandwidth, delta, margin)
	if (delta < margin) {} else { t.Errorf("FAILED! %v", delta) }
    
}

/*******************************************************************************
 * ACTUAL EVENT STREAM
 ******************************************************************************/

var mutex sync.Mutex

var producer_total uint64 = 1
var producer_checksum uint16 = 2

var consumer_total uint64 = 3
var consumer_checksum uint16 = 4

func producer(t * testing.T, limit uint64, delay time.Duration, output chan <- byte, done chan<- bool) {
    var total uint64 = 0
    var size int = 0
    var count int = 0
    var largest int = 0
    var datum [1] byte
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0
    
    mutex.Lock()
    fmt.Println("producer: begin.")
    mutex.Unlock()
    
    burst := cap(output)
    
    for limit > 0 {
        
        size = rand.Intn(burst) + 1
        if uint64(size) > limit {
            size = int(limit)
        }
        
        for index := size; index > 0; index -= 1 {
            datum[0] = byte(rand.Int31n(int32('~') - int32(' ') + 1) + int32(' '))
            c = fletcher.Checksum16(datum[:], &a, &b)
            output <- datum[0]
            total += 1  
        }
        
        count += 1
        
        limit -= uint64(size)
        
        if (size > largest) { largest = size }
         
        mutex.Lock()
        fmt.Printf("producer: produced=%vB total=%vB remaining=%vB.\n", size, total, limit)
        mutex.Unlock()
       
        time.Sleep(delay)
        
    }
    
    close(output)
       
    mean := float64(total) / float64(count)
    
    mutex.Lock()
    fmt.Printf("producer: end total=%vB/io mean=%vB/burst maximum=%vB/burst.\n", total, mean, largest);
    mutex.Unlock()
    
    producer_total = total
    producer_checksum = c
    
    done <- true
}

func shaper(t * testing.T, input <- chan byte, that gcra.Gcra, output net.PacketConn, address net.Addr, done chan<- bool) {
    var total uint64 = 0
    var okay bool = true
    var size int = 0
    var then ticks.Ticks = 0
    var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
    var alarmed bool = false
    var count int = 0
    var largest int = 0
    var before ticks.Ticks = 0
    var after ticks.Ticks = 0
    var interarrival ticks.Ticks = 0
    var bandwidth float64 = 0.0
    var fastest float64 = 0.0
        
    mutex.Lock()
    fmt.Println("shaper: begin.");
    mutex.Unlock()

    burst := cap(input)
    buffer := make([] byte, burst)
    
    frequency := float64(ticks.Frequency())
    
    then = ticks.Now()
    
    for {

        buffer[0], okay = <- input
        if !okay {
            break
        }
        total += 1

        for size = 1; (size < len(buffer)) && (len(input) > 0); size +=1 {
            buffer[size], okay = <- input
            if !okay {
                // Should never happen.
                break
            }
            total += 1
        }
        
        if (size > largest) { largest = size }
        
        now = ticks.Now()
        delay = that.Request(now)
        duration = 0
        for delay > 0 {
            duration += delay                       
            time.Sleep(time.Duration(delay))
            now = ticks.Now()
            delay = that.Request(now) 
        }
        
        after = ticks.Now()
        written, failure := output.WriteTo(buffer[:size], address)
        if failure != nil {
            t.Fatalf("shaper: failure=%v!\n", failure);
        }
        if (written != size) {
            t.Fatalf("shaper: written=%v size=%v!\n", written, size);
        }
        
        if (count > 0) {
            interarrival = after - before
            bandwidth = float64(size) / float64(interarrival)
            if (bandwidth > fastest) { fastest = bandwidth }
        }
        before = after
        count += 1
        
        fmt.Printf("shaper: delay=%vs written=%vB total=%vB.\n", float64(duration) / frequency, written, total);

        alarmed = !that.Commits(gcra.Events(size))
        if alarmed {
            t.Fatalf("shaper: alarmed=%v!\n", alarmed);
        }

    }
    
    now = ticks.Now()
    delay = that.Request(now)
    duration = 0
    for delay > 0 {
        duration += delay                       
        time.Sleep(time.Duration(delay))
        now = ticks.Now()
        delay = that.Request(now) 
    }
    
    buffer[0] = 0
    size = 1
    written, failure := output.WriteTo(buffer[0:size], address)
    if failure != nil {
        t.Fatalf("shaper: failure=%v!\n", failure);
    } 
    if (written != 1) {
        t.Fatalf("shaper: written=%v size=%v!\n", written, 1);
    }
        
    fmt.Printf("shaper: delay=%vs written=%vB total=%vB.\n", float64(duration) / frequency, written, total);

    alarmed = !that.Commits(gcra.Events(size))
    if alarmed {
        t.Fatalf("shaper: alarmed=%v!\n", alarmed);
    }

    now = ticks.Now()
    sustained := float64(total) * frequency / float64(now - then)
    mean := float64(total) / float64(count)
    peak := frequency * fastest

    mutex.Lock()
    fmt.Printf("shaper: end total=%vB mean=%vB/burst maximum=%vB/burst duration=%vs sustained=%vB/s peak=%vB/s.\n", total, mean, largest, duration, sustained, peak);
    mutex.Unlock()
    
    done <- true
}

func policer(t * testing.T, input net.PacketConn, that gcra.Gcra, output chan<- byte, done chan<- bool) {
    var total uint64 = 0
    var now ticks.Ticks = 0
    var admissable bool = false
    var eof bool = false
    
    mutex.Lock()
    fmt.Println("policer: begin.");
    mutex.Unlock()
    
    burst := cap(output)
    buffer := make([] byte, burst)
    
    for !eof {
    
        read, _, failure := input.ReadFrom(buffer)
        if failure != nil {
            t.Fatalf("policer: failure=%v!\n", failure);
        }
        
        if (read <= 0) {
            t.Fatalf("policer: read=%v!\n", read);
        }
        
        now = ticks.Now()
        admissable = that.Admits(now, gcra.Events(read))
        if admissable {
            mutex.Lock()
            fmt.Printf("policer: admitted=%vB total=%vB.\n", read, total)
            mutex.Unlock()
            for index := 0; index < read; index += 1 {
                if buffer[index] == 0 {
                    eof = true
                    break
                }
                output <- buffer[index]
                total += 1
            }
        } else {
            mutex.Lock()
            fmt.Printf("policer: policed=%vB total=%vB?\n", read, total);         
            mutex.Unlock()
            if buffer[read - 1] == 0 {
                eof = true
            }
        }
    
    }
    
    close(output)
    
    mutex.Lock()
    fmt.Println("policer: end");
    mutex.Unlock()
    
    done <- true
}

func consumer(t * testing.T, input <-chan byte, done chan<- bool) {
    var total uint64 = 0
    var okay bool = true
    var size int = 0
    var a uint8 = 0
    var b uint8 = 0
    var c uint16 = 0
   
    mutex.Lock()
    fmt.Println("consumer: begin.");
    mutex.Unlock()

    burst := cap(input)
    buffer := make([] byte, burst)
    
    for {

        buffer[0], okay = <- input
        if !okay {
            break
        }
        total += 1

        for size = 1; (size < len(buffer)) && (len(input) > 0); size +=1 {
            buffer[size], okay = <- input
            if !okay {
                // Should never happen.
                break
            }
            total += 1
        }
                
        c = fletcher.Checksum16(buffer[:size], &a, &b)

    }
    
    mutex.Lock()
    fmt.Printf("consumer: end total=%vB.\n", total);
    mutex.Unlock()
    
    consumer_total = total
    consumer_checksum = c
    
    done <- true
}

func TestContractActual(t * testing.T) {
    const PEAK int = 1024				// bytes per second
    const SUSTAINED int = 512			// bytes per second
    const BURST int = 64				// bytes
    const DURATION int = 60				// seconds
    const DELAY time.Duration = 1000	// nanoseconds
    var failure error
    
    mutex.Lock()
    fmt.Println("Beginning.")
    mutex.Unlock()
    
    done := make(chan bool, 4)
    defer close(done)
    
    supply := make(chan byte, BURST)
    // producer closes to signal EOF to shaper.
    
    demand := make(chan byte, BURST)
    // policer closes to signal EOF to consumer.
        
    source, failure := net.ListenPacket("udp", ":5555")
    if failure != nil {
        t.Fatal(failure)
    }
    defer source.Close()
    
    mutex.Lock()
    fmt.Printf("source=%+v.\n", source);
    mutex.Unlock()    
           
    sink, failure := net.ListenPacket("udp", ":0")
    if failure != nil {
        t.Fatal(failure)
    }
    defer sink.Close()
    
    mutex.Lock()
    fmt.Printf("sink=%+v.\n", sink);
    mutex.Unlock()    
 
    destination, failure := net.ResolveUDPAddr("udp", "localhost:5555")
    if failure != nil {
        t.Fatal(failure)
    }
     
    mutex.Lock()
    fmt.Printf("destination=%+v.\n", destination);
    mutex.Unlock()    
   
    frequency := ticks.Frequency()
    peak := frequency / ticks.Ticks(PEAK)
    jitter := peak / 200
    sustained := frequency / ticks.Ticks(SUSTAINED)
    burst := gcra.Events(BURST)
    now := ticks.Now()
    
    shape := New(peak, 0, sustained, burst, now)
    
    mutex.Lock()
    fmt.Printf("shape=%+v.\n", shape);
    mutex.Unlock()    
    
    police := New(peak, jitter, sustained, burst, now)
     
    mutex.Lock()
    fmt.Printf("police=%+v.\n", police);
    mutex.Unlock()    
   
    mutex.Lock()
    fmt.Println("Starting.")
    mutex.Unlock()
   
    go consumer(t, demand, done)
    go policer(t, source, police, demand, done)
    go shaper(t, supply, shape, sink, destination, done)
    go producer(t, uint64(DURATION) * uint64(SUSTAINED), DELAY, supply, done)
    
    mutex.Lock()
    fmt.Println("Waiting.")
    mutex.Unlock()
    
    <- done
    <- done
    <- done
    <- done
       
    mutex.Lock()
    fmt.Println("Checking.")
    mutex.Unlock()

    if (consumer_total == producer_total) {} else { t.Fatalf("consumer_total=%v producer_total=%v\n", consumer_total, producer_total) }
    if (consumer_checksum == producer_checksum) {} else { t.Fatalf("consumer_checksum=%v producer_checksum=%v\n", consumer_checksum, producer_checksum) }
   
    mutex.Lock()
    fmt.Println("Ending.")
    mutex.Unlock()
}

