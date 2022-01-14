package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
)

func NewAdSetInsightRepository(db *gorm.DB) *AdGroupsInsightRepository {
	return &AdGroupsInsightRepository{db: db}
}

type AdGroupsInsightRepository struct {
	db *gorm.DB
}

func (r *AdGroupsInsightRepository) Create(m []*models.AdGroupInsight) ([]*models.AdGroupInsight, error) {
	if len(m) == 0 {
		return nil, nil
	}

	return m, r.db.Create(&m).Error
}
