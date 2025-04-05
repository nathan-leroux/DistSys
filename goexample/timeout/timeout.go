package main

import (
    "fmt"
    "time"
)

// if 'timeout' happens, reset back to zero
type Listing struct {
    count int
}

// block for a certain amount of time, if allowed to finish, subtract 1 from Listing
func timer(timeout chan bool) {

    fmt.Println("timer sleeping")
    time.Sleep(time.Second *2)
    fmt.Println("timeout reached")
    timeout<- true
}

func main() {
    l := Listing{1}

    // start the timeout
    timeout := make(chan bool)
    go timer(timeout)

    fmt.Println("main sleeping")
    time.Sleep(time.Second * 3)

    select {
    case <-timeout:
        fmt.Println("timeout found")
        l.count = 2

    default:
        fmt.Println("no timeout")
        l.count = 0
    }

    fmt.Printf("final val: %v\n",l.count)
}
