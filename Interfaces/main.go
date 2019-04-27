package main

import "fmt"

type Artist interface {
    // Call Play function with any object
    Play()
}

type Singer struct {
    Name string
}

type Dancer struct {
    Name string
}

func (s Singer) Play() {
    // Singer object comes here
    fmt.Printf("%s sings beautiful songs\n", s.Name)
}

func (d Dancer) Play() {
    // Dancer object comes here
    fmt.Printf("%s dances in a classical fashion\n", d.Name)
}

func main() {
    var artist1 Singer
    artist1.Name = "Kartheek"
    artist1.Play()

    var artist2 Dancer
    artist2.Name = "Vinay"
    artist2.Play()
}
