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

func printmatrix() {

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
 //  looptimes := num1 
 //  for looptimes > 0 {
	// fmt.Print("Enter the number of mines (1 - 24): ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	//
	// 	var err error
	// 	num1, err = strconv.Atoi(input)
	// 	if err == nil && num1 > 0 && num1 <= 24 {
	// 		break
	// 	}
	// 	fmt.Println("Invalid input. Please enter a number between 1 and 24.")
	//
 //  }
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			value := 5*i + j + 1 
			if matrix[i][j] == 1 {
				formatted := fmt.Sprintf(" X(%2d) ", value)
				fmt.Print(formatted)
			} else {
				formatted := fmt.Sprintf(" .(%2d) ", value)
				fmt.Print(formatted)
			}
		}
		fmt.Println()
		fmt.Println()
	}
}
