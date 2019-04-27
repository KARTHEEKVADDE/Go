package main

import (
    "fmt"
    "strconv"
)

func strToIntConversion() {
    fmt.Println("String to Integer Value Conversion")

    var ourInteger int
    // use the strconv package to convert our string '12345' to an integer value
    ourInteger, err := strconv.Atoi("12345")

    // if there has been an error then handle it here
    if err != nil {
        fmt.Println(err)
    }

    // this should print out 12345
    fmt.Println(ourInteger)
}

func intToStringConversion() {
    fmt.Println("integer to string conversion")

    var ourString string

    ourString = strconv.Itoa(12345)

    // print out our string value
    fmt.Println(ourString)

}

func main() {
    strToIntConversion()
    intToStringConversion()
}
