package main

import "fmt"

/**
* Public method
 */
func Say() string {
	return "hello world"
}

func main() {
	fmt.Println("hello world")
	fmt.Println(Say())
}
