package main

import (
	"fmt"

	"local/proto-buff/proto"
)

func main() {
	// Create a new Person
	person := &proto.Person{
		Name:  "John Doe",
		Id:    1,
		Email: "john.doe@example.com",
	}

	// Do something with the person...
	fmt.Printf("Name: %s, ID: %d, Email: %s\n", person.Name, person.Id, person.Email)
}

//go mod init local/proto-buff
//make generate
//make build
