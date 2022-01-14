package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
	"strings"
)

func NewAdCampaignsRepository(db *gorm.DB) *AdCampaignsRepository {
	return &AdCampaignsRepository{db: db}
}

type AdCampaignsRepository struct {
	db *gorm.DB
}

func (r *AdCampaignsRepository) AttachToBrand(brandId int32, adCampaignIds []int64) error {
	if len(adCampaignIds) == 0 {
		return nil
	}

	placeholders := make([]string, 0, len(adCampaignIds))
	args := make([]interface{}, 0, 2*len(adCampaignIds))

	for _, campaignId := range adCampaignIds {
		placeholders = append(placeholders, `(?, ?)`)
		args = append(args, brandId, campaignId)
	}

	sb := strings.Builder{}
	sb.WriteString(`INSERT INTO ad_g_brand_campaigns(brand_id, campaign_id) VALUES `)
	sb.WriteString(strings.Join(placeholders, ","))
	sb.WriteString(`ON CONFLICT DO NOTHING`)
	cmd := sb.String()

	return r.db.Exec(cmd, args...).Error
}

func (r *AdCampaignsRepository) Save(campaigns []*models.AdCampaign) ([]*models.AdCampaign, error) {
	if len(campaigns) == 0 {
		return nil, nil
	}

	campaignIds := make([]int64, len(campaigns))
	for i, c := range campaigns {
		campaignIds[i] = c.Id
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		tx.Exec("lock table ad_g_campaigns in exclusive mode")

		existingCampaigns, err := r.findGCampaigns(tx, campaignIds)
		if err != nil {
			return err
		}
		campaignsSet := map[int64]struct{}{}
		for _, c := range existingCampaigns {
			campaignsSet[c.Id] = struct{}{}
		}

		var forCreate []*models.AdCampaign
		var forUpdate []*models.AdCampaign
		for _, campaign := range campaigns {
			if _, ok := campaignsSet[campaign.Id]; ok {
				forUpdate = append(forUpdate, campaign)
			} else {
				forCreate = append(forCreate, campaign)
			}
		}

		if len(forCreate) > 0 {
			if err = tx.Create(forCreate).Error; err != nil {
				return err
			}
		}

		if len(forUpdate) > 0 {
			if err = tx.Save(forUpdate).Error; err != nil {
				return err
			}
		}
		campaigns = append(forCreate, forUpdate...)

		return nil
	})

	return campaigns, err
}

// UpdateById deletes the record by given id and inserts a new one.
func (r *AdCampaignsRepository) UpdateById(id string, campaign *models.AdCampaign) error {
	deleteSql := `DELETE FROM ad_g_campaigns WHERE campaign_id = ?`

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(deleteSql, id).Error; err != nil {
			return err
		}
		return tx.Create(campaign).Error
	})
}

func (r *AdCampaignsRepository) findGCampaigns(tx *gorm.DB, ids []int64) ([]*models.AdCampaign, error) {
	var campaigns []*models.AdCampaign
	return campaigns, tx.Table("ad_g_campaigns").Where(`campaign_id IN (?)`, ids).Find(&campaigns).Error
}
