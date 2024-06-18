package main

import (
  "fmt"
  "time"

  "sync"
)


func work(timer time.Duration, wg * sync.WaitGroup, s string, ch chan string) string {
  fmt.Println("Started Work.......")
  time.Sleep(timer)
  fmt.Println("Done work !!!1")
  wg.Done()
  ch<- s
}

func main() {
  start := time.Now()
  wg := sync.WaitGroup{}
  ch := make(chan string)
  wg.Add(2)

  val1 := go work(2 * time.Second, &wg. "work22", ch)
  val2 := go work(4 * time.Second, &wg, "workkk1", ch)

  wg.Wait()
  fmt.Printf("val1 is ------ %v ",val1)
  fmt.Printf("val2 is ------ %v ", val2)
  fmt.Printf("Total time is ------ %v ",time.Since(start))

}
