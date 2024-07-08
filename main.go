package main

import (
	"log"
	"mxaof_blog/dao/redis"
	"mxaof_blog/router"
)

func main() {
	err := redis.RedisInit()
	if err != nil {
		log.Fatal(err)
	}
	router.NewRouter()
}
