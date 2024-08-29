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
	ctx    = context.Background()
	stream = "paisesStream"
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func send() {
	defer wg.Done()

	listPaises := []string{"Brasil", "Argentina", "Chile", "Uruguai", "Paraguai", "Bolívia", "Peru", "Colômbia", "Venezuela", "Equador", "Guiana", "Suriname", "Guiana Francesa"}

	for _, item := range listPaises {
		err := rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{"pais": item},
		}).Err()

		if err != nil {
			fmt.Println("Erro ao enviar para o stream:", err)
			return
		}
	}
}

func main() {
	wg.Add(1)
	go send()
	wg.Wait()
}
