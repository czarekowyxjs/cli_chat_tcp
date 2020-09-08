package main

import (
	"cli_chat_tcp/client/pkg/chat"
)

func main() {
	c := chat.New("localhost", 6742)
	c.Run()
}