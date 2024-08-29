package main

import (
	"fmt"
	"net"
	"sync"
)

var wg = sync.WaitGroup{}

func send() {
	listPaises := []string{"Brasil", "Argentina", "Chile", "Uruguai", "Paraguai", "Bolívia", "Peru", "Colômbia", "Venezuela", "Equador", "Guiana", "Suriname", "Guiana Francesa"}

	for _, item := range listPaises {
		conn, err := net.Dial("tcp", "localhost:3005")
		if err != nil {
			return
		}

		defer conn.Close()

		_, err = conn.Write([]byte(item))

		if err != nil {
			fmt.Println(err)

		}

	}
	wg.Done()
}

func main() {
	wg.Add(1)
	for range 1 {
		go send()
	}
	wg.Wait()
}
