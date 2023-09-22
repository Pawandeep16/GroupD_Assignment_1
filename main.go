package main

import (
	"fmt"
)

func main() {

	a := 25
	b := 35

	sum := a + b

	fmt.Println(sum)
	swapNumbers()
        weekdays()
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
