package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi world")
	server := serverApp(":4000")
	server.Listen
}
