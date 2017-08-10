package network

import "net"
import "os"

const SERVER_INFO = ":19850"

func Listen() (listener net.Listener) {
	listener, err := net.Listen("tcp", SERVER_INFO)

	if err != nil {
		os.Exit(1)
	}

	return
}
