package server

import (
	"log"
	"net"
	"strconv"
	"os"
	"bufio"
	"github.com/lithammer/shortuuid"
	"regexp"
)

type User struct {
	uid string
	username string
	connection net.Conn
}

type Server struct {
	host string
	port int
	users []User
	commands map[string]func(string, net.Conn)
}

func New(host string, port int) Server {
	return Server{
		host: host,
		port: port,
		users: make([]User, 0),
		commands: make(map[string]func(string, net.Conn), 0),
	}
}

func (server Server) Run() {
	listener, err := net.Listen("tcp", server.host + ":" + strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Unable to listen %v", err)
		os.Exit(1)
	}
	log.Println("Server is listening...")

	server.loadCommands()
	server.loop(listener)
}

func (server Server) loop(listener net.Listener) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Listener error %v", err)
			continue
		}

		go server.handleConnection(conn)
	}
}

func (server Server) handleConnection(conn net.Conn) {
	go server.createUser(conn)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
	
		if len(message) == 0 {
			server.removeUser(conn)
			break
		}

		if err != nil {
			log.Fatalf("Unable to handle message from %v, error is %v", conn, err)
		}

		server.parseMessage(message, conn)
	}
}

func (server Server) createUser(conn net.Conn) {
	log.Printf("New connection %v", conn)

	server.users = append(server.users, User{
		uid: shortuuid.New(),
		username: "",
		connection: conn,
	})
}

func (server Server) removeUser(conn net.Conn) {
	log.Printf("%v left server", conn)

	var filteredUsers = make([]User, 0)

	for i := 0; i < len(server.users); i++ {
		if conn != server.users[i].connection {
			filteredUsers = append(filteredUsers, server.users[i])
		}
	}

	server.users = filteredUsers
}

func (server Server) parseMessage(message string, conn net.Conn) {
	log.Printf("%v: %v", conn, message)

	if server.isCommand(message) {
		server.emitCommand(message, conn)
	} else {
		server.emitMessage(message, conn)
	}
}	

func (server Server) isCommand(message string) bool {
	match, _ := regexp.MatchString("^/", message)

	return match
}	

func (server Server) emitCommand(command string, conn net.Conn) {
	parsedCommand := server.parseCommand(command)

	if parsedCommand != nil {

		server.selectCommand(parsedCommand[1], parsedCommand[2], conn)

	} else {
		//handle error - empty command
	}
}

func (server Server) emitMessage(message string, conn net.Conn) {

}

func (server Server) join(username string, conn net.Conn) {
	userExists := true

	for i := 0; i < len(server.users); i++ {
		if server.users[i].connection == conn {
			server.users[i].username = username
			userExists = false
			break
		}
	}

	if(userExists) {
		//emitMessageToUser - username already exists
	} else {
		//emitMessageToUser - succesfully joined
	}
}

func (server Server) loadCommands() {
	server.commands["join"] = server.join
}

func (server Server) parseCommand(command string) []string {
	re := regexp.MustCompile(`^/([a-zA-Z]+)\s+([[:print:]]+)`)
	matches := re.FindStringSubmatch(command)

	return matches
}

func (server Server) selectCommand(name string, content string, conn net.Conn) {
	if commandFunc, found := server.commands[name]; found {
		commandFunc(content, conn)
	} else {
		//handle error - command not found
	}
}