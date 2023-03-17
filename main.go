package main

import (
	"fmt"
	"github.com/ChangQingAAS/Learing_GIN/config"
	"github.com/ChangQingAAS/Learing_GIN/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.InitRouter(engine)
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
