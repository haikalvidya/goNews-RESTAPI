// package redis

// import (
// 	"encoding/json"
// 	"context"
// 	"time"
// 	"fmt"

// 	"github.com/go-redis/redis"
// )

// // declaration for the storage service and redis
// var (
// 	storageService = &StorageService{}
// 	ctx = context.Background()
// )

// // define struct of wrapper for raw redis client
// type StorageService struct {
// 	redisClient *redis.Client
// }

// // set timer for cache redis duration
// const CacheDuration = 6 * time.Hour

// // init the storage service and return a store pointer
// func InitializeStorage() *StorageService {
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr:		"localhost:6379",
// 		Password:	"", // no password
// 		DB:			0,
// 	})

// 	pong, err := redisClient.Ping(ctx).Result()
// 	if err != nil {
// 		panic(fmt.Sprintf("Error init Redis: %v", err))
// 	}

// 	fmt.Printf("Redis started successfully: ping message = {%s}", pong)
// 	storageService.redisClient = redisClient
// 	return storageService
// }

// // set a key and value
// func Set(key string, data interface{}) error {
// 	value, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	err := storageService.redisClient.Set(ctx, key, value, CacheDuration).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Check Existance of a key
// func IsExists(key string) bool {
// 	result, err := storageService.redisClient.Exists(ctx, key).Result()
// 	if err != nil {
// 		return false
// 	}
// 	return result > 0
// }

// // get a key
// func Get(key string) interface{} {
// 	result, err := storageService.redisClient.Get(ctx, key).Result()
// 	if err != nil {
// 		return nil 
// 	}
// 	return result
// }

// // delete a key
// func Delete(key string) error {
// 	return storageService.redisClient.Del(ctx, key).Error
// }