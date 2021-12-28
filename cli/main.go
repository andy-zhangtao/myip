package main

import (
	kg "github.com/KataSpace/Kata-Gin"
	"github.com/andy-zhangtao/myip/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r = kg.RegisterRouter(r, nil, nil, &service.MyIP{})
	r.Run(":9000")
}
