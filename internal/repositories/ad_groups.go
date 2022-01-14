package repositories

import (
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gorm.io/gorm"
)

func NewAdGroupsRepository(db *gorm.DB) *AdGroupsRepository {
	return &AdGroupsRepository{db: db}
}

type AdGroupsRepository struct {
	db *gorm.DB
}

func (r *AdGroupsRepository) Save(groups []*models.AdGroup) ([]*models.AdGroup, error) {
	if len(groups) == 0 {
		return nil, nil
	}

	groupIds := make([]int64, len(groups))
	for i, c := range groups {
		groupIds[i] = c.Id
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		tx.Exec("lock table ad_g_groups in exclusive mode")

		existingGroups, err := r.findGGroups(tx, groupIds)
		if err != nil {
			return err
		}
		campaignsSet := map[int64]struct{}{}
		for _, c := range existingGroups {
			campaignsSet[c.Id] = struct{}{}
		}

		var forCreate []*models.AdGroup
		var forUpdate []*models.AdGroup
		for _, group := range groups {
			if _, ok := campaignsSet[group.Id]; ok {
				forUpdate = append(forUpdate, group)
			} else {
				forCreate = append(forCreate, group)
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
		groups = append(forCreate, forUpdate...)

		return nil
	})

	return groups, err
}

// UpdateById deletes the record by given id and inserts a new one.
func (r *AdGroupsRepository) UpdateById(id string, group *models.AdGroup) error {
	deleteSql := `DELETE FROM ad_g_groups WHERE id = ?`

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(deleteSql, id).Error; err != nil {
			return err
		}
		return tx.Create(group).Error
	})
}

func (r *AdGroupsRepository) findGGroups(tx *gorm.DB, ids []int64) ([]*models.AdGroup, error) {
	var groups []*models.AdGroup
	return groups, tx.Table("ad_g_groups").Where(`id IN (?)`, ids).Find(&groups).Error
}
