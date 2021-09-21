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
		v1.POST("/campaigns", apis.CreateCampaign)
		v1.PATCH("/campaigns/:id", apis.UpdateCampaign)
		v1.DELETE("/campaigns/:id", apis.DeleteCampaign)
	}

	return r
}
