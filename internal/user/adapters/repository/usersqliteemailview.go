package repository

import (
	"fmt"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"gorm.io/gorm"
)

type SQLiteUserEmailView struct {
	UserID           string `gorm:"primaryKey;not null"`
	Email            string
	HasVerifiedEmail bool
	CreatedAt        int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt        int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// TableName overrides grom default table name
func (SQLiteUserEmailView) TableName() string {
	return "users_email_view"
}

func (r *SQLiteUserRepository) saveUserEmailView(tx *gorm.DB, user ua.User) (ua.UserEmailView, error) {
	emptyUserEmailView := ua.UserEmailView{}

	uev := &SQLiteUserEmailView{}
	uev.FromDomain(user)

	result := tx.
		Where(SQLiteUserEmailView{UserID: user.ID().String()}).
		Assign(*uev).
		FirstOrCreate(&uev)

	if result.Error != nil {
		return emptyUserEmailView, fmt.Errorf("persist user email view: %w", result.Error)
	}

	return uev.toDomain()
}

func (d *SQLiteUserEmailView) FromDomain(de ua.User) {
	d.UserID = de.ID().String()
	d.Email = de.Email()
	d.HasVerifiedEmail = de.HasVerifiedEmail()
}

func (d SQLiteUserEmailView) toDomain() (ua.UserEmailView, error) {
	userID, err := uuid.Parse(d.UserID)
	if err != nil {
		return ua.UserEmailView{}, fmt.Errorf("parse persisted user email view ID: %w", err)
	}

	return ua.NewUserEmailView(userID, d.Email, d.HasVerifiedEmail), nil
}
