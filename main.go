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
	averageofthreenumbers()
	Even_Odd()
	switchFunction()
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

func averageofthreenumbers() {
    //Declare 4 integer type variables
    var num1 int =10
    var num2 int =20
    var num3 int =30
    var avg int=0 

    avg=(num1+num2+num3)/3
    
    fmt.Println("Average is: ",avg)
}


func switchFunction() {
    var val int=3
    
     switch val{ 
       case 4-1: 
            fmt.Println("INDIA") 
       case 4-2: 
            fmt.Println("USA") 
       case 4-3: 
            fmt.Println("UK") 
       case 4-4: 
            fmt.Println("CHINA")
       default:  
            fmt.Println("Invalid value") 
   }  
}
