package main

import (
    "fmt"
    "math"
)

// go does not have classes.
// methods can be defined on types however
//
// a method is a function with a special receiver argument
// in the below example, the 'Abs' method has the reciever of type Point
// methods are still functions, just with syntactic sugar
//
// methods can only be declared on types defined in the same package as the method.
// methods can be defined on pointer recievers too.
// pointer recievers are able to modify underlying values, this makes them more common than value ones
//
// go as a convience allows you to use a value for both value and pointer recievers
// in general, all methods on a type should have either value or pointer receivers, but not both.
type Point struct {
    X, Y float64
}

type Line struct {
    l1 Point
    l2 Point
}

// expand written as a regular function
func Expand(v Point, f float64) {
    v.X = v.X + f
    v.Y = v.Y + f
}


// Abs written as a member function, by value
func (v *Point) Abs() float64{
    if v == nil {
        fmt.Println("bing bong")
        return 0
    }

    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale written as a member function by reference
// member functions need to be written by reference in order to modify the type's values
// more common than value member functions
func (v *Point) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}


func (l *Line) Length() float64{
    l.l1.Y = 20

    return (l.l1.X - l.l2.X) * (l.l1.Y - l.l2.Y)
}
// 
// Interfaces
//
// An interface type is defined as a set of method signatures
// a value of an interface type can hold any value that implements those methods.
type MyInter interface {
    Abs() float64
}

// a type implements an interface if it implements all its methods, there is no 'implements' keywork
// so they are done implicitly
//
// interface methods can be called when they have been fed nil pointers.
// interface methods called when the interface has not been instantiated at all are a runtime error
// 'Stringer' is very common interface defined in 'fmt' that is used to represent the string form of types
func inter_test() {
    var i MyInter
    var P *Point

    // i.Abs() gives runtime error

    // nil implementation
    i = P
    fmt.Println(i.Abs())

    // not nil
    i = &Point{3, 4}
    fmt.Println(i.Abs())
}

// The empty interface
// can hold any val (any val implements no methods)
// used by code that handles unknown types
func empty_interface() {
    var i interface{}
    i = 21

    fmt.Println(i)
}


// Type assertions
// provides a way to check an interfaces type
// type switches allow serveral type assertions in series
// allows handling of multiple options when the type is unknown
func t_asserts() {
    var i interface{} = "balls"

    // asserts string, or panic
    s := i.(string)
    fmt.Println(s)
    
    // asserts string, no panic
    s, ok := i.(string)
    fmt.Println(s, ok)
    
    // asserts float, no panic
    f, ok := i.(float64)
    fmt.Println(f, ok)
}

func t_switch(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("value is str\n")
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Println("dunno bout that one chief")
    }
}



func main() {
    l := new(Line)

    l.l1 = Point{1,2}
    l.l2 = Point{3,4}

    fmt.Println(l)
    l.Length()
    fmt.Println(l)
}

