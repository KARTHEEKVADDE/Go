package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  fmt.Println("Tell Something Friend!")
  fmt.Println("---------------------")
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("hi", text) == 0 {
      fmt.Println("hello, Kartheek!")
      fmt.Println("I'm going! Keep Smiling! :)")
      break
    }
    fmt.Println("You said this : ", text, "\nAlright! Say me hi!")    
  }
}
