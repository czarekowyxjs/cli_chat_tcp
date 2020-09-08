package server

import (
	"log"
	"net"
	"strconv"
	"os"
	"bufio"
	"cli_chat_tcp/server/pkg/command"
	"cli_chat_tcp/server/pkg/user"
)

type Server struct {
	host string
	port int
	users []user.User
	commands map[string]func(string, net.Conn)
}

func New(host string, port int) Server {
	return Server{
		host: host,
		port: port,
		users: make([]user.User, 0),
		commands: make(map[string]func(string, net.Conn), 0),
	}
}

func (server *Server) Run() {
	listener, err := net.Listen("tcp", server.host + ":" + strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Unable to listen %v", err)
		os.Exit(1)
	}

	log.Println("Server is listening...")

	server.initCommands()
	server.loop(listener)
}

func (server *Server) loop(listener net.Listener) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Listener error %v", err)
			continue
		}

		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	server.newConnection(conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
	
		if len(message) == 0 {
			server.removeConnection(conn)
			break
		}

		if err != nil {
			log.Fatalf("Unable to handle message from %v, error is %v", conn, err)
			break
		}

		c := command.New(message, conn, server.commands)
		c.Exec()
	}
}

func (server *Server) newConnection(conn net.Conn) {
	server.users = append(server.users, user.New(conn))
}

func (server *Server) removeConnection(conn net.Conn) {
	server.users = user.RemoveUserByConnection(server.users, conn)
}

func (server *Server) sendMessage(content string, conn net.Conn) {
	conn.Write([]byte(content+"\n"))
}

func (server *Server) initCommands() {
	server.commands["join"] = server.join
	server.commands["message"] = server.message
}

func (server *Server) join(username string, conn net.Conn) {
	users, err := user.SetUsername(server.users, username, conn)

	if err != nil {
		log.Print(err)
		server.sendMessage("101 SERVER Username already exists", conn)
	} else {
		server.sendMessage("100 SERVER Successfully joined", conn)
	}

	server.users = users
}

func (server *Server) message(content string, conn net.Conn) {

}