package main

import "fmt"

// func update(name **string ) {
//   **name = "aaa"


// }

type Person2 struct{
  name string
  age int
}

func (p *Person2) change() *Person2 {
  p.name = "aaa"
  p.age = 1111
  return p
}
func (p *Person2) change2() Person2 {
  p.name = "bbb"
  p.age = p.age
  return *p
}




func main()  {
  person := Person2{
    name: "sanjay",
    age: 23,
  }
  p := person.change2();
  fmt.Println(p); 
  fmt.Println(person);
  
}
