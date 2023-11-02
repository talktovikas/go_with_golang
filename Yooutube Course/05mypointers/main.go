package main

import "fmt"

func main() {
	fmt.Println("My Pointer Module")
	mypointer()
}

func mypointer() {

	var myptr *int
	fmt.Println(myptr)
	// fmt.Println(*myptr)  The Famous Segmentation fault
	value := 37
	myptr = &value
	fmt.Println(myptr)
	fmt.Printf("Type of it %T", myptr)
	fmt.Println(*myptr)
	fmt.Printf("Type of this %T", *myptr)

}
