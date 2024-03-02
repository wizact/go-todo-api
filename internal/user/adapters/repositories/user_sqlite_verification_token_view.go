package repository

import (
	"context"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
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

func (r *UserSqliteRepository) createOrUpdateUserTokenView(ctx context.Context, tx *gorm.DB, user ua.User) (ua.UserTokenView, error) {
	emptyUserTokenView := ua.UserTokenView{}

	utv := &SqliteUserTokenView{}
	utv.FromDomainEntityToDbModel(user)

	result := tx.
		Where(SqliteUserTokenView{UserID: user.UserId().String()}).
		Assign(*utv).
		FirstOrCreate(&utv)

	if result.Error != nil {
		return emptyUserTokenView, result.Error
	}

	return utv.FromDbModelToDomainEntity(), nil
}

func (dbm *SqliteUserTokenView) FromDomainEntityToDbModel(de ua.User) {
	det := de.Token()
	dbm.UserID = de.UserId().String()
	dbm.VerificationToken = det.VerificationToken()
	dbm.VerificationSalt = det.VerificationSalt()
}

func (dbm SqliteUserTokenView) FromDbModelToDomainEntity() ua.UserTokenView {
	de := ua.NewUserTokenView(uuid.MustParse(dbm.UserID), dbm.VerificationToken, dbm.VerificationSalt)
	return de
}
