package kClient

import (
	"douSheng/kitex/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"log"
)

var kClient echo.Client

func KClient() {
	c, err := echo.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	//for {
	//	req := &api.Message{}
	//	resp, err := c.Message(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(resp)
	//}
	kClient = c
}
