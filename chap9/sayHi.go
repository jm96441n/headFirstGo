package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from:", m)
}

func main() {
	greeter := MyType("a value here")
	greeter.sayHi()
	anotherOne := MyType("another one")
	anotherOne.sayHi()
}
