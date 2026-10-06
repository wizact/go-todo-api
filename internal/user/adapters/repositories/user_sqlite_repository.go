package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	"github.com/wizact/go-todo-api/internal/user/domain"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	"gorm.io/gorm"
)

type UserSqliteRepository struct {
	connection *dbinfra.SqliteConnection
}

func (r *UserSqliteRepository) Connection(cnn *dbinfra.SqliteConnection) {
	r.connection = cnn
}

func (r *UserSqliteRepository) GetConnection() *dbinfra.SqliteConnection {
	return r.connection
}

func (r *UserSqliteRepository) FindById(ctx context.Context, id uuid.UUID) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	u := &SqliteUserAggregate{UserID: id.String()}

	result := db.WithContext(ctx).Limit(1).First(u)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, domain.ErrUserIDNotFound
	}

	if result.Error != nil {
		return emptyUser, fmt.Errorf("find user by ID: %w", result.Error)
	}

	de := u.FromDbModelToDomainEntity()

	return de, nil
}

func (r *UserSqliteRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	uev := &SqliteUserEmailView{Email: email}
	result := db.WithContext(ctx).Where(uev).First(uev)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, domain.ErrUserEmailNotFound
	}

	if result.Error != nil {
		return emptyUser, fmt.Errorf("find user by email: %w", result.Error)
	}

	de := uev.FromDbModelToDomainEntity()

	u, err := r.FindById(ctx, de.Id())

	if err != nil {
		return emptyUser, err
	}

	return u, nil
}

func (r *UserSqliteRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	emptyUser := ua.User{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	var persistedUser ua.User
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		record := &SqliteUserAggregate{}
		record.FromDomainEntityToDbModel(user)

		if err := tx.Create(record).Error; err != nil {
			return fmt.Errorf("persist user aggregate: %w", err)
		}

		persistedUser = record.FromDbModelToDomainEntity()
		if _, err := r.saveUserEmailView(tx, persistedUser); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return emptyUser, fmt.Errorf("create user: %w", err)
	}

	return persistedUser, nil
}

func (r *UserSqliteRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	emptyUser := ua.User{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	var persistedUser ua.User
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		record := SqliteUserAggregate{UserID: user.UserId().String()}
		if err := tx.First(&record).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return domain.ErrUserIDNotFound
			}
			return fmt.Errorf("find user aggregate for update: %w", err)
		}

		updatedRecord := SqliteUserAggregate{}
		updatedRecord.FromDomainEntityToDbModel(user)
		record.ValueData = updatedRecord.ValueData
		if err := tx.Save(&record).Error; err != nil {
			return fmt.Errorf("persist user aggregate update: %w", err)
		}

		persistedUser = record.FromDbModelToDomainEntity()
		if _, err := r.saveUserEmailView(tx, persistedUser); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return emptyUser, fmt.Errorf("update user: %w", err)
	}

	return persistedUser, nil
}

type SqliteUserAggregate struct {
	UserID    string          `gorm:"primaryKey;not null"`
	ValueData SqliteUserModel `gorm:"serializer:json"`
	CreatedAt int64           `gorm:"autoCreateTime:milli"`
	UpdatedAt int64           `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
}

type SqliteUserModel struct {
	ID          string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Email       string
	CountryCode string
	AreaCode    string
	Number      string

	LocationLong float64
	LocationLat  float64

	HasVerifiedEmail bool
	IsActive         bool
}

// TableName overrides grom default table name
func (SqliteUserAggregate) TableName() string {
	return "users_aggregate"
}

func (dbm *SqliteUserAggregate) FromDomainEntityToDbModel(de ua.User) {
	dbm.UserID = de.UserId().String()
	deu := de.User()
	deup := deu.Phone()
	location := de.Location()
	longitude, latitude := location.Coordinates()
	fn, ln := deu.Name()
	dbm.ValueData = SqliteUserModel{
		ID:               de.UserId().String(),
		FirstName:        fn,
		LastName:         ln,
		DateOfBirth:      deu.DateOfBirth(),
		Email:            deu.Email(),
		CountryCode:      deup.CountryCode(),
		AreaCode:         deup.AreaCode(),
		Number:           deup.Number(),
		LocationLong:     longitude,
		LocationLat:      latitude,
		HasVerifiedEmail: de.HasVerifiedEmail(),
		IsActive:         de.IsActive(),
	}
}

func (dbm SqliteUserAggregate) FromDbModelToDomainEntity() ua.User {
	ph := model.NewPhoneNumber(dbm.ValueData.CountryCode, dbm.ValueData.AreaCode, dbm.ValueData.Number)
	mu := model.RehydrateUser(uuid.MustParse(dbm.UserID), dbm.ValueData.FirstName, dbm.ValueData.LastName, dbm.ValueData.DateOfBirth, dbm.ValueData.Email, ph)

	dl := model.NewLocation(dbm.ValueData.LocationLong, dbm.ValueData.LocationLat)

	status := ua.RegistrationStatus{
		IsActive:         dbm.ValueData.IsActive,
		HasVerifiedEmail: dbm.ValueData.HasVerifiedEmail,
	}
	return ua.RehydrateUser(mu, dl, status)
}
