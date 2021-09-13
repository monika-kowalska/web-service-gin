package controllers

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/config"
	"github.com/monika-kowalska/web-service-gin/models"
)

func FindCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	config.DB.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"data": campaigns})
	fmt.Printf("%#v\n", campaigns)
}

func CreateCampaign(c *gin.Context) {
	// Validate input
	var input models.CreateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create campaign
	campaign := models.Campaign{Title: input.Title, Author: input.Author}
	config.DB.Create(&campaign)

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func FindCampaign(c *gin.Context) { // Get model if exist
	var campaign models.Campaign

	err := config.DB.Where("id = ?", c.Param("id")).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func UpdateCampaign(c *gin.Context) {
	// Get model if exist
	var campaign models.Campaign
	if err := config.DB.Where("id = ?", c.Param("id")).First(&campaign).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&campaign).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign

	err := config.DB.Where("id = ?", c.Param("id")).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&campaign)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
