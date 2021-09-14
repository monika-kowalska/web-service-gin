package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/daos"
	"github.com/monika-kowalska/web-service-gin/models"
	"github.com/monika-kowalska/web-service-gin/services"
)

func campaignService() *services.CampaignService {
	return services.NewCampaignService(daos.NewCampaignDAO())
}

func GetCampaign(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if campaign, err := campaignService().Get(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		c.JSON(http.StatusOK, campaign)
	}
}

func GetCampaigns(c *gin.Context) {
	campaigns := campaignService().GetCampaigns()
	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}

func CreateCampaign(c *gin.Context) {
	var input models.CreateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	campaign := campaignService().CreateCampaign(input)
	c.JSON(http.StatusOK, campaign)
}

func UpdateCampaign(c *gin.Context) {
	var input models.UpdateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign, err := campaignService().UpdateCampaign(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func DeleteCampaign(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	_, err := campaignService().DeleteCampaign(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
