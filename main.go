package main

import (
	"fmt"
	"time"
)

func main() {

	a := 25
	b := 35

	sum := a + b

	fmt.Println(sum)
	swapNumbers()
	weekdays()

	Even_Odd()
}

// weekdaysfunction
func weekdays() {

	// Using time.Now().Weekday() function.
	dt := time.Now().Weekday()
	fmt.Println(dt.String())
}
func swapNumbers() {

	var number1, number2, number3 int
	number1 = 16
	number2 = 20
	fmt.Println("Numbers before swapping: \n Number 1 =", number1, "\n Number 2 =", number2)

	// swapping the numbers
	number3 = number1
	number1 = number2
	number2 = number3

	// printing the numbers after swapping
	fmt.Println("Numbers after swapping:\n Number 1 =", number1, "\n Number 2 =", number2, "\n(Swap within the function)")

}

func Even_Odd() {
	var num int = 0
	fmt.Println("Enter number: ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println("Number is EVEN")
	} else {
		fmt.Println("Number is ODD")
	}
}
