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
)

type Person struct {
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
  return 2*r.width + 2*r.height
}

func newPerson(name string, age int) *Person {
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


func main() {
  p1 := Person{name: "Sanjay", age: 25}
  p2 := Person{name: "Sanjay", age: 25}



  if p1 == p2 {
    fmt.Println("p1 and p2 are equal")
  } else {
    fmt.Println("p1 and p2 are not equal")
  }

  p3 := newPerson("sanjuu",69)

  fmt.Println(p3);
  fmt.Println(p3.name);
  fmt.Println(p3.age);
  
  

  r := rect{height : 10, width :5}

  fmt.Println(r.area())
  fmt.Println(r.peri())

  r1 := &r

  r1.height = 11
  fmt.Println(r1.area())
  fmt.Println(r1.peri())

  r2 := rect{height: 10, width: 5}
  
  // var d Describe

  d := r2

  fmt.Println(d.Describer())

  // messages := make(chan int)
  //
  // go func() {messages <- 11}()
  // 
  // go func() {
  //
  // for i :=0;i<1000;i++ {
  //   
  //   messages <- i
  // }
  //
  //   close(messages)
  // }()
  // 
  //   for val := range messages {
  //     fmt.Println(val)
  //   }


  // sync channel 

  syncchan := make(chan string,1)

  go worker(syncchan)
  
  <- syncchan

  fmt.Println("Done sync....")


  

}



