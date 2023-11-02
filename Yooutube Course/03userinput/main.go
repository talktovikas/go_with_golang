package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("UserInputModule")

	//Taking input

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter something here.")
	input, _ := reader.ReadString('\n')
	fmt.Println("User Input is", input)
	fmt.Printf("User Input Type is : %T", input)

}
