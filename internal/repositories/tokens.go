package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
)

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

type TokenRepository struct {
	db *gorm.DB
}

func (r *TokenRepository) Delete(brandId int32) error {
	sql := `DELETE FROM ad_g_access_tokens WHERE brand_id = ?`
	return r.db.Exec(sql, brandId).Error
}

func (r *TokenRepository) Update(brandId int32, token string) error {
	return r.db.Model(&models.AccessToken{}).Where("brand_id = ?", brandId).UpdateColumn("token", token).Error
}

func (r *TokenRepository) GetToken(brandId int32) (*models.AccessToken, error) {
	token := &models.AccessToken{}
	res := r.db.Table(token.TableName()).Where("brand_id = ?", brandId).Find(token)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return token, nil
}
