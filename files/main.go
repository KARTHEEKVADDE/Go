package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    fmt.Println("Hey Kartheek! What's the Topic?")
    fmt.Println("Read/Write Operations on Files", "Let's start!")
    // Create Bytes Array
    fmt.Println("***Prepared my Data")
    mydata := []byte("All the data I wish to write to a file\n")
    // write your data to file(file gets created if not found)
    // WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.data", mydata, 0777)

    if err != nil {
        fmt.Println(err)
    }else{
        fmt.Println("***Edited the file with prepared data!")
    }
    // Read the file
    data, err := ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }else{
        fmt.Println("***Reading the file!")
    }
    // Print it out
    fmt.Println(string(data))
    // Open the file in Write mode to append something
    f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }else{
        fmt.Print("***Adding something to the file!")
    }
    defer f.Close() // let file be opened till the end
    // Append something
    if p, err := f.WriteString("new data that wasn't there earlier\n"); err != nil {
        panic(err)
    }else{
      fmt.Println("--Insert Count:", p,"characters--Done!")
    }

    data, err = ioutil.ReadFile("myfile.data")
    if err != nil {
        fmt.Println(err)
    }else{
      fmt.Println("***Check it now!")
    }
    // Check the content again
    fmt.Print(string(data))
    fmt.Println("***File successfully changed!")
}
