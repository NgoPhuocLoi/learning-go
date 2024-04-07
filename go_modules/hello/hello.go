package main

import (
	"example/greetings"
	"fmt"

	// "example/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Loi", "Cris", "Mena"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	message, error := greetings.Hello("Loi")

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
