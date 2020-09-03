package main

import (
	"log"
	"net"
	//"github.com/inancgumus/screen"
	"os"
	"bufio"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6742")

	if err != nil {
		log.Fatalf("Unable to connect to server %v", err)
		os.Exit(1)
	}

	fmt.Println("Hello, you are not logged in. \nIf you want to join to chat, type \"/join your_username\" \n(your username is temporary, only for your chating time)")

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")

		message, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Unable to handle message %v", err)
		}

		if _, err = conn.Write([]byte(message)); err != nil {
			log.Fatalf("Unable to send a message %v", err)
		}
	}

}