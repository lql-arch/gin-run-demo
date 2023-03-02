package main

import (
	"douSheng/cmd/handler"
	api "douSheng/cmd/relation/kitex_gen/api/relationfunc"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8893")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.RelationFuncImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
