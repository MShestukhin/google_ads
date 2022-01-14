package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
)

func NewCampaignInsightRepository(db *gorm.DB) *CampaignInsightRepository {
	return &CampaignInsightRepository{db: db}
}

type CampaignInsightRepository struct {
	db *gorm.DB
}

func (r *CampaignInsightRepository) Create(m []*models.CampaignInsight) ([]*models.CampaignInsight, error) {
	if len(m) == 0 {
		return nil, nil
	}

	return m, r.db.Create(&m).Error
}
