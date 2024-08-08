package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"protobuffer/interfaces"
	"protobuffer/services"

	"google.golang.org/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Substitua "localhost:8080" pela endereço IP e porta do seu servidor
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor: %v", err)
	}
	defer conn.Close()
	for {
		option := services.Prompt()
		if option == 4 {
			break
		}
		request := services.ExecuteOptionSelected(option)
		requestBytes, err := proto.Marshal(request)

		sizeBytes := make([]byte, 4) // Garante que temos exatamente 4 bytes
		size := uint32(len(requestBytes))
		for i := 0; i < 4; i++ {
			sizeBytes[i] = byte(size >> (24 - i*8))
		}

		// Passo 3: Enviar a mensagem de tamanho seguida pelos requestBytes
		conn.Write(sizeBytes)
		if err != nil {
			log.Fatalf("Erro ao serializar Request: %v", err)
		}

		_, err = conn.Write(requestBytes)
		if err != nil {
			log.Fatalf("Erro ao enviar Request: %v", err)
		}

		// Ler o tamanho da resposta
		sizeData := make([]byte, 4)
		_, err = conn.Read(sizeData)
		if err != nil {
			fmt.Println("Erro ao ler tamanho da resposta:", err)
			return
		}
		s := binary.BigEndian.Uint32(sizeData)

		// Ler a resposta
		responseData := make([]byte, s)
		_, err = conn.Read(responseData)
		if err != nil {
			fmt.Println("Erro ao ler resposta:", err)
			return
		}

		// Deserializar a resposta
		response := &interfaces.Response{}
		err = proto.Unmarshal(responseData, response)
		if err != nil {
			fmt.Println("Erro ao deserializar resposta:", err)
			return
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
	}
}
