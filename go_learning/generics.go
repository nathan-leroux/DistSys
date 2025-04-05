package main

import (
    "fmt"
)


// functions can be written to work on multiple types with type parameters
// type params apper between brackets, before the arguments
//
// comparable is a built in
func find_index[T comparable](s []T, x T) int{
    for i, v := range s {
        if v == x {
            return i
        }
    }

    // index not found
    return -1
}

// along with generic functions, go also supports generic types
// types can be parameterized, allowing generic data structures
// an example would be a linked list allowing any data type
type Llist[T any] struct {
    next *Llist[T]
    val T
}

func main() {
    // works on ints
    int_arr := []int{10, 20, 30, 40, 50}
    fmt.Println(find_index(int_arr, 30))

    // works on strings
    str_arr := []string{"mama", "dear", "come", "over", "here"}
    fmt.Println(find_index(str_arr, "dear"))
    
}

