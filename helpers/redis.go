package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/toluwase1/busha-assessment/models"
	"log"
	"os"
	"time"
)

var contxt = context.Background()

type cacheParams struct {
	host       string
	database      int
	password   string
	expiration time.Duration
}

func NewRedisCache(host string, database int, password string, expiration time.Duration) Cache {
	return &cacheParams{
		host:       host,
		password:   password,
		database:   database,
		expiration: expiration,
	}
}

func (cache *cacheParams) getClient() *redis.Client {
	var opts *redis.Options
	if os.Getenv("LOCAL") == "true" {
		redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))
		opts = &redis.Options{
			Addr:     redisAddress,
			Password: cache.password,
			DB:       cache.database,
		}
	} else {
		var err error
		opts, err = redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
	}
	return redis.NewClient(opts)
}

func (cache *cacheParams) Set(key string, value *[]models.MovieData) {
	client := cache.getClient()
	marshal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = client.Set(contxt, key, string(marshal), cache.expiration*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return
}

func (cache *cacheParams) GetMoviesFromCache(key string) *[]models.MovieData {
	client := cache.getClient()
	val, err := client.Get(contxt, key).Result()
	if err != nil {
		return nil
	}
	var movies []models.MovieData
	err = json.Unmarshal([]byte(val), &movies)
	if err != nil {
		panic(err)
	}
	log.Println("Movies retrieved from cache")
	return &movies
}

func (cache *cacheParams) SetCharToCache(key string, value []models.CharacterList) {
	client := cache.getClient()
	marshal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = client.Set(contxt, key, string(marshal), cache.expiration*time.Second).Err()
	if err != nil {
		panic(err)
	}
	log.Println("Added characters to cache")
	return
}

func (cache *cacheParams) GetCharactersFromCache(key string) []models.CharacterList {
	client := cache.getClient()
	val, err := client.Get(contxt, key).Result()
	if err != nil {
		return nil
	}
	var characters []models.CharacterList
	err = json.Unmarshal([]byte(val), &characters)
	if err != nil {
		panic(err)
	}
	log.Println("Characters retrieved from cache")
	return characters
}
