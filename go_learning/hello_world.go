// this is a comment
// curiously, golang is opinionated against shebangs

package main

import (
    "fmt"
    "time"
)

func time_now() {
    fmt.Println("Welcome to the ghetto playground")
    fmt.Println("The time is", time.Now())
}

func main() {
    fmt.Println("Hello world")
    time_now()
}


