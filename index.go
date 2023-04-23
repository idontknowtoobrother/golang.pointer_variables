package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "first message"
	var msgPointer *string = &msg

	fmt.Println("msg:", msg)
	fmt.Println("pointer of msg:", *msgPointer)

	changeMessagePointer(&msg, "new message")
	fmt.Println("changed msg:", msg)

	changeMessagePointer(msgPointer, "new message 2")
	fmt.Println("changed msg:", msg)

	upperAllLetter(&msg)
	fmt.Println("upper msg by pointer of msg:", msg)

}

func changeMessagePointer(aPointer *string, newMessage string) {
	*aPointer = newMessage
}

func upperAllLetter(strPointer *string) {
	*strPointer = strings.ToUpper(*strPointer)
}
