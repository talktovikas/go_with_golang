package main

import "fmt"

func main() {
	fmt.Println("Variables")
	var fname string = "Vikas"
	lname := "'Yadav'"
	var isonline bool
	fmt.Printf("FirstName is %s and Type is %T \n", fname, fname)
	fmt.Printf("LastName is %s and Type is %T \n", lname, lname)
	fmt.Printf("isonline is %s and Type is %T \n", isonline, isonline)
}
