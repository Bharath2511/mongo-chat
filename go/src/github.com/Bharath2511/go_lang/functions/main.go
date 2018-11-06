package main

import "fmt"

func gone(name string) string {
	return "Hello " + name
}

func sum(num1 int,num2 int) int {
	return num1+num2
}


func main() {
  fmt.Println(sum(2,2))	
  fmt.Println(gone("bharath"))
}