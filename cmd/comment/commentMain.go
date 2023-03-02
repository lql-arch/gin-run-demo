package main

import (
	api "douSheng/cmd/comment/kitex_gen/api/commentfunc"
	"douSheng/cmd/handler"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8892")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.CommentFuncImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
