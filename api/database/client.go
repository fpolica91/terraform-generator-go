package database

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "gGyz5vOxBMaTlAmbRE2GDZiCuSk=",
})

func Set(key string, value interface{}) error {
	ctx := context.Background()
	jsonValue, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}
	return RedisClient.Set(ctx, key, jsonValue, 0).Err()

}

// func Set(key string, value interface{}) error {
// 	ctx := context.Background()

// 	existing, e := RedisClient.Get(ctx, key).Result()
// 	if e != nil {
// 		fmt.Println("no existing data found")
// 		existing = "[]"
// 	}
// 	var configs []interface{}
// 	err := json.Unmarshal([]byte(existing), &configs)
// 	if err != nil {
// 		panic(err)
// 	}
// 	configs = append(configs, value)
// 	jsonValue, err := json.Marshal(configs)

// 	if err != nil {
// 		panic(err)
// 	}
// 	return RedisClient.Set(ctx, key, jsonValue, 0).Err()

// }
