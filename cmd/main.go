package main

import (
	"app/dbutil"
	"app/redis"
	"app/router"
)

func main() {
	dbutil.ConnectDB()

	router.Route()
	redis.ConnectRedis()

}
