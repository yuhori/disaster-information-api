package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuhori/disaster-information-api/controllers"
)

func Init() {
	r := gin.Default()
	r.GET("/high", controllers.High)
	r.Run()
}
