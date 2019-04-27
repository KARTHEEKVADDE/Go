package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    //"strconv"
)

// Users struct which contains
// an array of users
type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, er := ioutil.ReadAll(jsonFile)
    fmt.Println(er)

    // we initialize our Users array
    var users Users

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their social urls
    // as just an example
    fmt.Println(users)    
    n := users.Users
    for i := 0; i < len(n); i++ {
        fmt.Println("User Name: " + n[i].Name)
        fmt.Println("User Type: " + n[i].Type)
        fmt.Println("User Age: ", int(n[i].Age)) //strconv.Itoa(n[i].Age)
        fmt.Println("Facebook Url: " + n[i].Social.Facebook)
        fmt.Println("Twitter Url: " + n[i].Social.Twitter)
    }
    //Another way using Map with Empty Interface
    var result map[string]interface{}
    json.Unmarshal(byteValue, &result)

    fmt.Println(result["users"])
}
