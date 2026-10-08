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
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
	"gorm.io/gorm"
)

type SQLiteUserRepository struct {
	connection *dbinfra.SQLiteConnection
}

func (r *SQLiteUserRepository) SetConnection(cnn *dbinfra.SQLiteConnection) {
	r.connection = cnn
}

func (r *SQLiteUserRepository) Connection() *dbinfra.SQLiteConnection {
	return r.connection
}

func (r *SQLiteUserRepository) FindByID(ctx context.Context, id uuid.UUID) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	u := &SQLiteUserAggregate{UserID: id.String()}

	result := db.WithContext(ctx).Limit(1).First(u)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, domain.ErrUserIDNotFound
	}

	if result.Error != nil {
		return emptyUser, fmt.Errorf("find user by ID: %w", result.Error)
	}

	de, err := u.toDomain()
	if err != nil {
		return emptyUser, fmt.Errorf("map user aggregate: %w", err)
	}

	return de, nil
}

func (r *SQLiteUserRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	emptyUser := ua.User{}
	db, err := r.connection.Open(gorm.Config{})

	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	uev := &SQLiteUserEmailView{Email: email}
	result := db.WithContext(ctx).Where(uev).First(uev)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return emptyUser, domain.ErrUserEmailNotFound
	}

	if result.Error != nil {
		return emptyUser, fmt.Errorf("find user by email: %w", result.Error)
	}

	de, err := uev.toDomain()
	if err != nil {
		return emptyUser, fmt.Errorf("map user email view: %w", err)
	}

	u, err := r.FindByID(ctx, de.ID())

	if err != nil {
		return emptyUser, err
	}

	return u, nil
}

func (r *SQLiteUserRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	emptyUser := ua.User{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	var persistedUser ua.User
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		record := &SQLiteUserAggregate{}
		record.FromDomain(user)

		if err := tx.Create(record).Error; err != nil {
			return fmt.Errorf("persist user aggregate: %w", err)
		}

		mappedUser, err := record.toDomain()
		if err != nil {
			return fmt.Errorf("map persisted user aggregate: %w", err)
		}
		persistedUser = mappedUser
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

func (r *SQLiteUserRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	emptyUser := ua.User{}

	db, err := r.connection.Open(gorm.Config{})
	if err != nil {
		return emptyUser, fmt.Errorf("open user database: %w", err)
	}

	var persistedUser ua.User
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		record := SQLiteUserAggregate{UserID: user.ID().String()}
		if err := tx.First(&record).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return domain.ErrUserIDNotFound
			}
			return fmt.Errorf("find user aggregate for update: %w", err)
		}

		updatedRecord := SQLiteUserAggregate{}
		updatedRecord.FromDomain(user)
		record.ValueData = updatedRecord.ValueData
		if err := tx.Save(&record).Error; err != nil {
			return fmt.Errorf("persist user aggregate update: %w", err)
		}

		mappedUser, err := record.toDomain()
		if err != nil {
			return fmt.Errorf("map persisted user aggregate: %w", err)
		}
		persistedUser = mappedUser
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

type SQLiteUserAggregate struct {
	UserID    string          `gorm:"primaryKey;not null"`
	ValueData SQLiteUserModel `gorm:"serializer:json"`
	CreatedAt int64           `gorm:"autoCreateTime:milli"`
	UpdatedAt int64           `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
}

type SQLiteUserModel struct {
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
func (SQLiteUserAggregate) TableName() string {
	return "users_aggregate"
}

func (d *SQLiteUserAggregate) FromDomain(de ua.User) {
	d.UserID = de.ID().String()
	deu := de.User()
	deup := deu.Phone()
	location := de.Location()
	longitude, latitude := location.Coordinates()
	fn, ln := deu.Name()
	d.ValueData = SQLiteUserModel{
		ID:               de.ID().String(),
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

func (d SQLiteUserAggregate) toDomain() (ua.User, error) {
	userID, err := uuid.Parse(d.UserID)
	if err != nil {
		return ua.User{}, fmt.Errorf("parse persisted user ID: %w", err)
	}

	ph := model.NewPhoneNumber(d.ValueData.CountryCode, d.ValueData.AreaCode, d.ValueData.Number)
	mu := model.RehydrateUser(userID, d.ValueData.FirstName, d.ValueData.LastName, d.ValueData.DateOfBirth, d.ValueData.Email, ph)

	dl := model.NewLocation(d.ValueData.LocationLong, d.ValueData.LocationLat)

	status := ua.RegistrationStatus{
		IsActive:         d.ValueData.IsActive,
		HasVerifiedEmail: d.ValueData.HasVerifiedEmail,
	}
	return ua.RehydrateUser(mu, dl, status), nil
}
