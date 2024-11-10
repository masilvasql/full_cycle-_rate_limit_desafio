package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(host, port, password string, db int) (*redis.Client, error) {
	fmt.Printf("host: %s, port: %s, password: %s, db: %d\n", host, port, password, db)
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	if err := checkConnection(client); err != nil {
		return nil, err
	}

	return client, nil
}

func checkConnection(client *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	return err
}
