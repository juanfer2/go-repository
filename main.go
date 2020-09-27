package main

import (
	"fmt"
	
	"github.com/juanfer2/go-repository/server"
)

func main() {
	fmt.Println("hi world")
	server := serverApp(":4000")
	server.Listen()
}
