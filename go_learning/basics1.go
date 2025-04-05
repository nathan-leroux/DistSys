package main

import (
    "fmt"
    "math"
)


// exported names are capitalised
// if they're not they are not visible
func exported_names()  {
    fmt.Println(math.Pi)
}
// functions can take zero or one arguments
func add(x int, y int) int {
    return x+y
}

// when two or more consecutive function parameters you can omit the type til the end
func sneaky_add(x, y int) int {
    return x + y
}

// functions can return multiple results
func multi_add(x, y int) (int, int) {
    return x+y, x*y
}


// functions can name return values
// this is mainly done to document the meaning of the return values
// return statement without args returns the named returns
// known as a naked return
func split(sum int) (x, y int) {
    x = sum*4 / 9
    y = sum -x
    return 
}

//multiple variables of the same type can be declared and instantiated on the same line
// if the variable is intialised, you can skip the type
// := is the short variable operator, replacement for var, can't be used in global scope
func initial() {
    var a, b int
    var c, d = "balls", true
    e, f := 3, "four"
    fmt.Println(a, b, c, d, e, f)
}


//types
// mostly pretty standard
// int (unsigned and signed up to 64 bit)
// uintptr (guess its the size of a mem address)
// bool (true, false)
// string
// float32 float64
// complex64

// empty values
// 0, false and ""

// conversions are required to be explicit
func type_convert() {
    var i int = 32
    var f float64 = float64(i)
    var u uint = uint(f)

    fmt.Println(i, f, u)
}


// constants can be defined with the 'const' keyword
func def_const() {
    const mah_neck = "my back"

    fmt.Println(mah_neck)
}


func main()  {
    def_const()
}
