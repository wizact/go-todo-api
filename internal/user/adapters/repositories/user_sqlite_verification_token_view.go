package repository

import (
	"gorm.io/gorm"
)

type SqliteUserTokenView struct {
	UserID            string `gorm:"primaryKey;not null"`
	VerificationToken string
	VerificationSalt  string
	CreatedAt         int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt         int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// TableName overrides grom default table name
func (SqliteUserTokenView) TableName() string {
	return "users_token_view"
}
