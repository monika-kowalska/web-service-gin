package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/controllers"
	"github.com/monika-kowalska/web-service-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/campaigns", controllers.FindCampaigns)
	r.POST("/campaigns", controllers.CreateCampaign)
	r.GET("/campaigns/:id", controllers.FindCampaign)
	r.PATCH("/campaigns/:id", controllers.UpdateCampaign)
	r.DELETE("/campaigns/:id", controllers.DeleteCampaign)
	r.Run()
}
