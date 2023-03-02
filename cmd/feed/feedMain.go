package main

import (
	api "douSheng/cmd/feed/kitex_gen/api/feed"
	"douSheng/cmd/handler"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.FeedImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
