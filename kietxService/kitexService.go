package main

import (
	api "douSheng/kitex/kitex_gen/api/echo"
	service2 "douSheng/service"
	"log"
)

func main() {
	svr := api.NewServer(new(service2.EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
