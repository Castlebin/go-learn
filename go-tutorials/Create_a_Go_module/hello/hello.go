package main

import (
	"fmt"

	"github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
