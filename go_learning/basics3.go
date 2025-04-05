package main

import (
    "fmt"
)

// ah good ol pointers
// zero value is nil
// no pointer arithmetic, less seg faults :)
func pointers() {
    var p *int
    i := 1 

    p = &i
    *p++

    // prints 2
    fmt.Println(i)
}

// structs are a collection of fields
// instantiated with {}
// fields accessed with .
// can write pointers to structs
// you don't have to dereference to access fields, its implicit
//
// struct literals make use of Name: to specify fields to populate, all others are nil
type Vertex struct {
    x_val int
    y_val int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x_val: 1}  // y_val:0 is implicit
	v3 = Vertex{}      // x_val:0 and y_val:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func structs() {
    var new_vert Vertex
    p := &new_vert

    new_vert = Vertex{1,2}
    fmt.Println(new_vert)
    fmt.Println(new_vert.y_val)
    fmt.Println(p.x_val)
}

// defined by [n]T
// n is number of values
// T is type
// arrays cannot be resized
//
// slices are views of arrays
// like python [low:high]
// slice litterals can be defined with {} after definition
//
// len() and cap() are built in functions to messure slices
// len() is current length, cap() is underlying array len
func arrays() {
    var a [3]string
    var s []string
    s2 := []bool{true, false, true}

    a[0] = "knicks"
    a[1] = "in"
    a[2] = "three"

    fmt.Println(len(a))

    s = a[0:2]
    

    fmt.Println(a[0], a[1], a[2])
    fmt.Println(s)
    fmt.Println(s2)
}

// make is how you create dynamic arrays
// malloc vibes
// returns slice that refers to that array
// first num arg is len(), second is cap()
// cap() is optional
// append(slice ,extras of same type) can extend arrays
// if the underlying array is too small, new one created and returned
//
// slice literals are array literals without the length
// just need the type declared.
func dyn_array() {
    a := make([]int, 5, 10)

    // b := []int{1, 2, 3, 4}

    a = append(a, 1, 2, 3)
    fmt.Println(a)
}

// range keyword used to iterate over iterables
// by default first variable will be index
// second value
// either can be omitted with _
func for_range() {
   var powers = []int{1, 2, 4, 8, 16, 32, 64}

   for i, v := range powers{
       fmt.Printf("2**%d = %d\n", i, v)
   }

   for _, j := range powers {
       fmt.Printf("keepin it simple: %d\n", j)
   } 
}

// maps
// dictionary equivalent
// 
// can be done as a literal, but keys are required
//
// can be mutated in the ways that you would expect
// to check if a key exists two value assignment is used
func maps() {
    var m map[string]int

    m = make(map[string]int)
    m["one"] = 1

    var m_litteral = map[string]int {
	"one": 1,
	"two": 2,
	"three": 3,
    }

    fmt.Println(m_litteral["two"])
    fmt.Println(m["one"])

    m["two"] = 3
    delete(m, "one")


    fakie, ok := m["not a key"]
    if ok {
	fmt.Println(fakie)
    } else {
	fmt.Println("not ok")
    }
}

// functions are first class
// yippee
func shadow_wizard(input_func func(int) int) int{
    return input_func(69)
}

func first_class(x int) int{
    return 2*x
}




func main() {
    // fmt.Println(shadow_wizard(first_class))
    arrays()
}


