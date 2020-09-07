package user

import (
	"net"
	"log"
	"fmt"
	"github.com/lithammer/shortuuid"
)

type User struct {
	Uid string
	Username string
	Conn net.Conn
}

func New(conn net.Conn) User {
	log.Printf("New connection %v", conn)

	return User{
		Uid: shortuuid.New(),
		Username: "",
		Conn: conn,
	}
}

func RemoveUserByConnection(users []User, conn net.Conn) []User {
	log.Printf("%v left server", conn)

	var usersAfterRemove = make([]User, 0)

	for i := 0; i < len(users); i++ {
		if users[i].Conn != conn {
			usersAfterRemove = append(usersAfterRemove, users[i])
		}
	}

	return usersAfterRemove
}

func SetUsername(users []User, username string, conn net.Conn) ([]User, error) {

	if !isUsernameTaken(users, username) {
		for i := 0; i < len(users); i++ {
			if users[i].Conn == conn {
				users[i].Username = username
				break
			}
		}

		return users, nil
	} else {
		return users, fmt.Errorf("Username <%s> already taken", username)
	}
}

func isUsernameTaken(users []User, username string) bool {

	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return true
		}
	}

	return false
}
