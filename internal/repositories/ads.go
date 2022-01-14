package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewAdsRepository(db *gorm.DB) *AdsRepository {
	return &AdsRepository{db: db}
}

type AdsRepository struct {
	db *gorm.DB
}

func (r *AdsRepository) Save(ads []*models.Ads) error {
	if len(ads) == 0 {
		return nil
	}
	return r.db.Table(ads[0].TableName()).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(ads).Error
}
