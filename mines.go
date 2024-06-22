// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )
//
// var matrix [5][5]int32
//
// func calran() int {
// 	return rand.Intn(5)
// }
//
// func calculate(val int) {
// 	rand.Seed(time.Now().UnixNano())
// 	vallocal := val
// 	for vallocal > 0 {
// 		num1 := calran()
// 		num2 := calran()
// 		if matrix[num1][num2] == 0 {
// 			matrix[num1][num2] = 1
// 			vallocal--
// 		}
// 	}
// }
//
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	var num1 int
//
// 	for {
// 		fmt.Print("Enter the number of mines (1 - 24): ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
//
// 		var err error
// 		num1, err = strconv.Atoi(input)
// 		if err == nil && num1 > 0 && num1 <= 24 {
// 			break
// 		}
// 		fmt.Println("Invalid input. Please enter a number between 1 and 24.")
// 	}
//
// 	calculate(num1)
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			value := 5*i + j + 1 
// 			if matrix[i][j] == 1 {
// 				formatted := fmt.Sprintf(" -(%2d) ", value)
// 				fmt.Print(formatted)
// 			} else {
// 				formatted := fmt.Sprintf(" -(%2d) ", value)
// 				fmt.Print(formatted)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }


package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var matrix [5][5]int32
var revealed [5][5]bool

func calran() int {
	return rand.Intn(5)
}

func calculate(val int) {
	rand.Seed(time.Now().UnixNano())
	vallocal := val
	for vallocal > 0 {
		num1 := calran()
		num2 := calran()
		if matrix[num1][num2] == 0 {
			matrix[num1][num2] = 1
			vallocal--
		}
	}
}

func printcompleteBoard() {
  fmt.Println("-------------------------------------------")
  for i:= 0;i <5;i++ {
    for j:=0;j<5;j++ {
      if matrix[i][j] == 1 {
        fmt.Print(" 💣 ")
      } else {
        fmt.Print(" ⭐ ")
      }
    }
    fmt.Print("\n")
    fmt.Print("\n")
  }
}

func printBoard() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if revealed[i][j] {
				if matrix[i][j] == 1 {
          fmt.Print(" 💣 ")
				} else {
          fmt.Print(" ⭐ ")
				}
			} else {
        fmt.Print(" ⭕ ")
			}
		}
		fmt.Println() 
		fmt.Println() 
	}
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	var num1 int

	for {
		fmt.Print("Enter the number of mines (1 - 24): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		num1, err = strconv.Atoi(input)
		if err == nil && num1 > 0 && num1 <= 24 {
			break
		}
		fmt.Println("Invalid input. Please enter a number between 1 and 24.")
	}

	calculate(num1)
	fmt.Println("Game started! Enter cell numbers (1-25) to reveal them. Find a bomb and the game ends.")

	for {
		printBoard()

		fmt.Print("Enter a cell number to reveal (1-25): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		cellNum, err := strconv.Atoi(input)

		if err != nil || cellNum < 1 || cellNum > 25 {
			fmt.Println("Invalid input. Please enter a number between 1 and 25.")
			continue
		}

		row := (cellNum - 1) / 5
		col := (cellNum - 1) % 5

		if revealed[row][col] {
			fmt.Println("Cell already revealed. Choose another cell.")
			continue
		}

		revealed[row][col] = true

		if matrix[row][col] == 1 {
			fmt.Println("You hit a bomb! Game over!")
			printBoard()
      fmt.Println()
      printcompleteBoard()
			break
		}
	}
}
