package chat

import (
	"net"
	"os"
	"strconv"
	"fmt"
	"bufio"
	"cli_chat_tcp/client/pkg/render"
	"cli_chat_tcp/client/pkg/response"
)

type Chat struct {
	host string
	port int
	conn net.Conn
	render render.Render
	joined bool
}

func New(host string, port int) Chat {
	return Chat{
		host: host,
		port: port,
		conn: nil,
		render: render.New(),
		joined: false,
	}
}

func (c *Chat) Run() {
	conn, err := net.Dial("tcp", c.host + ":" + strconv.Itoa(c.port))

	if err != nil {
		fmt.Errorf("Unable to connect to server %v", err)
		os.Exit(1)
	}

	c.conn = conn
	c.render.Run()

	go c.handleServerMessage()
	c.handleUserInput()
}

func (c *Chat) handleUserInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Errorf("Unable to handle message %v", err)
		}

		if _, err = c.conn.Write([]byte(message)); err != nil {
			fmt.Errorf("Unable to send a message %v", err)
		}

		c.render.AddMessage("You", message)
		c.render.Run()
	}
}

func (c *Chat) handleServerMessage() {
	for {
		//message, err := bufio.NewReader(c.conn).ReadString('\n')

		res := response.New(c.conn)
		res.Get()

		// if len(message) == 0 {
		// 	fmt.Errorf("Server error - disconnected")
		// 	os.Exit(1)
		// }

		// if err != nil {
		// 	fmt.Errorf("Error with handling message from server %v", err)
		// 	os.Exit(1)
		// }

		// c.render.AddMessage("[SERVER]", message)
		// c.render.Run()
	}
}