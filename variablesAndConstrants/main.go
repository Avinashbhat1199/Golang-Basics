package main

import (
	"fmt"
)

func main() {
	fmt.Println("Variables and constrants ")

	//variable with var
	var username string = "avi"
	fmt.Println("username", username)
	fmt.Printf("type of username %T \n", username)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("type of username %T \n", isBool)

	var isint uint16 = 345
	fmt.Println(isint)
	fmt.Printf("type of username %T \n", isint)

	//variable  without var
	sum := 239
	fmt.Println("sum", sum)

}
