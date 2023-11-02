package main

import "fmt"

func main() {
	programming := make(map[string]string)
	fmt.Println(programming)
	programming["JS"] = "Javascript"
	programming["C"] = "C LANG"
	programming["GO"] = "Go Lang"
	fmt.Println(programming)
	fmt.Println("Go stands for", programming["GO"])
	//error
	fmt.Println("Go stands for", programming["q"])
	// fmt.Println("Go stands for",programming["GO"])
	// fmt.Println("Go stands for",programming["GO"])

}
