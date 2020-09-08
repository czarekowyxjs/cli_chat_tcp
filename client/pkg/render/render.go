package render

import (
	"github.com/inancgumus/screen"
	"fmt"
)

type Render struct {
	Messages []Message
	Announcement bool
}

type Message struct {
	sender string
	content string
}

func New() Render {
	return Render{
		Messages: make([]Message, 0),
		Announcement: false,
	}
}

func (r *Render) Run() {
	screen.Clear()
	r.renderAnnouncement()
	r.renderMessages()
	r.renderInput()
}

func (r *Render) AddMessage(sender string, content string) {
	r.Messages = append(r.Messages, Message{sender, content})
}

func (r * Render) renderMessages() {
	for i := 0; i < len(r.Messages); i++ {
		fmt.Printf("%s: %s", r.Messages[i].sender, r.Messages[i].content)
	}
}

func (r *Render) renderAnnouncement() {
	if !r.Announcement {
		fmt.Printf("Hello, you are not logged in. \nIf you want to join to chat, type \"/join your_username\" \n(your username is temporary, only for your chating time)\n\n")
	}
}

func (r *Render) renderInput() {
	fmt.Printf(">> ")
}