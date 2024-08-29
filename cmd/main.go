package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/sandronister/generate-devices/internal/service"
)

var wg = sync.WaitGroup{}
var j = 0
var server = 1

func send() {

	for i := 0; i < 333; i++ {
		conn, err := net.Dial("tcp", "localhost:3005")
		if err != nil {
			fmt.Println("Erro ao conectar:", err)
			return
		}

		defer conn.Close()
		data := make([]byte, 1024)
		for i := 0; i < 1024; i++ {
			data[i] = byte(i % 256) // Preencher com valores de 0 a 255
		}
		payload := service.NewPayload(0x1234, 0x01, data)
		binaryData, err := payload.ToBinary()
		if err != nil {
			fmt.Println("Erro ao converter para binário:", err)
			return
		}

		// Enviar o payload binário através da conexão
		_, err = conn.Write(binaryData)
		if err != nil {
			j++
			//fmt.Println("Erro ao enviar dados:", err)
			return
		}

		//fmt.Println("Payload enviado com sucesso")

	}
	fmt.Println("Erros:", j)
	j = 0
	server++
	fmt.Println("Server:", server)
	wg.Done()
}

func main() {
	wg.Add(1)
	for range 1 {
		go send()
	}
	wg.Wait()
}
