package main

import "fmt"

func main(){
  if len(os.Args) > 1 {
    fmt.Println("Hello, " + os.Arg[1])
  } else {
    fmt.Println("Hello, World")
  }
}
