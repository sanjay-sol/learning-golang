package main

import "fmt"

func update(name **string ) {
  **name = "aaa"


}

type person struct{
  name string
  age int
}

func change(p *person) {
  p.name = "asdf"
  p.age = 99
}

func main()  {
  num := "sanju"
  num2 := &num;
  num3 := &num2;
  fmt.Println(num);
  fmt.Println(*num2);
  fmt.Println(**num3);
  update(num3);
  fmt.Printf(num);
  person := person{
    name: "sanjay",
    age: 23,
  }
  change(&person);
  fmt.Println(person) 
  
}
