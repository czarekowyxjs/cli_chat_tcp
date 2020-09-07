package command

import (
	"log"
	"net"
	"regexp"
	"errors"
)

type Command struct {
	command string
	conn net.Conn
	serverCommands map[string]func(string, net.Conn)
}

type ParsedCommand struct {
	name string
	content string
}

func New(command string, conn net.Conn, serverCommands map[string]func(string, net.Conn)) Command {
	return Command{
		command: command,
		conn: conn,
		serverCommands: serverCommands,
	}
}

func (c *Command) Exec() {
	if c.isExecutable() {
		c.executeCommand()
	}
}

func (c *Command) isExecutable() bool {
	match, _ := regexp.MatchString("^/", c.command)

	return match
}

func (c *Command) executeCommand() {
	parsedCommand, err := c.parseCommand()

	if err != nil {
		log.Print(err)
	}

	if commandFunc, found := c.serverCommands[parsedCommand.name]; found {
		commandFunc(parsedCommand.content, c.conn)
	}
}

func (c *Command) parseCommand() (ParsedCommand, error) {
	re := regexp.MustCompile(`^/([a-zA-Z]+)\s+([[:print:]]+)`)
	matches := re.FindStringSubmatch(c.command)

	if matches == nil {
		return ParsedCommand{"", ""}, errors.New("Unable to parse command, probably empty command content")
	}

	return ParsedCommand{
		name: matches[1],
		content: matches[2],
	}, nil	
}