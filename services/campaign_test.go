package services

import (
	"errors"
	"testing"

	"github.com/monika-kowalska/web-service-gin/models"
	"github.com/stretchr/testify/assert"
)

func TestNewCampaignService(t *testing.T) {
	dao := newMockCampaignDAO()
	s := NewCampaignService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestCampaignService_Get(t *testing.T) {
	s := NewCampaignService(newMockCampaignDAO())
	campaign, err := s.Get(2)
	if assert.Nil(t, err) && assert.NotNil(t, campaign) {
		assert.Equal(t, "User2", campaign.Author)
		assert.Equal(t, "Second Campaign", campaign.Title)
	}

	campaign, err = s.Get(100)
	assert.NotNil(t, err)
	assert.Nil(t, campaign)
}

func TestCampaignService_GetCampaigns(t *testing.T) {
	s := NewCampaignService(newMockCampaignDAO())
	campaigns := s.GetCampaigns()
	assert.Equal(t, 2, len(*campaigns))

}

func TestCampaignService_CreateCampaign(t *testing.T) {
	s := NewCampaignService(newMockCampaignDAO())

	input := models.CreateCampaignInput{Title: "test", Author: "user1"}
	campaign := s.CreateCampaign(input)
	assert.Equal(t, "user1", campaign.Author)

}

func TestCampaignService_UpdateCampaign(t *testing.T) {
	s := NewCampaignService(newMockCampaignDAO())

	input := models.UpdateCampaignInput{Title: "test", Author: "user1"}
	campaign, err := s.UpdateCampaign(input)
	assert.Equal(t, "user1", campaign.Author)
	assert.Nil(t, err)

}

func TestCampaignService_DeleteCampaign(t *testing.T) {
	s := NewCampaignService(newMockCampaignDAO())

	campaign, err := s.DeleteCampaign(2)
	assert.Nil(t, err)
	assert.Equal(t, "User2", campaign.Author)
}

func newMockCampaignDAO() campaignDAO {
	return &mockCampaignDAO{
		records: []models.Campaign{
			{ID: 1, Title: "First Campaign", Author: "User1"},
			{ID: 2, Title: "Second Campaign", Author: "User2"},
		},
	}
}

// Mock Get function that replaces real Campaign DAO
func (m *mockCampaignDAO) Get(id uint) (*models.Campaign, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockCampaignDAO struct {
	records []models.Campaign
}

//typ od bloga mial tylko jedna funkcje Get(id) i ja testowal a ze ja mialam wszystkie crudowe
//funkcje to wszystkie je tu muszezamokowac i przetestowac
//to dlatego func newMockCampaignDAO() campaignDAO { zwaracalo mi error bo moje campaignDAO implementuje 5 metod
//i wszystkie musze zawrzec w tescie

func (m *mockCampaignDAO) GetAll() *[]models.Campaign {
	return &m.records
}

func (m *mockCampaignDAO) CreateCampaign(input models.CreateCampaignInput) *models.Campaign {
	campaign := models.Campaign{Title: input.Title, Author: input.Author}
	return &campaign
}

func (m *mockCampaignDAO) UpdateCampaign(input models.UpdateCampaignInput) (*models.Campaign, error) {
	campaign := models.Campaign{Title: input.Title, Author: input.Author}
	return &campaign, nil
}

func (m *mockCampaignDAO) DeleteCampaign(id uint) (*models.Campaign, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}
