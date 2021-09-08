package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/monika-kowalska/web-service-gin/models"
)

func TestCampaignsCRUD(t *testing.T) {
	dbTarget := "test.db"
	gin.SetMode(gin.TestMode)
	ts := httptest.NewServer(setupServer(dbTarget))

	models.DB.DropTableIfExists(&models.Campaign{}, "campaigns")
	models.ConnectDataBase(dbTarget)

	defer ts.Close()

	t.Run("Create Empty DB", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/campaigns", ts.URL))

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		assert.Equal(t, nil, err)
	})

	t.Run("Populate DB with campaigns", func(t *testing.T) {
		campaigns := []models.Campaign{
			{Title: "First Campaign", Author: "User1"},
			{Title: "Second Campaign", Author: "User2"},
		}

		for _, camp := range campaigns {
			models.DB.Create(&camp)
		}

		var campains_from_db []models.Campaign

		result := models.DB.Find(&campains_from_db)
		var num int64 = 2

		assert.Equal(t, result.RowsAffected, num)
	})

	t.Run("Retrieve Existing ID on Populated DB", func(t *testing.T) {
		resp, _ := http.Get(fmt.Sprintf("%s/campaigns/1", ts.URL))

		defer resp.Body.Close()

		expected := models.Campaign{
			Author: "User1",
			ID:     1,
			Title:  "First Campaign",
		}

		// jak nie potrzebuje err do niczego a funkcja zwraca 2 argumenty w tym error wlasnie i musze jej przypisac
		// w zwiazku z tym drugi argument to moge uzyc _

		body, _ := io.ReadAll(resp.Body)

		// parse json response body in order to get value of data:
		// json response body looks like that:  {"data":{"id":1,"title":"First Campaign","author":"User1"}}
		var c struct {
			Data models.Campaign `json:"data"`
		}

		_ = json.Unmarshal(body, &c)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expected, c.Data)
	})

	t.Run("Add new campaign", func(t *testing.T) {
		payload, _ := json.Marshal(models.Campaign{
			Title:  "Harry Potter",
			Author: "Dupa Jasia",
		})

		resp, err := http.Post(fmt.Sprintf("%s/campaigns", ts.URL), "application/json", bytes.NewReader(payload))

		var campains_from_db []models.Campaign
		result := models.DB.Find(&campains_from_db)
		var num int64 = 3

		assert.Equal(t, result.RowsAffected, num)
		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Update existing campaign", func(t *testing.T) {
		payload, _ := json.Marshal(map[string]interface{}{
			"author": "Lord Voldemort",
		})

		client := &http.Client{}
		url := fmt.Sprintf("%s/campaigns/3", ts.URL)
		req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var campaign models.Campaign

		error := models.DB.Where("id = ?", 3).First(&campaign).Error

		if err != nil {
			log.Fatal(error)
		}

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Lord Voldemort", campaign.Author)
	})

	t.Run("Delete existing campaign", func(t *testing.T) {
		client := &http.Client{}
		url := fmt.Sprintf("%s/campaigns/3", ts.URL)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var campains_from_db []models.Campaign
		result := models.DB.Find(&campains_from_db)
		var num int64 = 2

		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, result.RowsAffected, num)
	})
}
