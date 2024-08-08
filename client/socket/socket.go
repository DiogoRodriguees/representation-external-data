package socket

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"protobuffer/interfaces"

	"google.golang.org/protobuf/proto"
)

func WriteMessage(conn net.Conn, request *interfaces.Request) error {
	// Marshalling request
	requestBytes, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Erro ao serializar Request: %v", err)
		return err
	}

	// Send message size
	sizeBytes := make([]byte, 4)
	size := uint32(len(requestBytes))
	for i := 0; i < 4; i++ {
		sizeBytes[i] = byte(size >> (24 - i*8))
	}
	conn.Write(sizeBytes)

	// Send message
	_, err = conn.Write(requestBytes)
	if err != nil {
		log.Fatalf("Erro ao enviar Request: %v", err)
		return err
	}

	return nil
}

func ReadResponse(conn net.Conn) error {
	// Ler o tamanho da resposta
	sizeData := make([]byte, 4)
	_, err := conn.Read(sizeData)
	if err != nil {
		fmt.Println("Erro ao ler tamanho da resposta:", err)
		return err
	}
	s := binary.BigEndian.Uint32(sizeData)

	// Ler a resposta
	responseData := make([]byte, s)
	_, err = conn.Read(responseData)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return err
	}

	// Deserializar a resposta
	response := &interfaces.Response{}
	err = proto.Unmarshal(responseData, response)
	if err != nil {
		fmt.Println("Erro ao deserializar resposta:", err)
		return err
	}

	fmt.Println("\n-----------------------------------------------")
	fmt.Println("*                  Response                   *")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Status:", response.Status)
	fmt.Println("Mensagem:", response.Message)
	fmt.Println("Filmes: ")
	for _, movie := range response.Movies {
		fmt.Println("* ", movie.Title)
	}
	fmt.Println("-----------------------------------------------")

	return nil
}
