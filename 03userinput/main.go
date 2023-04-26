package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var avi string
	fmt.Println("hello user input ")
	reader := bufio.NewReader(os.Stdin)
	// comma eeror /comma ok
	//avi,error
	avi, _ = reader.ReadString('\n')
	fmt.Println(avi)

}
