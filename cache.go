package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

type CacheData struct {
	Value      string
	ExpiryTime time.Time
}

var cache = sync.Map{}

func InMemCache() {
	// set a value
	cache.Store("key1", CacheData{
		Value: "{\"name\": \"tonie\"}",
	})

	// retrieve
	if val, ok := cache.Load("key1"); !ok {
		cachedData := val.(CacheData)
		if !time.Now().Before(cachedData.ExpiryTime) {
			fmt.Println("cache has expired")
			go func() {
				// fetch from db
			}()
			// update our cache
		}
	}
}

func Distibuted() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// set a value
	if err := client.Set("key1", "{\"name\": \"tonie\"}", 2*time.Minute).Err(); err != nil {
		fmt.Println(err.Error())
	}

	// retrive value
	result := client.Get("key1")
	if result.Err() != nil {
		fmt.Println(result.Err().Error())
	}
	result.Val()
}

func FileBased() {
	//
}

func SetValue(key, value string, expiry time.Duration) error {
	cacheDir := "./cache"
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}

	cacheFile := fmt.Sprintf("%s/%s", cacheDir, key)
	expiryTime := time.Now().Add(expiry).Unix()
	data := []byte(fmt.Sprintf("%d:%s", expiryTime, value))
	return os.WriteFile(cacheFile, data, 0644)
}

func GetValue(key string, expiry time.Duration) (string, error) {
	cacheFile := fmt.Sprintf("./cache/%s", key)
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", err
	}

	parts := string(data)
	expirtyTime, err := strconv.ParseInt(parts[:13], 10, 64)
	if err != nil {
		return "", err
	}

	if time.Now().Unix() > expirtyTime {
		return "", errors.New("the value has expired")
	}

	return parts, nil
}
