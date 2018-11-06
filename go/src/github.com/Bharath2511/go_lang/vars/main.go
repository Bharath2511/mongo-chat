package main

import "fmt"

func main() {
  var age = 23
  var name = "Bharath"
  var isCool = true
  isCool = false
  fmt.Println(name,age)
  //for getting the type
  fmt.Printf("%T\n", name)
  fmt.Printf("%T\n", age)
  fmt.Printf("%T\n", isCool)
}