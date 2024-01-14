package repository

import (
	"context"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	"gorm.io/gorm"
)

type SqliteUserEmailView struct {
	UserID           string `gorm:"primaryKey;not null"`
	Email            string
	HasVerifiedEmail bool
	CreatedAt        int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt        int64          `gorm:"autoUpdateTime:milli"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// TableName overrides grom default table name
func (SqliteUserEmailView) TableName() string {
	return "users_email_view"
}

func (r *UserSqliteRepository) createOrUpdateUserEmailView(ctx context.Context, tx *gorm.DB, user ua.User) (ua.UserEmailView, error) {
	emptyUserEmailView := ua.UserEmailView{}

	uev := &SqliteUserEmailView{}
	uev.FromDomainEntityToDbModel(user)

	result := tx.
		Where(SqliteUserEmailView{UserID: user.UserId().String()}).
		Assign(*uev).
		FirstOrCreate(&uev)

	if result.Error != nil {
		return emptyUserEmailView, result.Error
	}

	return uev.FromDbModelToDomainEntity(), nil
}

func (dbm *SqliteUserEmailView) FromDomainEntityToDbModel(de ua.User) {
	dbm.UserID = de.UserId().String()
	dbm.Email = de.Email()
	dbm.HasVerifiedEmail = de.HasVerifiedEmail()
}

func (dbm SqliteUserEmailView) FromDbModelToDomainEntity() ua.UserEmailView {
	de := ua.NewUserEmailView(uuid.MustParse(dbm.UserID), dbm.Email, dbm.HasVerifiedEmail)
	return de
}
