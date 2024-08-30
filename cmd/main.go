package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Endereço do Redis
	})

	// Adiciona mensagens ao stream
	for i := 0; i < 10; i++ {
		err := rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: "mystream",
			Values: map[string]interface{}{"message": "Hello, World!", "id": i},
		}).Err()

		if err != nil {
			log.Fatalf("Erro ao adicionar mensagem ao stream: %v", err)
		}
	}
}
