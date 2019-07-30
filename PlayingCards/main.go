package main

import "fmt"
import "math"
import "math/rand"

func main() {
  fmt.Println("Hello World")
  var list[]int
  for i:=1;i<=4;i++{
    for j:=1;j<=13;j++{
      list = append(list, j)
    }
  }
  rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
  Game:
    Game(list)
    fmt.Print("Want to play once more! No(0)/Yes(1) :")
    var oncemore int
    fmt.Scan(&oncemore)
    if oncemore == 1{
      goto Game
    }
  fmt.Print("Good Bye! :)")
}
func Game(list []int){
  var Target = 21
  var Sum1, Sum2 int
  for i:=0; i<100; i++{
    fmt.Println("Player1 - Pick The Card : 1 - ", len(list))
    var choice1 int
    fmt.Scan(&choice1)
    Sum1 += list[choice1]
    list = append(list[:choice1], list[choice1+1:]...)
    rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
    fmt.Println("Player1 - Your Score : ", Sum1)
    fmt.Println("You Want to Hit(0) / Stand(1) ? : ")
    var hit_stand int
    fmt.Scan(&hit_stand)
    if hit_stand == 1{
      WinLose(Sum1, Sum2, Target)
      break
    }
    fmt.Println("Player2 - Pick The Card : 1 - ", len(list))
    var choice2 int
    fmt.Scan(&choice2)
    Sum2 += list[choice2]
    list = append(list[:choice2], list[choice2+1:]...)
    rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
    fmt.Println("Player2 - Your Score : ", Sum2)
    fmt.Println("You Want to Hit(0) / Stand(1) ? : ")
    var hit_stand2 int
    fmt.Scan(&hit_stand2)
    if hit_stand2 == 1{
      WinLose(Sum1, Sum2, Target)
      break
    }
  }
}
func WinLose(Sum1, Sum2 int, Target int){
  if math.Abs(float64(Sum1-Target)) < math.Abs(float64(Sum2-Target)){
    fmt.Println("Player1 is the Winner! :)\nPlayer2 is the Loser!  ):")
  }else{
    fmt.Println(("Player2 is the Winner! :)\nPlayer1 is the Loser! :("))
  }
}
