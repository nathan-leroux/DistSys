package main

import (
    "fmt"
    "runtime"
)

// classic structure, parenthesis are not required
// pre and post statements are optional
// no while loop as for becomes functionally the same
// 
// all statements are optional
// omitting all makes an infinite loop
func for_loop() {
    sum := 0
    j := 0

    for i:=0; i<10; i++ {
        sum += 1
    }

    for ; j<10 ; {
        sum += 1
        j++
    }

    // for {
    //     fmt.Println("...Infinite, This loop is ...")
    // }

    fmt.Println(sum)
}


// if statements also don't require parenthisis
// can sneak in a pre statement
// the prestatement exists only in scope of the if
// carries over to else statements too
func if_state() {
    x := true

    if x == true {
        fmt.Println("never lied")
    }

    if y:=2; y>1 {
        fmt.Println("not even once")
    } else {
        fmt.Println("ok maybe once", y)
    }

    // undefined y
    // fmt.Println(y)
}

// standard, no continuing after a successful case was found
// can be used with pre-statement
// can be used without any statement, effectively makes it if-else chain
func switch_state() {
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X")
    case "linux":
        fmt.Println("nerdville")
    default:
        fmt.Printf("%s.\n", os)
    }
}

// defers execution of a function until the surrounding function returns
// commonly used to pack up resources like closing files and releasing mutex's
// works by placing functions on a stack (last in first out)
func defer_state() {
    defer fmt.Println("my back")

    fmt.Println("my neck")
}


func main() {
    defer_state()
    fmt.Println("my pussi")
}
