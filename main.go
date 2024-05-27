package main

import (
	"TTMS_Web/conf"
	"TTMS_Web/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(conf.Config_.Service.HttpPort)
}
