package hydrachat

import (
	"fmt"
	"hydra/hydra/hlogger"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	r := CreateRoom("HydraChat")

	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error accepting connection from char client ", err)
		}

		go handleConnection(r, conn)
	}

	return err
}

func handleConnection(r *room, c net.Conn) {
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
