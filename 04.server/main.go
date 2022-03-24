package main

import (
	"github.com/go-mysql-org/go-mysql/server"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:4000")

	c, _ := l.Accept()

	// Create a connection with user root and an empty password.
	// You can use your own handler to handle command here.
	conn, _ := server.NewConn(c, "root", "", server.EmptyHandler{})

	for {
		conn.HandleCommand()
	}
}
