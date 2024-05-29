package main

import (
	"TTMS_Web/cache"
	"TTMS_Web/conf"
	"TTMS_Web/routes"
)

func main() {
	conf.Init()
	cache.InitRedis()
	r := routes.NewRouter()
	r.Run(conf.Config_.Service.HttpPort)
}
