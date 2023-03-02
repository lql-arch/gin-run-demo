package main

import (
	"douSheng/cmd/handler"
	api "douSheng/cmd/publish/kitex_gen/api/publish"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8891")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.PublishImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
