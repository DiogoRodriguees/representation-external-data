package app

import (
	"log"
	"net"
	"protobuffer/config"
	"protobuffer/logger"
	"protobuffer/messages"
	"protobuffer/services"
	"protobuffer/socket"
)

func Run() {
	log.Println(messages.StartingClient)

	conn, err := net.Dial(config.Socket.Type, config.Socket.Host)
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
		option := services.Prompt()
		if option == 4 {
			break
		}

		// Executeing option method
		request := services.ExecuteOptionSelected(option)

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
