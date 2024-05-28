
// go routines 

package main

import (
  "fmt"
  // "time"
)

// func func1(n int){
//   fmt.Println(n)
//   time.Sleep(1 * time.Second)
// }
//
func main() {
  // go channels 
  
  myChannel := make(chan string)
  
  go func() {
    myChannel <- "Hello"
    myChannel <- "World"
    close(myChannel)

  }()

  // printing both values

  // msgs := <- myChannel
  for msg := range myChannel {
    fmt.Println(msg)
  }
  // fmt.Println(msgs) 

  // go func1(2)

  // go func1(3)
  // go func1(4)
  // go func1(5)
  // 
  //
  // // loop {
  // //   time.Sleep(1 * time.Second)
  // // }
  //
  //
  //
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello3")
  // fmt.Println("Hello4")
  // fmt.Println("Hello5")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello")
  // fmt.Println("Hello222")
  //
}
