package harness

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
    "net"
    "fmt"
    "sync"
)

/*******************************************************************************
 * SIMULATED EVENT STREAM
 ******************************************************************************/

func SimulatedEventStream(t * testing.T, that gcra.Gcra, blocksize int, operations int) {
	var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var duration ticks.Ticks = 0
	var size gcra.Events = 0
	var maximum gcra.Events = 0
    var total uint64 = 0
    var admissable bool = false

	t.Log(that.String())
	
	for ii := 0; ii < operations; ii += 1 {

	    delay = that.Request(now)
	    now += delay
	    if now >= 0 {} else { t.Fatalf("OVERFLOW! %v\n", now) }
	    duration += delay
	    if duration >= 0 {} else { t.Log(that.String()); t.Fatalf("OVERFLOW! %v\n", duration) }

	    delay = that.Request(now)
	    if delay == 0 {} else { t.Log(that.String()); t.Fatalf("FAILED! %v\n", delay);  }

        size = gcra.Events(rand.Int63n(int64(blocksize))) + 1
	    if 0 < size {} else { t.Log(that.String()); t.Fatalf("FAILED! %v\n", size) }
	    if size <= gcra.Events(blocksize) {} else { t.Fatalf("FAILED! %v\n", size) }
	    if size > maximum { maximum = size }
	    total += uint64(size)
	    if total > 0 {} else { t.Fatalf("OVERFLOW! %v\n", total) }

	    admissable = that.Commits(size)
	    if admissable {} else { t.Log(that.String); t.Fatalf("FAILED! %v\n", admissable) }

	}
	
	delay = that.GetDeficit()
	now += delay
	if now >= 0 {} else { t.Fatalf("OVERFLOW! %v\n", now) }
	duration += delay
	if duration >= 0 {} else { t.Fatalf("OVERFLOW! %v\n", duration) }
	
	frequency := float64(ticks.Frequency())
	average := float64(total) / float64(operations)
	seconds := float64(duration) / frequency
	mean := seconds / float64(operations)
	actual := float64(total) * frequency / float64(duration)
	t.Logf("total=%vB mean=%vB/io maximum=%vB/io latency=%vs/io actual=%vB/s\n", total, average, maximum, mean, actual)
    
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
     
    burst := cap(output) - 1
   
    mutex.Lock()
    fmt.Printf("producer: begin burst=%vB.\n", burst)
    mutex.Unlock()
    
    for limit > 0 {
        
        size = rand.Intn(burst) + 1
        if uint64(size) > limit {
            size = int(limit)
        }
        
        for remain := size; remain > 0; remain -= 1 {
            datum[0] = byte(rand.Int31n(int32('~') - int32(' ') + 1) + int32(' '))
            c = fletcher.Checksum16(datum[:], &a, &b)
            output <- datum[0]
        }
        total += uint64(size)
        
        if (size > largest) { largest = size }
        
        count += 1

        datum[0] = 0
        output <- datum[0]
         
        mutex.Lock()
        fmt.Printf("producer: produced=%vB total=%vB remaining=%vB.\n", size, total, limit)
        mutex.Unlock()
        
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

func shaper(t * testing.T, input <- chan byte, that gcra.Gcra, output net.PacketConn, address net.Addr, done chan<- bool) {
    var total uint64 = 0
    var datum byte = 0
    var okay bool = true
    var size int = 0
    var then ticks.Ticks = 0
    var now ticks.Ticks = 0
    var delay ticks.Ticks = 0
    var accumulated ticks.Ticks = 0
    var alarmed bool = false
    var count int = 0
    var largest int = 0
    var before ticks.Ticks = 0
    var after ticks.Ticks = 0
    var rate float64 = 0.0
    var fastest float64 = 0.0
    var briefest ticks.Ticks = 0

    burst := cap(input) - 1
        
    mutex.Lock()
    fmt.Printf("shaper: begin burst=%vB.\n", burst);
    mutex.Unlock()
    
    buffer := make([] byte, burst)
    
    frequency := float64(ticks.Frequency())
    
    then = ticks.Now()
    
    for {

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
        
        now = ticks.Now()
        delay = that.Request(now)
        ticks.Sleep(delay)
        if count == 0 {
            briefest = delay
        } else if delay == 0 {
            // Do nothing
        } else if delay < briefest {
            briefest = delay
        } else {
            // Do nothing.
        }
        accumulated += delay
        
        now = ticks.Now()
        delay = that.Request(now)
        if delay > 0 {
            t.Fatalf("shaper: delay=%v!\n", delay)
        }
        
        written, failure := output.WriteTo(buffer[:size], address)
        if failure != nil {
            t.Fatalf("shaper: failure=%v!\n", failure);
        }
        if (written != size) {
            t.Fatalf("shaper: written=%v size=%v!\n", written, size);
        }

        alarmed = !that.Commits(gcra.Events(size))
        if alarmed {
            t.Logf("shaper: contract=%v!\n", that)
            t.Fatalf("shaper: alarmed=%v!\n", alarmed);
        }

        after = ticks.Now()
        if count > 0 {
            if after <= before {
                t.Fatalf("shaper: before=%v after=%v\n", before, after)
            }
            rate = float64(size) / float64(after - before)
            if (rate > fastest) { fastest = rate }
        }
        before = after
        
        fmt.Printf("shaper: delay=%vs written=%vB total=%vB rate=%vB/s.\n", float64(delay) / frequency, written, total, float64(rate) * frequency);

        count += 1

    }
    
    now = ticks.Now()
    that.Update(now)
    delay = that.GetDeficit()

    mutex.Lock()
    fmt.Printf("shaper: delay=%vs.\n", float64(delay) / frequency);
    mutex.Unlock()

    ticks.Sleep(delay)
    now = ticks.Now()
    that.Update(now)
    
    buffer[0] = 0
    size = 1
    written, failure := output.WriteTo(buffer[0:size], address)
    if failure != nil {
        t.Fatalf("shaper: failure=%v!\n", failure);
    } 
    if (written != 1) {
        t.Fatalf("shaper: written=%v size=%v!\n", written, 1);
    }

    average := (float64(accumulated) / float64(count)) / frequency
    minimum := float64(briefest) / frequency
    sustained := float64(total) * frequency / float64(now - then)
    mean := float64(total) / float64(count)
    peak := frequency * fastest

    mutex.Lock()
    fmt.Printf("shaper: end total=%vB mean=%vB/burst maximum=%vB/burst delay=%vs/burst minimum=%vs/burst sustained=%vB/s peak=%vB/s.\n", total, mean, largest, average, minimum, sustained, peak);
    mutex.Unlock()
    
    done <- true
}

func policer(t * testing.T, input net.PacketConn, that gcra.Gcra, output chan<- byte, done chan<- bool) {
    var total uint64 = 0
    var admitted uint64 = 0
    var policed uint64 = 0
    var now ticks.Ticks = 0
    var admissable bool = false
    var eof bool = false
    var count int = 0
    var largest int = 0

    burst := cap(output)
    
    mutex.Lock()
    fmt.Printf("policer: begin burst=%vB.\n", burst);
    mutex.Unlock()
    
    buffer := make([] byte, burst)
    
    for !eof {
    
        read, _, failure := input.ReadFrom(buffer)
        if failure != nil {
            t.Fatalf("policer: failure=%v!\n", failure);
        }
        if read <= 0 {
            t.Fatalf("policer: read=%v!\n", read);
        }
        if buffer[read - 1] == 0 {
            eof = true
            read -= 1
        }
        
        now = ticks.Now()
        
        if read > 0 {
            count += 1
            total += uint64(read)
            if (read > largest) { largest = read }
            admissable = that.Admits(now, gcra.Events(read))
            if admissable {
                admitted += uint64(read)
                for index := 0; index < read; index += 1 {
                    output <- buffer[index]
                }
                mutex.Lock()
                fmt.Printf("policer: read=%vB admitted=%vB total=%vB.\n", read, admitted, total)
                mutex.Unlock()
            } else {
                policed += uint64(read)
                mutex.Lock()
                fmt.Printf("policer: read=%vB policed=%vB total=%vB?\n", read, policed, total)  
                fmt.Printf("policer: contract=%v?\n", that)  
                mutex.Unlock()
            }
        } else if eof {
            that.Update(now)
        } else {
            // Do nothing.
        }
        
        ticks.Sleep(0)
    
    }
    
    close(output)
    
    if policed > 0 {
        t.Logf("policer: contract=%v!\n", that)
        t.Fatalf("policer: policed=%vB!\n", policed)
    }
    
    mean := float64(total) / float64(count)
    
    mutex.Lock()
    fmt.Printf("policer: end admitted=%vB policed=%vB total=%vB mean=%vB/burst maximum=%vB/burst.\n", admitted, policed, total, mean, largest)
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
        
        if (total % uint64(burst)) == 0 {
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

func ActualEventStream(t * testing.T, shape gcra.Gcra, police gcra.Gcra, supply chan byte, demand chan byte, total uint64) {
    var failure error
    var producertotal uint64 = 0
    var producerchecksum uint16 = 0
    var consumertotal uint64 = 0
    var consumerchecksum uint16 = 0
    
    fmt.Println("Beginning.")
       
    fmt.Printf("shape=%v.\n", shape);
    fmt.Printf("police=%v.\n", police);
    
    done := make(chan bool, 4)
    defer close(done)
        
    source, failure := net.ListenPacket("udp", ":5555")
    if failure != nil {
        t.Fatalf("FAILED! %v\n", failure)
    }
    defer source.Close()
    fmt.Printf("source=%+v.\n", source);
           
    sink, failure := net.ListenPacket("udp", ":0")
    if failure != nil {
        t.Fatalf("FAILED! %v\n", failure)
    }
    defer sink.Close()
    fmt.Printf("sink=%+v.\n", sink);
 
    destination, failure := net.ResolveUDPAddr("udp", "localhost:5555")
    if failure != nil {
        t.Fatalf("FAILED! %v\n", failure)
    }
    fmt.Printf("destination=%+v.\n", destination);
   
    fmt.Println("Starting.")
   
    go consumer(t, demand, &consumertotal, &consumerchecksum, done)
    go policer(t, source, police, demand, done)
    go shaper(t, supply, shape, sink, destination, done)
    go producer(t, total, supply, &producertotal, &producerchecksum, done)
    
    mutex.Lock()
    fmt.Println("Waiting.")
    mutex.Unlock()
    
    <- done
    <- done
    <- done
    <- done
       
    fmt.Println("Checking.")

    if (consumertotal == producertotal) {} else { t.Fatalf("FAILED! consumertotal=%v producertotal=%v\n", consumertotal, producertotal) }
    if (consumerchecksum == producerchecksum) {} else { t.Fatalf("FAILED! consumerchecksum=%v producerchecksum=%v\n", consumerchecksum, producerchecksum) }
    
    fmt.Printf("shape=%v.\n", shape);
    fmt.Printf("police=%v.\n", police);
   
    fmt.Println("Ending.")
}

