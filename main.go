package main

import (
	"fmt"
	"runtime"
)

func main() {

	a := 25
	b := 35

	sum := a + b

	fmt.Println(sum)
	calculator()
}

func calculator() {
	fmt.Println("Enter Your First Number: ")

	// var then variable name then variable type
	var first int
	fmt.Scanln(&first)

	fmt.Println("Enter Your Second Number: ")

	// var then variable name then variable type
	var second int
	fmt.Scanln(&second)

	fmt.Println("Press 1 for add, 2 for multiply and 3 for substract")

	switch os := runtime.GOOS; os {
	case "1":
		fmt.Println(first + second)
	case "2":
		fmt.Println(first * second)
	case "3":
		fmt.Println(first - second)
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

}
