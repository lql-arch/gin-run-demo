package main

import (
	"douSheng/cmd/handler"
	api "douSheng/cmd/user/kitex_gen/api/userinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8889")

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := api.NewServer(new(handler.UserInfoImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
