package main

import "fmt"

func main() {

	vikas := User{"Vikas", true}
	fmt.Printf("The user is %+v \n", vikas)
	makeoffline(&vikas)
	fmt.Printf("The user is %+v \n", vikas)
}

func makeoffline(user *User) {
	user.isonline = false
}

type User struct {
	name     string
	isonline bool
}
