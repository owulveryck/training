package training

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
)

const (
	// Local stdio
	Local = iota
	// Network listener on localhost:1234
	Network = iota
)

// ListenAddr for the network connexion
var ListenAddr = "localhost:1234"

// GetStdIO takes as a parameter Local or Network and return 3 channels:
// stdin, stdout, stderr
// it returns an error if the stdio parameter is Network and it cannot listen to
// ListenAddr
func GetStdIO(stdio int) (chan io.Reader, chan io.Writer, chan io.Writer, error) {
	stdin := make(chan io.Reader)
	stdout := make(chan io.Writer)
	stderr := make(chan io.Writer)
	switch stdio {

	case Network:
		l, err := net.Listen("tcp", ListenAddr)
		if err != nil {
			return stdin, stdout, stderr, err
		}
		go func() {
			defer l.Close()
			for {
				// Wait for a connection.
				conn, err := l.Accept()
				if err != nil {
					log.Println(err)
					continue
				}
				// Handle the connection in a new goroutine.
				// The loop then returns to accepting, so that
				// multiple connections may be served concurrently.
				stdin <- conn
				stdout <- conn
				stderr <- conn
			}
		}()
	case Local:
		go func() {
			stdin <- os.Stdin
			stdout <- os.Stdout
			stderr <- os.Stderr
		}()
	default:
		return stdin, stdout, stderr, errors.New("Unknown handler")
	}
	return stdin, stdout, stderr, nil
}
