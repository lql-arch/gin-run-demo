package main

import (
	"douSheng/cmd/handler"
	api "douSheng/cmd/message/kitex_gen/api/messagefunc"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8894")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.MessageFuncImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
