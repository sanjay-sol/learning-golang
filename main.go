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
)

type Person struct {
  name string
  age int
}

func newPerson(name string, age int) *Person {
  p := Person{name:            name , age: 433}
                     return &p
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
  
  
  

}



