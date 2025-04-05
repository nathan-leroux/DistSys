package main

import (
    "fmt"
    "example/greetings"
    "log"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("Big Dawg")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}


