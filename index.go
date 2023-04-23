package main

import "fmt"

func main() {
	msg := "first message"
	var msgPointer *string = &msg

	fmt.Println("msg:", msg)
	fmt.Println("pointer of msg:", *msgPointer)

	changeMessagePointer(&msg, "new message")
	fmt.Println("changed msg:", msg)
}

func changeMessagePointer(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
