package app

import (
	"log"
	"net"
	"protobuffer/logger"
	"protobuffer/logger/messages"
	"protobuffer/movie"
	"protobuffer/socket"
)

func Run() {
	log.Println(messages.StartingClient)

	conn, err := net.Dial(socket.Config.Type, socket.Config.Host)
	if err != nil {
		logger.Error(messages.FailedOnStartSocket, err)
	}
	defer conn.Close()

	start(conn)

	log.Println(messages.StoppingClient)
}

func start(conn net.Conn) {
	for {
		// Read option selected by user
		option := movie.ShowMenu()
		if option == 4 {
			break
		}

		movieManager := movie.Init()
		createRequest := movieManager.Option[option]
		request := createRequest()

		// Writing message
		err := socket.WriteMessage(conn, request)
		if err != nil {
			logger.Error(messages.FailedOnWriteMessage, err)
			return
		}

		// Reading response
		err = socket.ReadResponse(conn)
		if err != nil {
			logger.Error(messages.FailedOnReadMessage, err)
			return
		}
	}
}
