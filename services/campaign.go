package services

import "github.com/monika-kowalska/web-service-gin/models"

type campaignDAO interface {
	Get(id uint) (*models.Campaign, error)
	GetAll() *[]models.Campaign
	CreateCampaign(input models.CreateCampaignInput) *models.Campaign
	UpdateCampaign(input models.UpdateCampaignInput) (*models.Campaign, error)
	DeleteCampaign(id uint) (*models.Campaign, error)
}

type CampaignService struct {
	dao campaignDAO
}

func NewCampaignService(dao campaignDAO) *CampaignService {
	return &CampaignService{dao}
}

func (s *CampaignService) Get(id uint) (*models.Campaign, error) {
	return s.dao.Get(id)
}

func (s *CampaignService) GetCampaigns() *[]models.Campaign {
	return s.dao.GetAll()
}

func (s *CampaignService) CreateCampaign(input models.CreateCampaignInput) *models.Campaign {
	return s.dao.CreateCampaign(input)
}

func (s *CampaignService) UpdateCampaign(input models.UpdateCampaignInput) (*models.Campaign, error) {
	return s.dao.UpdateCampaign(input)
}

func (s *CampaignService) DeleteCampaign(id uint) (*models.Campaign, error) {
	return s.dao.DeleteCampaign(id)
}
