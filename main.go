package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/controllers"
	"github.com/monika-kowalska/web-service-gin/models"
)

func main() {
	setupServer("test.db").Run()
}

func setupServer(dbTarget string) *gin.Engine {
	r := gin.Default()

	models.ConnectDataBase(dbTarget)

	r.GET("/campaigns", controllers.FindCampaigns)
	r.POST("/campaigns", controllers.CreateCampaign)
	r.GET("/campaigns/:id", controllers.FindCampaign)
	r.PATCH("/campaigns/:id", controllers.UpdateCampaign)
	r.DELETE("/campaigns/:id", controllers.DeleteCampaign)

	return r
}
