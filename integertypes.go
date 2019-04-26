package main

import "fmt"

func main() {
  fmt.Println("Hi Kartheek! :)\n")
  fmt.Println("Can you tell the integer types!\n")
  fmt.Println("Of course! Let's start!\n")

  // unsigned int with 8 bits
  // Can store: 0 to 255
  var unsignedInt8 uint8
  fmt.Println("\n unsignedInt8 - Default Value : ", unsignedInt8)
  
  // signed int with 8 bits
  // Can store: -127 to 127
  var signedInt8 int8
  fmt.Println("\n signedInt8 - Default Value : ", signedInt8)

  // unsigned int with 16 bits
  var unsignedInt16 uint16
  fmt.Println("\n unsignedInt16 - Default Value : ", unsignedInt16)

  // signed int with 16 bits
  var signedInt16 int16
  fmt.Println("\n signedInt16 - Default Value : ", signedInt16)

  // unsigned int with 32 bits
  var unsignedInt32 uint32
  fmt.Println("\n unsignedInt32 - Default Value : ", unsignedInt32)

  // signed int with 32 bits
  var signedInt32 int32
  fmt.Println("\n signedInt32 - Default Value : ", signedInt32)

  // unsigned int with 64 bits
  var unsignedInt64 uint64
  fmt.Println("\n unsignedInt64 - Default Value : ", unsignedInt64)

  // signed int with 64 bits
  var signedInt64 int64
  fmt.Println("\n signedInt64 - Default Value : ", signedInt64)
}
