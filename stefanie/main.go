package main

import (
	"fmt"
)


func Greeting() string {
	return "Hello, world!"
}

func main() {
	message := Greeting()
	fmt.Printf("%s\n", message)
}
