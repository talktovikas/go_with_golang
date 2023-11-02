package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Inside Conversion module")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Input a Number to increase it with one :")
	input, _ := reader.ReadString('\n')
	intinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Increased Output is", intinput+1)
	}
	// timenow := time.Now()
	// fmt.Printf("%T", timenow)
}
