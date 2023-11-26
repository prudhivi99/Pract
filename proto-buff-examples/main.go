package main

import (
	"fmt"

	"local/proto-buff-examples/proto"
)

func printSimple() *proto.Person {
	return &proto.Person{
		Name:  "ABC1",
		Id:    1,
		Email: "ABC1@example.com",
	}
}

func printComplex() *proto.Complex {
	return &proto.Complex{
		OneDummy: &proto.Dummy{Id: 42, Name: "Name1"},
		MultipleDummies: []*proto.Dummy{
			{Id: 43, Name: "My name 2"},
			{Id: 44, Name: "My name 3"},
		},
	}
}

func printOneOf(message interface{}) {
	switch x := message.(type) {
	case *proto.Result_Id:
		    fmt.Printf("This is an Id: %d\n", message.(*proto.Result_Id).Id)
	case *proto.Result_Message:
		    fmt.Printf("This is a message: %s\n", message.(*proto.Result_Message).Message)
	default:
		fmt.Printf("message has unexpected type: %T\n", x);
	}
} 

func printMap() *proto.MapExample {
	message := &proto.MapExample {
		Ids: map[string]*proto.IdWrapper{
                "myId": {Id: 42},
				"myid2": {Id: 43},
				"myid3": {Id: 44},
		},
	}
	return message
}

//fun printEnum() *proto.Enumeration {
//	return &proto.Enumeration{
//		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
//	}
//}

func main() {

	// Print simple structure
	fmt.Println(printSimple())

	// Print complex structure
	fmt.Println(printComplex())

	// Print enumeration
	//fmt.Println(printEnum())

	//Print oneOf
	printOneOf(&proto.Result_Id{Id: 33})
	printOneOf(&proto.Result_Message{Message: "My name"})
	fmt.Println(printMap())
}

//go mod init local/proto-buff
//make generate
//make build
