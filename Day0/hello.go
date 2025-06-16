package main

import "fmt"

func say_hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(say_hello("kishore"))
}
