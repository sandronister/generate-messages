package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	wg     = sync.WaitGroup{}
	rdb    *redis.Client
	stream = "ruptela.com"
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func send() {
	defer wg.Done()

	listPaises := []string{
		"Brasil",
		"Argentina",
		"Chile",
		"Uruguai",
		"Paraguai",
		"Bolívia",
		"Peru",
		"Colômbia",
		"Venezuela",
		"Equador",
		"Guiana",
		"Suriname",
		"Guiana Francesa",
		"Estados Unidos",
		"Canadá",
		"México",
		"Cuba",
		"República Dominicana",
		"Costa Rica",
		"Panamá",
		"China",
		"Japão",
		"Coreia do Sul",
		"Índia",
		"Rússia",
		"Alemanha",
		"França",
		"Itália",
		"Espanha",
		"Portugal",
	}

	type message struct {
		Index int
		Item  string
	}

	for i, item := range listPaises {
		msm := message{
			Index: i,
			Item:  item,
		}

		message := fmt.Sprintf("%v", msm)

		err := rdb.LPush(context.Background(), stream, message).Err()
		if err != nil {
			fmt.Println("Erro ao adicionar item à fila:", err)
			return
		}

		fmt.Println("Item adicionado à fila:", item)

	}
}

func main() {
	wg.Add(1)
	go send()
	wg.Wait()
}
