package main

import "fmt"

func main() {
	normal()
	err()
	err1()
}

func normal() {
	defer catch()
	fmt.Println("Normal")
}

func err() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	fmt.Println("Error")
	panic("HIAAAA")
	fmt.Println("After")
}

func err1() {
	defer catch()
	fmt.Println("Error")
	panic("HIAAAA")
	fmt.Println("After")
}

func catch() {
	msg := recover()
	fmt.Println(msg)
}
