package main

import (
    "fmt"
    //"strconv"
)

var initCounter int

func init() {
    fmt.Println("Called First in Order of Declaration")
    initCounter++
}

func init() {
    fmt.Println("Called second in order of declaration")
    initCounter++
}

var hi = 0
func AnswerToLife() int {
  fmt.Print("hi")
    return 42
}

func init() {
    hi = AnswerToLife()
    fmt.Println(hi)
    initCounter++
}

func main() {
    fmt.Printf("Init Counter: %d\n", initCounter)
    if hi == 0 {
        fmt.Println("It's all a lie.")
    }
}
