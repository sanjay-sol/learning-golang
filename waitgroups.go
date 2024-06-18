package main

import (
  "fmt"
  "time"
)


func work(timer time.Duration, s string, ch chan string) {
  fmt.Println("Started Work.......")
  time.Sleep(timer)
  fmt.Println("Done work !!!1")
  ch<- s
}

func main() {
  start := time.Now()
  // wg := sync.WaitGroup{}
  ch := make(chan string)
  // wg.Add(2)

  go work(2 * time.Second, "work--1", ch)
  go work(4 * time.Second, "work--2", ch)

  // wg.Wait()
  val1 := <-ch
  fmt.Printf("val1 is ------ %v ",val1)
  val2 := <-ch
  fmt.Printf("val2 is ------ %v ", val2)
  fmt.Printf("Total time is ------ %v ",time.Since(start))

}
