package main

import "fmt"

func update(name **string ) {
  **name = "aaa"


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

}

type person struct{
  name string
  age int
}
