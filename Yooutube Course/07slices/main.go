package main

import "fmt"

func main() {
	slices()
}

func slices() {

	var one = []int{1, 2, 3}
	fmt.Println(one)
	one = append(one, 1, 2, 3)
	fmt.Println(one)
	two := append(one[3:])
	fmt.Println(two)
	whatismake()

}

func whatismake() {
	newarr := make([]int, 4)
	fmt.Println(newarr)
	newarr[2] = 1
	fmt.Println(newarr)
	deletesecondindex(newarr)

}

//delete second index

func deletesecondindex(array []int) {
	index := 2
	newslice := append(array[:index], array[index+1:]...)
	fmt.Println("New array with second element deleted is :")
	fmt.Println(newslice)
}
