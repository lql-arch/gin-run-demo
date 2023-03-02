package main

import (
	api "douSheng/cmd/favorite/kitex_gen/api/favorite"
	"douSheng/cmd/handler"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8890")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.FavoriteImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
