package main

import "fmt"

func f1() {
	f2()
	fmt.Println("F1")
}

func f2() {
	f3()
	fmt.Println("F2")
}

func f3() {
	fmt.Println("F3")
}

func main() {
	f1()
}