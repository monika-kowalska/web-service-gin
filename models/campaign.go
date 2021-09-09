package models

type Campaign struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateCampaignInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateCampaignInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
