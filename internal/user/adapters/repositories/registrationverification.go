package repository

type SqliteRegistrationVerification struct {
	UserID       string `gorm:"primaryKey;not null"`
	SecretDigest string `gorm:"not null"`
	ExpiresAt    int64  `gorm:"not null"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli;not null"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;not null"`
}

// TableName overrides GORM's default table name.
func (SqliteRegistrationVerification) TableName() string {
	return "user_registration_verifications"
}
