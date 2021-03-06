package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

// our struct which contains the complete
// array of all Users in the file
type Users struct {
    XMLName xml.Name `xml:"users"`
    Users   []User   `xml:"user"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type User struct {
    XMLName xml.Name `xml:"user"`
    Type    string   `xml:"type,attr"`
    Name    string   `xml:"name"`
    Social  Social   `xml:"social"`
}

// a simple struct which contains all our
// social links
type Social struct {
    XMLName  xml.Name `xml:"social"`
    Facebook string   `xml:"facebook"`
    Twitter  string   `xml:"twitter"`
    Youtube  string   `xml:"youtube"`
}

func main() {

    // Open our xmlFile
    xmlFile, err := os.Open("users.xml")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.xml")
    // defer the closing of our xmlFile so that we can parse it later on
    defer xmlFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(xmlFile)

    // we initialize our Users array
    var users Users
    // we unmarshal our byteArray which contains our
    // xmlFiles content into 'users' which we defined above
    xml.Unmarshal(byteValue, &users)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    n := users.Users
    fmt.Println(users, "\n------")
    fmt.Println(n, "\n------")

    for i := 0; i < len(n); i++ {
        fmt.Println("User Type: " + n[i].Type)
        fmt.Println("User Name: " + n[i].Name)
        fmt.Println("Facebook Url: " + n[i].Social.Facebook)
        fmt.Println("Twitter Url: " + n[i].Social.Twitter)
        fmt.Println("Youtube Url: " + n[i].Social.Youtube)
    }

}
