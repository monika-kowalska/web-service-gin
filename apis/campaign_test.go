package apis

import (
	"net/http"
	"testing"

	"github.com/monika-kowalska/web-service-gin/test_data"
)

func TestCampaign(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a Campaign", "GET", "/api/v1/campaigns/:id", "/api/v1/campaigns/1", "", GetCampaign, http.StatusOK, path + "/campaign_1.json"},
		{"t2 - get a Campaign not Present", "GET", "/api/v1/campaigns/:id", "/api/v1/campaigns/99", "", GetCampaign, http.StatusNotFound, ""},
		{"t3 - get all campaigns", "GET", "api/v1/campaigns", "/api/v1/campaigns", "", GetCampaigns, http.StatusOK, path + "/campaigns.json"},
		{"t4 - create Campaign", "POST", "api/v1/campaigns", "/api/v1/campaigns", "{\"title\":\"Cat Food Campaign\",\"author\":\"Pearl\"}", CreateCampaign, http.StatusOK, path + "/campaign_2.json"},
		{"t5 - update Campaign", "PATCH", "api/v1/campaigns/:id", "/api/v1/campaigns/1", "{\"id\":1,\"author\":\"Perlovekocie\"}", UpdateCampaign, http.StatusOK, path + "/campaign_updated.json"},
		{"t6 - delete Campaign", "DELETE", "api/v1/campaigns/:id", "/api/v1/campaigns/1", "", DeleteCampaign, http.StatusOK, path + "/campaign_deleted.json"},
	})
}
