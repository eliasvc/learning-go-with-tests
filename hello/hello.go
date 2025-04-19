package main

import "fmt"

func HelloWorld(name string) string {
	if name == "" {
		return "Hello, world!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
func main() {
	fmt.Println(HelloWorld("Chris"))
}
