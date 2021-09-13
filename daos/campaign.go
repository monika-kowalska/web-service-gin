package daos

import (
	"github.com/monika-kowalska/web-service-gin/config"
	"github.com/monika-kowalska/web-service-gin/models"
)

// CampaignDAO persists campaign data in database
type CampaignDAO struct{}

//NewCampaignDAO creates a new CampaignDAO
func NewCampaignDAO() *CampaignDAO {
	return &CampaignDAO{}
}

//Get does the actual query to database
func (dao *CampaignDAO) Get(id uint) (*models.Campaign, error) {
	var campaign models.Campaign

	err := config.DB.Where("id = ?", id).First(&campaign).Error
	return &campaign, err
}

func (dao *CampaignDAO) GetAll() *[]models.Campaign {
	var campaigns []models.Campaign
	config.DB.Find(&campaigns)
	return &campaigns
}

// func (dao *CampaignDAO) CreateCampaign(input *models.CreateCampaignInput) *models.Campaign {
// 	campaign := models.Campaign{Title: input.Title, Author: input.Author}
// 	config.DB.Create(&campaign)
// 	return &campaign
// }
