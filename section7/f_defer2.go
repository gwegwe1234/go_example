package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Print("Hi! ")
	}()
}

func main() {
	sayHello("Golang!")
}
