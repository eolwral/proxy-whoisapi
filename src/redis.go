package main

import (
    "log"
    "time"
    "github.com/fzzy/radix/redis"
)

type RedisClient struct {
   redisConn *redis.Client
}

func GetConnection() *RedisClient {
	redisConn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(1)*time.Second)
	if err != nil {
		log.Println("redis.DialTimeout: ", err)
	}
	return &RedisClient { redisConn }
}

func (client *RedisClient) Set (key string, value string) {
	result := client.redisConn.Cmd("SETEX", key, "86400", value)
	if result.Err !=  nil  {
		log.Println("set:  ",  result.Err)
	}
	return
}	

func (client *RedisClient) Get (key string) string {
	values, err :=  client.redisConn.Cmd("GET", key).Str()
    if err != nil  { return "" }
	return values
}

func (client *RedisClient) Close() {
   client.redisConn.Close()
}
