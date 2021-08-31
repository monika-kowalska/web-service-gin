package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monika-kowalska/web-service-gin/models"
)

func FindCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	models.DB.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}

type CreateCampaignInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateCampaignInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func CreateCampaign(c *gin.Context) {
	// Validate input
	var input CreateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create campaign
	campaign := models.Campaign{Title: input.Title, Author: input.Author}
	models.DB.Create(&campaign)

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func FindCampaign(c *gin.Context) { // Get model if exist
	var campaign models.Campaign

	err := models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func UpdateCampaign(c *gin.Context) {
	// Get model if exist
	var campaign models.Campaign
	if err := models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&campaign).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign

	err := models.DB.Where("id = ?", c.Param("id")).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&campaign)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
