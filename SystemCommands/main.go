package main

import (
    "fmt"
    "os/exec"
    "runtime"
)

func execute() {

    // here we perform the ls command.
    // we can store the output of this in our out variable
    // and catch any errors in err
    out, err := exec.Command("ls", "-ltr").Output()
    // if there is an error with our execution
    // handle it here
    if err != nil {
        fmt.Printf("%s", err)
    }
    // as the out variable defined above is of type []byte we need to convert
    // this to a string or else we will see garbage printed out in our console
    // this is how we convert it to a string
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)

    // let's try the pwd command here
    out, err = exec.Command("pwd").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output = string(out[:])
    fmt.Println(output)
}

func main() {
    if runtime.GOOS == "windows" {
        fmt.Println("Can't Execute this on a windows machine")
    } else {
        fmt.Println("Operating System:",runtime.GOOS)
        execute()
    }
}
