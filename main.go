package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//go kClient.KClient()

	r := gin.Default()

	initRouter(r)

	err := r.Run()
	if err != nil {
		panic("run failed.")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
