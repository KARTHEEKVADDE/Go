package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)
func scanner() {
  fmt.Println("Friend! You have come inside scanner function!")
  fmt.Print("Want to go out? y/n :")
  
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    if scanner.Text() == "y"{
      fmt.Print("Bye Friend!")
      return
    }
    fmt.Println("You said : ", scanner.Text())
    fmt.Print("Want to go out? y/n :")
  }
}
func main() {
  fmt.Println("Tell Something Friend!")
  fmt.Println("---------------------")
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("-> ")
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("Oh sorry! Something went wrong! Try again!")
    }
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("hi", text) == 0 {
      fmt.Println("hello, Kartheek!")
      fmt.Println("I'm going! Keep Smiling! :)")
      break
    }
    fmt.Println("You said this : ", text, "\nAlright! Say me hi!")    
  }
  
  scanner()
}
