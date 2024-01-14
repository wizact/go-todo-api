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
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// TableName overrides grom default table name
func (SqliteUserEmailView) TableName() string {
	return "users_email_view"
}

func (r *UserSqliteRepository) CreateOrUpdateUserEmailView(ctx context.Context, user ua.User) (ua.UserEmailView, error) {
	emptyUserEmailView := ua.UserEmailView{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUserEmailView, err
	}

	uev := &SqliteUserEmailView{}
	uev.FromDomainEntityToDbModel(user)

	result := db.
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
