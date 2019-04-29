package main
import ("fmt"
        "os")
type user struct{
  name string
  mobNo int
  pwd string
  amt int
}
var data = make(map[int]*user)
func main(){
  welcome()  
}
func welcome(){
  fmt.Println("Welcome To Mini Bank")
  fmt.Println("Human! Choose your Option!")
  fmt.Println("1.Sign Up(New Customer)")
  fmt.Println("2.Log In(Existing Customer)")
  fmt.Println("3.Exit")
  var a int
  fmt.Scanf("%d",&a)
  fmt.Print(a)
  if a == 1{
    signUp()
  }else if a == 3 {
    exit()
  }else{
    logIn()
  }
}
func signUp(){
  var acNo, mobNo int; var name, pwd, cfPwd string
  fmt.Println("Please enter your account number!")
  fmt.Scanf("%d", &acNo)
  fmt.Println("Please enter your name!")
  fmt.Scanf("%s", &name)
  fmt.Println("Please enter your mobile number!")
  fmt.Scanf("%d", &mobNo)
  password:
    fmt.Println("Please create your password!")
    fmt.Scanf("%s", &pwd)
    fmt.Println("Please enter your password again!")
    fmt.Scanf("%s", &cfPwd)
    if pwd == cfPwd{
      createProfile(acNo, mobNo, name, pwd)
    }else{
      fmt.Print("Passwords didn't match!!!")
      goto password
    }
  fmt.Println("3.Exit")
}
func logIn(){
  var acNo int; var pwd string
  login:
    fmt.Println("Please enter your account number!")
    fmt.Scanf("%d", &acNo)
    fmt.Println("Please enter your password!")
    fmt.Scanf("%s", &pwd)
    fmt.Print(data[acNo], data[acNo].name)
    if data[acNo].pwd == pwd{
      fmt.Println("You have logged in successfully!")
      userProfile(acNo)
    }else{
      fmt.Println("Please enter correct credentials!")
      goto login
    }
}
func createProfile(acNo, mobNo int, name, pwd string){
  amt := data[acNo].amt
  data[acNo] = &user{name, mobNo, pwd, amt}
  fmt.Print(data)
  fmt.Println("Profile Created/Modified Successfully!")
  fmt.Println("Redirecting to Welcome Page!")
  welcome()
}
func userProfile(acNo int){
  fmt.Println("Hey Human! Welcome to the User Profile!")
  fmt.Println("User Name:", data[acNo].name)
  fmt.Println("Account Number:", acNo)
  fmt.Println("Mobile Number:", data[acNo].mobNo)
  fmt.Println("Password:", data[acNo].pwd)
  fmt.Println("Balance:", data[acNo].amt)
  fmt.Println("Choose your Option!")
  fmt.Println("1.Update Your Details 2.Deposit Money 3.Withdraw Money 4.Transfer Money 5.Log Out")
  var a int
  fmt.Scanf("%d", &a)
  if a==1{
    signUp()
  }else if a==2{
    deposit(acNo)
  }else if a==3{
    withdraw(acNo)
  }else if a==4{
    transfer(acNo)
  }else{
    welcome()
  }
}
func exit(){
  os.Exit(1)
}
func deposit(acNo int){
  fmt.Print("How much? ")
  var howMuch int
  fmt.Scanf("%d", &howMuch)
  data[acNo].amt += howMuch
  fmt.Println(data, data[acNo])
  userProfile(acNo)
}
func withdraw(acNo int){
  fmt.Print("How much? ")
  var howMuch int
  fmt.Scanf("%d", &howMuch)
  data[acNo].amt -= howMuch
  fmt.Println(data, data[acNo])
  userProfile(acNo)
}
func transfer(acNo int){
  fmt.Print("To Whose Account? ")
  var toWhom int
  fmt.Scanf("%d", &toWhom)
  fmt.Print("How much? ")
  var howMuch int
  fmt.Scanf("%d", &howMuch)
  data[acNo].amt -= howMuch
  data[toWhom].amt += howMuch
  fmt.Println(data, data[acNo])
  userProfile(acNo)
}
