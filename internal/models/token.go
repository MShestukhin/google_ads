package models

import (
	"database/sql"
)

type AccessToken struct {
	BrandId      int32
	Token        string
	RefreshToken string
	ExpiresAt    sql.NullTime
}

func (AccessToken) TableName() string {
	return "ad_g_access_tokens"
}
