package main

import (
    "fmt"
    "time"
    "sync"
)

// goroutines are lightweight threads managed by the Go runtime.
//
// channels are how you can send and receive values between go routines
// allows synchronization between routines without locks or other primatives
// block by default until otherside is ready
// once control returns to shell, all go routines are killed
//
// channels can be buffered aswell, kind of like io streams
// sends to buffered channels block when buffer is full
// receives block when the buffer is empty
//
// a sender can close a channel to indicate that no more values will be send
// receivers can test if a channel has been closed before trying to read
// useful in for range loops
//
// only a sender should close a channel, trying to send on a closed channel causes a panic
//
// select statement lets a routine wait on multiple operations
// blocks when none are ready
// when one can run, it executes
// select statements can have a default case as well
// allows to see if something is ready without blocking
//
// Classic Mutex
// found in sync.Mutex
// good when no communication is needed
type SafeCounter struct {
    mu sync.Mutex
    value int
}

func (counter *SafeCounter) Inc(amount int) {
    counter.mu.Lock()
    defer counter.mu.Unlock()

    counter.value += amount
}


func (counter *SafeCounter) Value() int{
    counter.mu.Lock()
    defer counter.mu.Unlock()

    return counter.value
}



func spam(s string) {
    for i := 0; i<5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func give_spam(spam_list []string, c chan string) {
    defer close(c)

    for _, r := range spam_list {
        time.Sleep(250 * time.Millisecond)
        c <- r
    }
}


func eat_spam(hunger int, c chan string) {
    time.Sleep(1000 * time.Millisecond)

    for i:=0; i<hunger; i++ {
        food, ok := <-c
        if !ok {
            fmt.Println("Out of spam!")
            return
        }
        fmt.Println(food)
    }
    fmt.Println("Full!")
}


func main() {
    // channel_tres := make(chan string, 10)
    //
    // go give_spam([]string{"bacon", "eggs", "sausages", "spam", "spam", "spam"}, channel_tres)
    // eat_spam(8, channel_tres)

    my_count := SafeCounter{value:0}

    for i:=0; i<1000; i++ {
        go my_count.Inc(1)
    }

    time.Sleep(time.Second)
    fmt.Println(my_count.Value())
    

}

