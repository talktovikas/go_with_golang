package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func generate_IPaddress() string {
	PartA := strconv.Itoa(rand.Intn(300))
	PartB := strconv.Itoa(rand.Intn(300))
	PartC := strconv.Itoa(rand.Intn(300))
	PartD := strconv.Itoa(rand.Intn(300))
	Result := PartA + "." + PartB + "." + PartC + "." + PartD
	return Result
}
func broken_ip_bits(IP string) (int, int, int, int) {
	SplitArray := strings.Split(IP, ".")
	A, _ := strconv.Atoi(SplitArray[0])
	B, _ := strconv.Atoi(SplitArray[1])
	C, _ := strconv.Atoi(SplitArray[2])
	D, _ := strconv.Atoi(SplitArray[3])
	return A, B, C, D
}

func decide_auth(A int, B int, C int, D int) bool {
	if A < 256 && B < 256 && C < 256 && D < 256 {
		return true
	} else {
		return false
	}
}

func generate_response(Flag bool, IP string) {
	if Flag == true {
		fmt.Println("Yes, %s  is a valid IP Address ", IP)
	} else {
		fmt.Println("No, %s is Not a Valid IP Address", IP)
	}
}
func ip_validation() {
	NewIP := generate_IPaddress()
	A, B, C, D := broken_ip_bits(NewIP)
	Flag := decide_auth(A, B, C, D)
	generate_response(Flag, NewIP)
	time.Sleep(1 * time.Second)
	ip_validation()
}

func main() {
	ip_validation()
}
