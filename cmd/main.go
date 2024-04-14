package main

import (
	"app/redis"
	"app/router"
)

func main() {
	router.Route()
	redis.ConnectRedis()
	//middleware.SaveLogRequest()
}
