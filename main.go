// package main
//
//
// import (
//   "os"
//   "log"
//   "readLines"
//   "compareFiles"
//   "github.com/sanjay-sol/learning-go/file"
// )
//
// func main() {
//     if len(os.Args) < 3 {
//         log.Fatalf("Usage: %s <file1> <file2>\n", os.Args[0])
//     }
//
//     file1Path := os.Args[1]
//     file2Path := os.Args[2]
//
//     file1Lines, err := readLines(file1Path)
//     if err != nil {
//         log.Fatalf("Failed to read file %s: %s\n", file1Path, err)
//     }
//
//     file2Lines, err := readLines(file2Path)
//     if err != nil {
//         log.Fatalf("Failed to read file %s: %s\n", file2Path, err)
//     }
//
// }
//
//


package main 

import (
  "fmt"
  "time"
  // "github.com/sanjay-sol/learning-golang/hello"  
)

   type   Person struct {
  name string
  age int
}

type rect struct {
  width int
  height int
  

}

type Describe interface {
  Describer() string
}

func (r rect) Describer() string {
  
 return fmt.Sprintf("%v , %v", r.width, r.height)
}

func (r *rect) area() int {
  return r.width * r.height
}

func (r * rect) peri() int {
  return 2*r.width + 4*r.height
}

funcs newPerson(name string, ageint) *Person {
  p := Person{name: name , age: 433}
  return &p
}


// go channel syscronization


func worker(done chan string)  {
  fmt.Println("wrking....")
  time.Sleep(time.Second)
  fmt.Println("Done")

  done <- "Hello guys"
}

// channel directions


func ping2(ping chan<- string, msg string ){
  ping <- msg
}


func pong2(ping <-chan string, pong chan<- string){
  msg := <- ping
  pong <- msg
}


// Done channel 



func donechan(done <-chan string) {
 npm install `@gardenfi/core` 
  for {
    select {
    case <-done:
      return
    default:
      fmt.Println("Doing work until done ................")
  }
  }

  

}

func main() {
  main2()
//   p1 := Person{name: "Sanjay", age: 25}
//   p2 := Person{name: "Sanjay", age: 25}
//
//
//
//   if p1 == p2 {
//     fmt.Println("p1 and p2 are equal")
//   } else {
//     fmt.Println("p1 and p2 are not equal")
//   }
//
//   p3 := newPerson("sanjuu",69)
//
//   fmt.Println(p3);
//   fmt.Println(p3.name);
//   fmt.Println(p3.age);
//   
//   
//
//   r := rect{height : 10, width :5}
//
//   fmt.Println(r.area())
//   fmt.Println(r.peri())
//
//   r1 := &r
//
//   r1.height = 11
//   fmt.Println(r1.area())
//   fmt.Println(r1.peri())
//
//   r2 := rect{height: 10, width: 5}
//   
//   // var d Describe
//
//   d := r2
//
//   fmt.Println(d.Describer())
//
//   // messages := make(chan int)
//   //
//   // go func() {messages <- 11}()
//   // 
//   // go func() {
//   //
//   // for i :=0;i<1000;i++ {
//   //   
//   //   messages <- i
//   // }
//   //
//   //   close(messages)
//   // }()
//   // 
//   //   for val := range messages {
//   //     fmt.Println(val)
//   //   }
//
//
//   // sync channel 
//
//   syncchan := make(chan string,1)
//
//   go worker(syncchan)
//   
//   fmt.Println(<- syncchan)
//
//   fmt.Println("Done sync....")
//
//   // channel directions
//   ping := make(chan string,1)
//   pong := make(chan string,1)
//
//   ping2(ping,"hello from the channel directions ")
//   pong2(ping, pong)
//
//   fmt.Println(<-pong)
//
//   // select 
//
//   chan1 := make(chan string)
//   chan2 := make(chan string)
//   
//   go func() {
//     time.Sleep(3*time.Second)
//     chan1 <- "wait for 3 sec"
//   }()
//
//   go func() {
//     time.Sleep(time.Second)
//     chan2 <- "wait for 1 sec"
//   }()
//
//
//
//
//
//   for i := 0;i<2;i++ {
//     select {
//     case msg := <- chan1:
//       fmt.Println("message from channel 1 %v", msg)
//       
//     case msg2 := <- chan2:
//       fmt.Println("message from channel 2 ", msg2)
//     }
//   }
//
//   
//   // testing range 
//   
//   // mapex := make(map[string]string)
//   //
//   // for i:=;i<5;i++ {
//   //   map[i+""] = 'a'+i+""
//   // }
//
//  // fmt.Println(mapex)
//   
//
//   // Packages testing
//
//   hello.Hello()
//
//
//   // for select 
//
//   chan11 := make(chan string, 5)
//
//   chars := []string{"a", "b","c","d","e"}
//
//   for _, val := range chars {
//     select {
//     case chan11 <- val :
//   }
//   }
//
//   close(chan11)
//
//   for val := range chan11 {
//     fmt.Println(val)
//   }
//
//
// boolchan := make(chan string)
//
//
//   go donechan(boolchan)
//
//   time.Sleep(time.Second * 3)
//
//   // boolchan <- "true"
//
//   close(boolchan)
//
    newchann := make(chan string)

    go func(newchann <-chan string) {
        for {
            select {
            case val, ok := <-newchann:
                if !ok {
                    fmt.Println("Channel closed, exiting.")
                    return
                }
                fmt.Println("Received:", val)
            default:
                fmt.Println("No new messages. Waiting...")
                time.Sleep(1 * time.Second)
            }
        }
    }(newchann)


    for i := 0; i < 3; i++ {
        newchann <- fmt.Sprintf("Message %d", i)
        time.Sleep(time.Second * 1)
    }

    for j := 2; j > 0; j-- {
        time.Sleep(time.Second * 1)
        newchann <- fmt.Sprintf("Special message %d", j)
    }

    close(newchann)
    time.Sleep(time.Second * 2) 
   

    /// testing select with multiple channels

    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
      time.Sleep(time.Second)
      ch1 <- "one"
    }()

    go func() {
      time.Sleep(time.Second)
      ch2 <- "two"
    }()

    for i := 0; i < 2; i++ {
      select {
      case msg1 := <- ch1:
        fmt.Println("message 1", msg1)
      case msg2 := <- ch2:
        fmt.Println("message 2", msg2)
      }
    }

    // select with default case

    ch3 := make(chan string)
    ch4 := make(chan string)

    go func() {
      time.Sleep(time.Second)
      ch3 <- "one"
    }()

    go func() {
      time.Sleep(time.Second)
      ch4 <- "two"
    }()

    for i := 0; i < 2; i++ {
      select {
      case msg1 := <- ch3:
        fmt.Println("message 1", msg1)
      case msg2 := <- ch4:
        fmt.Println("message 2", msg2)
      default:
        fmt.Println("No message received")
      }
    }

  // wait groups 





}



