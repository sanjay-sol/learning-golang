package main

import (
  "fmt"
  "time"
  "sync"
)


func work(timer time.Duration, s string, ch chan string, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Println("Started Work.......")
  time.Sleep(timer)
  fmt.Println("Done work !!!1")
  ch <- s + "Hello"
}

func main() {
  start := time.Now()
  var wg sync.WaitGroup
  ch := make(chan string)
  wg.Add(2)

  go work(2 * time.Second, "work--1", ch, &wg) 
  go work(4 * time.Second, "work--2", ch, &wg)

  go func() {
    wg.Wait()
    close(ch)
  }()

  for val := range ch{
    fmt.Printf("val from channel %v", val)
  }

  fmt.Printf("Total time is ------ %v ",time.Since(start))

}
