package response

import (
	"bufio"
	"net"
	"fmt"
	"os"
	"regexp"
)

type Response struct {
	conn net.Conn
	parsedResponse ParsedResponse
}

type ParsedResponse struct {
	code int
	sender string
	content string
}

func New(conn net.Conn) Response {
	return Response{
		conn: conn,
		parsedResponse: ParsedResponse{0, "", ""},
	}
}

func (r *Response) Get() {
	message := r.readMessage()
	r.parseMessage(message)
}

func (r *Response) readMessage() string {
	message, err := bufio.NewReader(r.conn).ReadString('\n')

	if len(message) == 0 {
		fmt.Errorf("Server error - disconnected")
		os.Exit(1)
	}

	if err != nil {
		fmt.Errorf("Error with handling message from server %v", err)
		os.Exit(1)
	}

	return message
}

func (r *Response) parseMessage(message string) {
	fmt.Print(message)
	re := regexp.MustCompile(`^([0-9]{3})\s([[:print:]]+)`)
	matches := re.FindStringSubmatch(message)

	fmt.Print(matches)

	//to finish
}
