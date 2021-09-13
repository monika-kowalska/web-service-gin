package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/apis"
	"github.com/monika-kowalska/web-service-gin/config"
)

func main() {
	setupServer("local.db").Run()
}

func setupServer(dbTarget string) *gin.Engine {
	r := gin.Default()

	config.ConnectDataBase(dbTarget)

	v1 := r.Group("/api/v1")

	{
		v1.GET("/campaigns/:id", apis.GetCampaign)
		v1.GET("campaigns", apis.GetCampaigns)
		// v1.POST("/campaigns", apis.CreateCampaign)
	}

	// r.GET("/campaigns", controllers.FindCampaigns)
	// r.POST("/campaigns", controllers.CreateCampaign)
	// r.GET("/campaigns/:id", controllers.FindCampaign)
	// r.PATCH("/campaigns/:id", controllers.UpdateCampaign)
	// r.DELETE("/campaigns/:id", controllers.DeleteCampaign)

	return r
}
