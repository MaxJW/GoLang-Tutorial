package main

import "fmt"

//Hello - Print out something
func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
