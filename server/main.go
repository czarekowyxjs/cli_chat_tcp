package main

import (
	"cli_chat_tcp/server/pkg/server"
)

func main() {
	s := server.New("localhost", 6742)
	s.Run()
}