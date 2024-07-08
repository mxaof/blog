package main

import (
	"fmt"
	"log"
	"mxaof_blog/core"
	"mxaof_blog/dao/global"
	"mxaof_blog/dao/mysql"
	"mxaof_blog/dao/redis"
	"mxaof_blog/router"
)

func main() {
	err := core.ViperInit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("globalConfig:", global.GlobalConfig)
	err = redis.RedisInit()
	if err != nil {
		log.Fatal(err)
	}
	err = mysql.MysqlInit()
	if err != nil {
		log.Fatal(err)
	}
	router.NewRouter()
}
