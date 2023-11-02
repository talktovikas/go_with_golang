package main

import "fmt"

func main() {
	arrayfun()

}

func arrayfun() {
	var arraylist [4]string
	arraylist[0] = "0"
	arraylist[1] = "1"
	arraylist[2] = "2"
	arraylist[3] = "3"

	fmt.Println("List is", arraylist)
	fmt.Println("List is", len(arraylist)) // This will always show 4 despite how many we will fill here.

	var newarray = []string{"1", "2", "3"}
	fmt.Println(newarray)
}
