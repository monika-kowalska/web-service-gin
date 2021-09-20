package daos

import (
	"testing"

	"github.com/monika-kowalska/web-service-gin/test_data"
	"github.com/stretchr/testify/assert"
)

func TestCampaignDAO_Get(t *testing.T) {
	test_data.ResetDB()
	dao := NewCampaignDAO()

	campaign, err := dao.Get(1)

	expected := map[string]string{"Author": "John Doe", "Title": "First Campaign"}

	assert.Nil(t, err)
	assert.Equal(t, expected["Author"], campaign.Author)
	assert.Equal(t, expected["Title"], campaign.Title)
}

func TestCampaignDAO_GetNotPresent(t *testing.T) {
	test_data.ResetDB()
	dao := NewCampaignDAO()

	campaign, err := dao.Get(99)

	assert.NotNil(t, err)
	assert.Equal(t, "", campaign.Author)
	assert.Equal(t, "", campaign.Title)
}
